"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""
from collections import namedtuple
from ryu.ofproto.ofproto_v1_4 import OFPP_LOCAL

from .base import MagmaController
from magma.pipelined.app.li_mirror import LIMirrorController
from magma.pipelined.openflow import flows
from magma.pipelined.bridge_util import BridgeTools
from magma.pipelined.openflow.magma_match import MagmaMatch
from magma.pipelined.openflow.registers import load_direction, Direction, \
    PASSTHROUGH_REG_VAL

from ryu.lib.packet import ether_types

# ingress and egress service names -- used by other controllers
INGRESS = "ingress"
EGRESS = "egress"
PHYSICAL_TO_LOGICAL = "middle"


class InOutController(MagmaController):
    """
    A controller that sets up an openflow pipeline for Magma.

    The EPC controls table 0 which is the first table every packet touches.
    This controller owns the ingress and output portions of the pipeline, the
    first table a packet hits after the EPC controller's table 0 and the last
    table a packet hits before exiting the pipeline.
    """

    APP_NAME = "inout"

    InOutConfig = namedtuple(
        'InOutConfig',
        ['gtp_port', 'uplink_port_name', 'mtr_ip', 'mtr_port', 'li_port_name'],
    )

    def __init__(self, *args, **kwargs):
        super(InOutController, self).__init__(*args, **kwargs)
        self.config = self._get_config(kwargs['config'])
        self._uplink_port = OFPP_LOCAL
        self._li_port = None
        #TODO Alex do we want this to be cofigurable from swagger?
        if self.config.mtr_ip:
            self._mtr_service_enabled = True
        else:
            self._mtr_service_enabled = False
        if (self.config.uplink_port_name):
            self._uplink_port = BridgeTools.get_ofport(self.config.uplink_port_name)
        if (self._service_manager.is_app_enabled(LIMirrorController.APP_NAME)
                and self.config.li_port_name):
            self._li_port = BridgeTools.get_ofport(self.config.li_port_name)
            self._li_table = self._service_manager.get_table_num(
                LIMirrorController.APP_NAME)
        self._ingress_tbl_num = self._service_manager.get_table_num(INGRESS)
        self._midle_tbl_num = \
            self._service_manager.get_table_num(PHYSICAL_TO_LOGICAL)
        self._egress_tbl_num = self._service_manager.get_table_num(EGRESS)

    def _get_config(self, config_dict):
        port_name = None
        mtr_ip = None
        mtr_port = None
        li_port_name = None
        if 'ovs_uplink_port_name' in config_dict:
            port_name = config_dict['ovs_uplink_port_name']

        if 'mtr_ip' in config_dict:
            self._mtr_service_enabled = True
            mtr_ip = config_dict['mtr_ip']
            mtr_port = config_dict['ovs_mtr_port_number']
        if 'li_local_iface' in config_dict:
            li_port_name = config_dict['li_local_iface']

        return self.InOutConfig(
            gtp_port=config_dict['ovs_gtp_port_number'],
            uplink_port_name=port_name,
            mtr_ip=mtr_ip,
            mtr_port=mtr_port,
            li_port_name=li_port_name
        )

    def initialize_on_connect(self, datapath):
        self.delete_all_flows(datapath)
        self._install_default_egress_flows(datapath)
        self._install_default_ingress_flows(datapath)
        self._install_default_middle_flows(datapath)

    def cleanup_on_disconnect(self, datapath):
        self.delete_all_flows(datapath)

    def delete_all_flows(self, datapath):
        flows.delete_all_flows_from_table(datapath, self._ingress_tbl_num)
        flows.delete_all_flows_from_table(datapath, self._midle_tbl_num)
        flows.delete_all_flows_from_table(datapath, self._egress_tbl_num)

    def _install_default_middle_flows(self, dp):
        """
        Egress table is the last table that a packet touches in the pipeline.
        Output downlink traffic to gtp port, uplink trafic to LOCAL

        Raises:
            MagmaOFError if any of the default flows fail to install.
        """
        next_tbl = self._service_manager.get_next_table_num(PHYSICAL_TO_LOGICAL)

        # Allow passthrough pkts(skip enforcement and send to egress table)
        ps_match = MagmaMatch(passthrough=PASSTHROUGH_REG_VAL)
        flows.add_resubmit_next_service_flow(dp, self._midle_tbl_num, ps_match,
            actions=[], priority=flows.PASSTHROUGH_PRIORITY,
            resubmit_table=self._egress_tbl_num)

        match = MagmaMatch()
        flows.add_resubmit_next_service_flow(dp,
            self._midle_tbl_num, match,
            actions=[], priority=flows.DEFAULT_PRIORITY,
            resubmit_table=next_tbl)

        if self._mtr_service_enabled:
            match = MagmaMatch(eth_type=ether_types.ETH_TYPE_IP,
                               ipv4_dst=self.config.mtr_ip)
            flows.add_output_flow(dp,
                self._midle_tbl_num, match,
                [], priority=flows.UE_FLOW_PRIORITY,
                output_port=self.config.mtr_port)

    def _install_default_egress_flows(self, dp):
        """
        Egress table is the last table that a packet touches in the pipeline.
        Output downlink traffic to gtp port, uplink trafic to LOCAL

        Raises:
            MagmaOFError if any of the default flows fail to install.
        """
        downlink_match = MagmaMatch(direction=Direction.IN)
        flows.add_output_flow(dp, self._egress_tbl_num, downlink_match, [],
                              output_port=self.config.gtp_port)

        uplink_match = MagmaMatch(direction=Direction.OUT)
        flows.add_output_flow(dp, self._egress_tbl_num, uplink_match, [],
                              output_port=self._uplink_port)

    def _install_default_ingress_flows(self, dp):
        """
        Sets up the ingress table, the first step in the packet processing
        pipeline.

        This sets up flow rules to annotate packets with a metadata bit
        indicating the direction. Incoming packets are defined as packets
        originating from the LOCAL port, outgoing packets are defined as
        packets originating from the gtp port.

        All other packets bypass the pipeline.

        Note that the ingress rules do *not* install any flows that cause
        PacketIns (i.e., sends packets to the controller).

        Raises:
            MagmaOFError if any of the default flows fail to install.
        """
        parser = dp.ofproto_parser
        next_table = self._service_manager.get_next_table_num(INGRESS)

        # set traffic direction bits
        # set a direction bit for outgoing (pn -> inet) traffic.
        match = MagmaMatch(in_port=self.config.gtp_port)
        actions = [load_direction(parser, Direction.OUT)]
        flows.add_resubmit_next_service_flow(dp, self._ingress_tbl_num, match,
                                             actions=actions,
                                             priority=flows.DEFAULT_PRIORITY,
                                             resubmit_table=next_table)

        # set a direction bit for incoming (internet -> UE) traffic.
        match = MagmaMatch(in_port=OFPP_LOCAL)
        actions = [load_direction(parser, Direction.IN)]
        flows.add_resubmit_next_service_flow(dp, self._ingress_tbl_num, match,
                                             actions=actions,
                                             priority=flows.DEFAULT_PRIORITY,
                                             resubmit_table=next_table)

        # set a direction bit for incoming (internet -> UE) traffic.
        match = MagmaMatch(in_port=self._uplink_port)
        actions = [load_direction(parser, Direction.IN)]
        flows.add_resubmit_next_service_flow(dp, self._ingress_tbl_num, match,
                                             actions=actions,
                                             priority=flows.DEFAULT_PRIORITY,
                                             resubmit_table=next_table)

        # Send RADIUS requests directly to li table
        if self._li_port:
            match = MagmaMatch(in_port=self._li_port)
            actions = [load_direction(parser, Direction.IN)]
            flows.add_resubmit_next_service_flow(dp, self._ingress_tbl_num,
                match, actions=actions, priority=flows.DEFAULT_PRIORITY,
                resubmit_table=self._li_table)

        # set a direction bit for incoming (mtr -> UE) traffic.
        if self._mtr_service_enabled:
            match = MagmaMatch(in_port=self.config.mtr_port)
            actions = [load_direction(parser, Direction.IN)]
            flows.add_resubmit_next_service_flow(dp, self._ingress_tbl_num,
                match, actions=actions, priority=flows.DEFAULT_PRIORITY,
                resubmit_table=next_table)
