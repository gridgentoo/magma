#!/usr/bin/env python3
"""
Copyright (c) 2019-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.


ServiceManager manages the lifecycle and chaining of network services,
which are cloud managed and provide discrete network functions.

These network services consist of Ryu apps, which operate on tables managed by
the ServiceManager. OVS provides a set number of tables that can be
programmed to match and modify traffic. We split these tables two categories,
main tables and scratch tables.

All apps from the same service are associated with a main table, which is
visible to other services and they are used to forward traffic between
different services.

Apps can also optionally claim additional scratch tables, which may be
required for complex flow matching and aggregation use cases. Scratch tables
should not be accessible to apps from other services.
"""
# pylint: skip-file
# pylint does not play well with aioeventlet, as it uses asyncio.async which
# produces a parse error

import asyncio
import logging
from concurrent.futures import Future
from collections import namedtuple, OrderedDict
from typing import List

import aioeventlet
from lte.protos.mconfig.mconfigs_pb2 import PipelineD
from lte.protos.mobilityd_pb2_grpc import MobilityServiceStub
from lte.protos.session_manager_pb2_grpc import LocalSessionManagerStub
from magma.pipelined.app.base import ControllerType
from magma.pipelined.app import of_rest_server
from magma.pipelined.app.access_control import AccessControlController
from magma.pipelined.app.tunnel_learn import TunnelLearnController
from magma.pipelined.app.vlan_learn import VlanLearnController
from magma.pipelined.app.arp import ArpController
from magma.pipelined.app.dpi import DPIController
from magma.pipelined.app.enforcement import EnforcementController
from magma.pipelined.app.ipfix import IPFIXController
from magma.pipelined.app.li_mirror import LIMirrorController
from magma.pipelined.app.enforcement_stats import EnforcementStatsController
from magma.pipelined.app.inout import EGRESS, INGRESS, PHYSICAL_TO_LOGICAL, \
    InOutController
from magma.pipelined.app.ue_mac import UEMacAddressController
from magma.pipelined.app.xwf_passthru import XWFPassthruController
from magma.pipelined.app.startup_flows import StartupFlows
from magma.pipelined.app.check_quota import CheckQuotaController
from magma.pipelined.rule_mappers import RuleIDToNumMapper, \
    SessionRuleToVersionMapper
from ryu.base.app_manager import AppManager

from magma.common.service import MagmaService
from magma.common.service_registry import ServiceRegistry
from magma.configuration import environment

# Type is either Physical or Logical, highest order_priority is at zero
App = namedtuple('App', ['name', 'module', 'type', 'order_priority'])


class Tables:
    __slots__ = ['main_table', 'type', 'scratch_tables']

    def __init__(self, main_table, type, scratch_tables=None):
        self.main_table = main_table
        self.type = type
        self.scratch_tables = scratch_tables
        if self.scratch_tables is None:
            self.scratch_tables = []


class TableNumException(Exception):
    """
    Exception used for when table number allocation fails.
    """
    pass


class TableRange():
    """
    Used to generalize different table ranges.
    """

    def __init__(self, start: int, end: int):
        self._start = start
        self._end = end
        self._next_table = self._start

    def allocate_table(self):
        if (self._next_table == self._end):
            raise TableNumException('Cannot generate more tables. Table limit'
                                    'of %s reached!' % self._end)
        table_num = self._next_table
        self._next_table += 1
        return table_num

    def allocate_tables(self, count: int):
        if self._next_table + count >= self._end:
            raise TableNumException('Cannot generate more tables. Table limit'
                                    'of %s reached!' % self._end)
        tables = [self.allocate_table() for i in range(0, count)]
        return tables

    def get_next_table(self, table: int):
        if table + 1 < self._next_table:
            return table + 1
        else:
            return self._end


class _TableManager:
    """
    TableManager maintains an internal mapping between apps to their
    main and scratch tables.
    """

    INGRESS_TABLE_NUM = 1
    PHYSICAL_TO_LOGICAL_TABLE_NUM = 10
    EGRESS_TABLE_NUM = 20
    LOGICAL_TABLE_LIMIT_NUM = EGRESS_TABLE_NUM  # exclusive
    SCRATCH_TABLE_START_NUM = EGRESS_TABLE_NUM + 1  # 21
    SCRATCH_TABLE_LIMIT_NUM = 255  # exclusive

    def __init__(self):
        self._table_ranges = {
            ControllerType.PHYSICAL: TableRange(self.INGRESS_TABLE_NUM + 1,
                                                self.PHYSICAL_TO_LOGICAL_TABLE_NUM),
            ControllerType.LOGICAL:
                TableRange(self.PHYSICAL_TO_LOGICAL_TABLE_NUM + 1,
                           self.EGRESS_TABLE_NUM)
        }
        self._scratch_range = TableRange(self.SCRATCH_TABLE_START_NUM,
                                         self.SCRATCH_TABLE_LIMIT_NUM)
        self._tables_by_app = {
            INGRESS: Tables(main_table=self.INGRESS_TABLE_NUM,
                            type=ControllerType.SPECIAL),
            PHYSICAL_TO_LOGICAL: Tables(
                main_table=self.PHYSICAL_TO_LOGICAL_TABLE_NUM,
                type=ControllerType.SPECIAL),
            EGRESS: Tables(main_table=self.EGRESS_TABLE_NUM,
                           type=ControllerType.SPECIAL),
        }

    def _allocate_main_table(self, type: ControllerType) -> int:
        if type not in self._table_ranges:
            raise TableNumException('Cannot generate a table for %s' % type)
        return self._table_ranges[type].allocate_table()

    def register_apps_for_service(self, apps: List[App]):
        """
        Register the apps for a service with a main table. All Apps must share
        the same contoller type
        """
        if not all(apps[0].type == app.type for app in apps):
            raise TableNumException('Cannot register apps with different'
                                    'controller type')
        table_num = self._allocate_main_table(apps[0].type)
        for app in apps:
            self._tables_by_app[app.name] = Tables(main_table=table_num,
                                                   type=app.type)

    def register_apps_for_table0_service(self, apps: List[App]):
        """
        Register the apps for a service with main table 0
        """
        for app in apps:
            self._tables_by_app[app.name] = Tables(main_table=0, type=app.type)

    def get_table_num(self, app_name: str) -> int:
        if app_name not in self._tables_by_app:
            raise Exception('App is not registered: %s' % app_name)
        return self._tables_by_app[app_name].main_table

    def get_next_table_num(self, app_name: str) -> int:
        """
        Returns the main table number of the next service.
        If there are no more services after the current table, return the
        EGRESS table
        """
        if app_name not in self._tables_by_app:
            raise Exception('App is not registered: %s' % app_name)

        app = self._tables_by_app[app_name]
        if app.type == ControllerType.SPECIAL:
            if app_name == INGRESS:
                return self._table_ranges[ControllerType.PHYSICAL].get_next_table(app.main_table)
            elif app_name == PHYSICAL_TO_LOGICAL:
                return self._table_ranges[ControllerType.LOGICAL].get_next_table(app.main_table)
            else:
                raise TableNumException('No next table found for %s' % app_name)
        return self._table_ranges[app.type].get_next_table(app.main_table)

    def is_app_enabled(self, app_name: str) -> bool:
        return app_name in self._tables_by_app or \
            app_name == InOutController.APP_NAME

    def allocate_scratch_tables(self, app_name: str, count: int) -> \
            List[int]:

        tbl_nums = self._scratch_range.allocate_tables(count)
        self._tables_by_app[app_name].scratch_tables.extend(tbl_nums)
        return tbl_nums

    def get_scratch_table_nums(self, app_name: str) -> List[int]:
        if app_name not in self._tables_by_app:
            raise Exception('App is not registered: %s' % app_name)
        return self._tables_by_app[app_name].scratch_tables

    def get_all_table_assignments(self) -> 'OrderedDict[str, Tables]':
        resp = OrderedDict(sorted(self._tables_by_app.items(),
                                  key=lambda kv: (kv[1].main_table, kv[0])))
        # Include table 0 when it is managed by the EPC, for completeness.
        if not any(table in ['ue_mac', 'xwf_passthru'] for table in self._tables_by_app):
            resp['mme'] = Tables(main_table=0, type=None)
            resp.move_to_end('mme', last=False)
        return resp


class ServiceManager:
    """
    ServiceManager manages the service lifecycle and chaining of services for
    the Ryu apps. Ryu apps are loaded based on the services specified in the
    YAML config for static apps and mconfig for dynamic apps.
    ServiceManager also maintains a mapping between apps to the flow
    tables they use.

    Currently, its use cases include:
        - Starting all Ryu apps
        - Flow table number lookup for Ryu apps
        - Main & scratch tables management
    """

    UE_MAC_ADDRESS_SERVICE_NAME = 'ue_mac'
    ARP_SERVICE_NAME = 'arpd'
    ACCESS_CONTROL_SERVICE_NAME = 'access_control'
    TUNNEL_LEARN_SERVICE_NAME = 'tunnel_learn'
    VLAN_LEARN_SERVICE_NAME = 'vlan_learn'
    IPFIX_SERVICE_NAME = 'ipfix'
    RYU_REST_SERVICE_NAME = 'ryu_rest_service'
    RYU_REST_APP_NAME = 'ryu_rest_app'
    STARTUP_FLOWS_RECIEVER_CONTROLLER = 'startup_flows'
    CHECK_QUOTA_SERVICE_NAME = 'check_quota'
    LI_MIRROR_SERVICE_NAME = 'li_mirror'
    XWF_PASSTHRU_NAME = 'xwf_passthru'

    # Mapping between services defined in mconfig and the names and modules of
    # the corresponding Ryu apps in PipelineD. The module is used for the Ryu
    # app manager to instantiate the app.
    # Note that a service may require multiple apps.
    DYNAMIC_SERVICE_TO_APPS = {
        PipelineD.ENFORCEMENT: [
            App(name=EnforcementController.APP_NAME,
                module=EnforcementController.__module__,
                type=EnforcementController.APP_TYPE,
                order_priority=500),
            App(name=EnforcementStatsController.APP_NAME,
                module=EnforcementStatsController.__module__,
                type=EnforcementStatsController.APP_TYPE,
                order_priority=501),
        ],
        PipelineD.DPI: [
            App(name=DPIController.APP_NAME, module=DPIController.__module__,
                type=DPIController.APP_TYPE,
                order_priority=700),
        ],
    }

    # Mapping between the app names defined in pipelined.yml and the names and
    # modules of their corresponding Ryu apps in PipelineD.
    STATIC_SERVICE_TO_APPS = {
        UE_MAC_ADDRESS_SERVICE_NAME: [
            App(name=UEMacAddressController.APP_NAME,
                module=UEMacAddressController.__module__,
                type=None,
                order_priority=0),
        ],
        ARP_SERVICE_NAME: [
            App(name=ArpController.APP_NAME, module=ArpController.__module__,
                type=ArpController.APP_TYPE,
                order_priority=200)
        ],
        ACCESS_CONTROL_SERVICE_NAME: [
            App(name=AccessControlController.APP_NAME,
                module=AccessControlController.__module__,
                type=AccessControlController.APP_TYPE,
                order_priority=400),
        ],
        TUNNEL_LEARN_SERVICE_NAME: [
            App(name=TunnelLearnController.APP_NAME,
                module=TunnelLearnController.__module__,
                type=TunnelLearnController.APP_TYPE,
                order_priority=300),
        ],
        VLAN_LEARN_SERVICE_NAME: [
            App(name=VlanLearnController.APP_NAME,
                module=VlanLearnController.__module__,
                type=VlanLearnController.APP_TYPE,
                order_priority=500),
        ],
        RYU_REST_SERVICE_NAME: [
            App(name=RYU_REST_APP_NAME,
                module='ryu.app.ofctl_rest',
                type=None,
                order_priority=0),
        ],
        STARTUP_FLOWS_RECIEVER_CONTROLLER: [
            App(name=StartupFlows.APP_NAME,
                module=StartupFlows.__module__,
                type=StartupFlows.APP_TYPE,
                order_priority=0),
        ],
        CHECK_QUOTA_SERVICE_NAME: [
            App(name=CheckQuotaController.APP_NAME,
                module=CheckQuotaController.__module__,
                type=CheckQuotaController.APP_TYPE,
                order_priority=300),
        ],
        IPFIX_SERVICE_NAME: [
            App(name=IPFIXController.APP_NAME,
                module=IPFIXController.__module__,
                type=IPFIXController.APP_TYPE,
                order_priority=800),
        ],
        LI_MIRROR_SERVICE_NAME: [
            App(name=LIMirrorController.APP_NAME,
                module=LIMirrorController.__module__,
                type=LIMirrorController.APP_TYPE,
                order_priority=900),
        ],
        XWF_PASSTHRU_NAME: [
            App(name=XWFPassthruController.APP_NAME,
                module=XWFPassthruController.__module__,
                type=XWFPassthruController.APP_TYPE,
                order_priority=0),
        ],
    }

    # Some apps do not use a table, so they need to be excluded from table
    # allocation.
    STATIC_APP_WITH_NO_TABLE = [
        RYU_REST_APP_NAME,
        StartupFlows.APP_NAME,
    ]

    def __init__(self, magma_service: MagmaService):
        self._magma_service = magma_service
        # inout is a mandatory app and it occupies:
        #   table 1(for ingress)
        #   table 10(for middle)
        #   table 20(for egress)
        self._apps = [App(name=InOutController.APP_NAME,
                          module=InOutController.__module__,
                          type=None,
                          order_priority=0)]
        self._table_manager = _TableManager()
        self.session_rule_version_mapper = SessionRuleToVersionMapper()

        apps = self._get_static_apps()
        apps.extend(self._get_dynamic_apps())
        apps.sort(key=lambda x: x.order_priority)

        self._apps.extend(apps)
        # Filter out reserved apps and apps that don't need a table
        for app in apps:
            if app.name in self.STATIC_APP_WITH_NO_TABLE:
                continue
            # UE MAC service must be registered with Table 0
            if app.name in [self.UE_MAC_ADDRESS_SERVICE_NAME, self.XWF_PASSTHRU_NAME]:
                self._table_manager.register_apps_for_table0_service([app])
                continue
            self._table_manager.register_apps_for_service([app])

    def _get_static_apps(self):
        """
        _init_static_services populates app modules and allocates a main table
        for each static service.
        """
        static_services = self._magma_service.config['static_services']
        static_apps = \
            [app for service in static_services for app in
             self.STATIC_SERVICE_TO_APPS[service]]

        return static_apps

    def _get_dynamic_apps(self):
        """
        _init_dynamic_services populates app modules and allocates a main table
        for each dynamic service.
        """
        dynamic_services = []
        for service in self._magma_service.mconfig.services:
            if service not in self.DYNAMIC_SERVICE_TO_APPS:
                # Most likely cause: the config contains a deprecated
                # pipelined service.
                # Fix: update the relevant network's network_services settings.
                logging.warning(
                    'Mconfig contains unsupported network_services service: %s',
                    service,
                )
                continue
            dynamic_services.append(service)

        dynamic_apps = [app for service in dynamic_services for
                        app in self.DYNAMIC_SERVICE_TO_APPS[service]]
        return dynamic_apps

    def load(self):
        """
        Instantiates and schedules the Ryu app eventlets in the service
        eventloop.
        """
        manager = AppManager.get_instance()
        manager.load_apps([app.module for app in self._apps])
        contexts = manager.create_contexts()
        contexts['rule_id_mapper'] = RuleIDToNumMapper()
        contexts[
            'session_rule_version_mapper'] = self.session_rule_version_mapper
        contexts['app_futures'] = {app.name: Future() for app in self._apps}
        contexts['config'] = self._magma_service.config
        contexts['mconfig'] = self._magma_service.mconfig
        contexts['loop'] = self._magma_service.loop
        contexts['service_manager'] = self

        sessiond_chan = ServiceRegistry.get_rpc_channel(
            'sessiond', ServiceRegistry.LOCAL)
        mobilityd_chan = ServiceRegistry.get_rpc_channel(
            'mobilityd', ServiceRegistry.LOCAL)
        contexts['rpc_stubs'] = {
            'mobilityd': MobilityServiceStub(mobilityd_chan),
            'sessiond': LocalSessionManagerStub(sessiond_chan),
        }

        # Instantiate and schedule apps
        for app in manager.instantiate_apps(**contexts):
            # Wrap the eventlet in asyncio so it will stop when the loop is
            # stopped
            future = aioeventlet.wrap_greenthread(app,
                                                  self._magma_service.loop)

            # Schedule the eventlet for evaluation in service loop
            asyncio.ensure_future(future)

        # In development mode, run server so that
        if environment.is_dev_mode():
            server_thread = of_rest_server.start(manager)
            future = aioeventlet.wrap_greenthread(server_thread,
                                                  self._magma_service.loop)
            asyncio.ensure_future(future)

    def get_table_num(self, app_name: str) -> int:
        """
        Args:
            app_name: Name of the app
        Returns:
            The app's main table number
        """
        return self._table_manager.get_table_num(app_name)

    def get_next_table_num(self, app_name: str) -> int:
        """
        Args:
            app_name: Name of the app
        Returns:
            The main table number of the next service.
            If there are no more services after the current table,
            return the EGRESS table
        """
        return self._table_manager.get_next_table_num(app_name)

    def is_app_enabled(self, app_name: str) -> bool:
        """
        Args:
             app_name: Name of the app
        Returns:
            Whether or not the app is enabled
        """
        return self._table_manager.is_app_enabled(app_name)

    def allocate_scratch_tables(self, app_name: str, count: int) -> List[int]:
        """
        Args:
            app_name:
                Each scratch table is associated with an app. This is used to
                help enforce scratch table isolation between apps.
            count: Number of scratch tables to be claimed
        Returns:
            List of scratch table numbers
        Raises:
            TableNumException if there are no more available tables
        """
        return self._table_manager.allocate_scratch_tables(app_name, count)

    def get_scratch_table_nums(self, app_name: str) -> List[int]:
        """
        Returns the scratch tables claimed by the given app.
        """
        return self._table_manager.get_scratch_table_nums(app_name)

    def get_all_table_assignments(self):
        """
        Returns: OrderedDict of app name to tables mapping, ordered by main
        table number, and app name.
        """
        return self._table_manager.get_all_table_assignments()
