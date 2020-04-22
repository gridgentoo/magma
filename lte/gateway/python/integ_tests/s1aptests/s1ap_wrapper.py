""""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import os
import time
import ctypes

import s1ap_types
from integ_tests.common.magmad_client import MagmadServiceGrpc

# from integ_tests.cloud.cloud_manager import CloudManager
from integ_tests.common.mobility_service_client import MobilityServiceGrpc
from integ_tests.common.service303_utils import GatewayServicesUtil
from integ_tests.common.subscriber_db_client import SubscriberDbGrpc
from integ_tests.s1aptests.s1ap_utils import (
    MagmadUtil,
    MobilityUtil,
    S1ApUtil,
    SubscriberUtil,
)
from integ_tests.s1aptests.util.traffic_util import TrafficUtil


class TestWrapper(object):
    """
    Module wrapping boiler plate code for all test setups and cleanups.
    """

    TEST_IP_BLOCK = "192.168.128.0/24"
    MSX_S1_RETRY = 2

    def __init__(self):
        """
        Initialize the various classes required by the tests and setup.
        """
        self._s1_util = S1ApUtil()
        self._enBConfig()

        subscriber_client = SubscriberDbGrpc()
        mobility_client = MobilityServiceGrpc()
        magmad_client = MagmadServiceGrpc()
        self._sub_util = SubscriberUtil(subscriber_client)
        # Remove existing subscribers to start
        self._sub_util.cleanup()
        self._mobility_util = MobilityUtil(mobility_client)
        self._mobility_util.cleanup()
        self._magmad_util = MagmadUtil(magmad_client)
        # gateway tests don't require restart, just wait for healthy now
        self._gateway_services = GatewayServicesUtil()
        self.wait_gateway_healthy = True
        if not self.wait_gateway_healthy:
            self.init_s1ap_tester()

        self._configuredUes = []
        self._ue_idx = 0  # Index of UEs already used in test
        self._trf_util = TrafficUtil()

    def init_s1ap_tester(self):
        """
        Initialize the s1ap tester and the UEApp.

        Doing this separately allows initialization to occur during
        tests rather than during setup stage.
        """
        # config ip first, because cloud tests will restart gateway
        self.configIpBlock()

        self._s1setup()
        self._configUEApp()

    @property
    def _test_cloud(self):
        test_cloud = os.getenv("MAGMA_S1APTEST_USE_CLOUD") is not None
        return test_cloud

    @property
    def _test_oai_upstream(self):
        return os.getenv("TEST_OAI_UPSTREAM") is not None

    def _enBConfig(self):
        """Helper to configure the eNB"""
        # Using exaggerated prints makes the stdout easier to read.
        print("************************* Enb tester config")
        req = s1ap_types.FwNbConfigReq_t()
        req.cellId_pr.pres = True
        req.cellId_pr.cell_id = 10
        assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.ENB_CONFIG, req) == 0
        response = self._s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.ENB_CONFIG_CONFIRM.value
        res = response.cast(s1ap_types.FwNbConfigCfm_t)
        assert res.status == s1ap_types.CfgStatus.CFG_DONE.value

    def _issue_s1setup_req(self):
        """ Issue the actual setup request and get the response"""
        req = None
        assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.ENB_S1_SETUP_REQ, req) == 0
        response = self._s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.ENB_S1_SETUP_RESP.value
        return response.cast(s1ap_types.FwNbS1setupRsp_t)

    def _s1setup(self):
        """Helper to setup s1 to the EPC"""
        print("************************* S1 setup")
        res = self._issue_s1setup_req()

        retry = 0
        while retry < TestWrapper.MSX_S1_RETRY:
            if (
                res.res == s1ap_types.S1_setp_Result.S1_SETUP_FAILED.value
                and res.waitIe.pres == 1
            ):
                print(
                    "Received time to wait in S1-Setup-Failure" " message is",
                    res.waitIe.val,
                )
                time.sleep(res.waitIe.val)
                res = self._issue_s1setup_req()
                retry += 1
            else:
                # Not a failure in setup.
                break

        assert res.res == s1ap_types.S1_setp_Result.S1_SETUP_SUCCESS.value

    def _configUEApp(self):
        """ Update the internal configuration of the UEApp"""
        print("************************* UE App config")
        req = s1ap_types.ueAppConfig_t()
        req.nasProcGuardTimer_pr.pres = True
        req.nasProcGuardTimer_pr.nas_proc_guard_timer = 5
        assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.UE_APPL_CONFIG, req) == 0
        response = self._s1_util.get_response()
        assert s1ap_types.tfwCmd.UE_APP_CONFIG_COMPLETE_IND.value == response.msg_type

    def _getAddresses(self, *ues):
        """ Retrieve IP addresses for the given UEs

        Will put None for IPs in the cases where a UE has been included that
        doesn't have a cached IP (e.g. the UE has not yet been attached)

        Args:
            ues (list(s1ap_types.ueAppConfig_t)): the UEs whose IPs we want

        Returns a list of ipaddress.ip_address objects, corresponding in order
            with the input UE parameters
        """
        return [self._s1_util.get_ip(ue.ue_id) for ue in ues]

    def configIpBlock(self):
        """ Removes any existing allocated blocks, then adds the ones used for
        testing """
        print("************************* Configuring IP block")
        self._mobility_util.remove_all_ip_blocks()
        self._mobility_util.add_ip_block(self.TEST_IP_BLOCK)
        print("************************* Waiting for IP changes to propagate")
        self._mobility_util.wait_for_changes()

    def configUEDevice(self, num_ues):
        """ Configure the device on the UE side """
        reqs = self._sub_util.add_sub(num_ues=num_ues)
        for i in range(num_ues):
            print(
                "************************* UE device config for ue_id ", reqs[i].ue_id
            )
            assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.UE_CONFIG, reqs[i]) == 0
            response = self._s1_util.get_response()
            assert s1ap_types.tfwCmd.UE_CONFIG_COMPLETE_IND.value == response.msg_type
            self._configuredUes.append(reqs[i])
        self.check_gw_health_after_ue_load()

    def configAPN(self, imsi, apn_list):
        """ Configure the APN """
        self._sub_util.config_apn_data(imsi, apn_list)

    def configUEDevice_ues_same_imsi(self, num_ues):
        """ Configure the device on the UE side with same IMSI and
        having different ue-id"""
        reqs = self._sub_util.add_sub(num_ues=num_ues)
        for i in range(num_ues):
            print(
                "************************* UE device config for ue_id ", reqs[i].ue_id
            )
            assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.UE_CONFIG, reqs[i]) == 0
            response = self._s1_util.get_response()
            assert s1ap_types.tfwCmd.UE_CONFIG_COMPLETE_IND.value == response.msg_type
            self._configuredUes.append(reqs[i])
        for i in range(num_ues):
            reqs[i].ue_id = 2
            print(
                "************************* UE device config for ue_id ", reqs[i].ue_id
            )
            assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.UE_CONFIG, reqs[i]) == 0
            response = self._s1_util.get_response()
            assert s1ap_types.tfwCmd.UE_CONFIG_COMPLETE_IND.value == response.msg_type
            self._configuredUes.append(reqs[i])

        self.check_gw_health_after_ue_load()

    def configUEDevice_without_checking_gw_health(self, num_ues):
        """ Configure the device on the UE side """
        reqs = self._sub_util.add_sub(num_ues=num_ues)
        for i in range(num_ues):
            print(
                "************************* UE device config for ue_id ", reqs[i].ue_id
            )
            assert self._s1_util.issue_cmd(s1ap_types.tfwCmd.UE_CONFIG, reqs[i]) == 0
            response = self._s1_util.get_response()
            assert s1ap_types.tfwCmd.UE_CONFIG_COMPLETE_IND.value == response.msg_type
            self._configuredUes.append(reqs[i])

    def check_gw_health_after_ue_load(self):
        """ Wait for the MME only after adding entries to HSS """
        if self.wait_gateway_healthy:
            self._gateway_services.wait_for_healthy_gateway()
            self.init_s1ap_tester()
            self.wait_gateway_healthy = False

    def configDownlinkTest(self, *ues, **kwargs):
        """ Set up an downlink test, returning a TrafficTest object

        Args:
            ues (s1ap_types.ueConfig_t): the UEs to test
            kwargs: the keyword args to pass into generate_downlink_test

        Returns: a TrafficTest object, the traffic test generated based on the
            given UEs
        """
        ips = self._getAddresses(*ues)
        for ip, ue in zip(ips, ues):
            if not ip:
                raise ValueError(
                    "Encountered invalid IP for UE ID %s."
                    " Are you sure the UE is attached?" % ue
                )
        return self._trf_util.generate_traffic_test(ips, is_uplink=False, **kwargs)

    def configUplinkTest(self, *ues, **kwargs):
        """ Set up an uplink test, returning a TrafficTest object

        Args:
            ues (s1ap_types.ueConfig_t): the UEs to test
            kwargs: the keyword args to pass into generate_uplink_test

        Returns: a TrafficTest object, the traffic test generated based on the
            given UEs
        """
        ips = self._getAddresses(*ues)
        for ip, ue in zip(ips, ues):
            if not ip:
                raise ValueError(
                    "Encountered invalid IP for UE ID %s."
                    " Are you sure the UE is attached?" % ue
                )
        return self._trf_util.generate_traffic_test(ips, is_uplink=True, **kwargs)

    def get_gateway_services_util(self):
        """ Not a property, so return object is callable """
        return self._gateway_services

    @property
    def ue_req(self):
        """ Get a configured UE """
        req = self._configuredUes[self._ue_idx]
        self._ue_idx += 1
        return req

    @property
    def s1_util(self):
        return self._s1_util

    @property
    def mobility_util(self):
        return self._mobility_util

    @property
    def traffic_util(self):
        return self._trf_util

    @property
    def magmad_util(self):
        return self._magmad_util

    def cleanup(self):
        time.sleep(0.5)
        print("************************* send SCTP SHUTDOWN")
        self._s1_util.issue_cmd(s1ap_types.tfwCmd.SCTP_SHUTDOWN_REQ, None)
        self._s1_util.cleanup()
        self._sub_util.cleanup()
        self._trf_util.cleanup()
        self._mobility_util.cleanup()

        # Cloud cleanup needs to happen after cleanup for
        # subscriber util and mobility util
        # if self._test_cloud:
        #    self._cloud_manager.clean_up()

    def multiEnbConfig(self, num_of_enbs, enb_list=None):
        if enb_list is None:
            enb_list = []
        req = s1ap_types.multiEnbConfigReq_t()
        req.numOfEnbs = num_of_enbs
        # ENB Parameter column index initialization
        CELLID_COL_IDX = 0
        TAC_COL_IDX = 1
        ENBTYPE_COL_IDX = 2
        PLMNID_COL_IDX = 3
        PLMN_LENGTH_IDX = 4

        for idx1 in range(num_of_enbs):
            req.multiEnbCfgParam[idx1].cell_id = enb_list[idx1][CELLID_COL_IDX]
            req.multiEnbCfgParam[idx1].tac = enb_list[idx1][TAC_COL_IDX]
            req.multiEnbCfgParam[idx1].enbType = enb_list[idx1][ENBTYPE_COL_IDX]
            req.multiEnbCfgParam[idx1].plmn_length = enb_list[idx1][PLMN_LENGTH_IDX]
            for idx2 in range(req.multiEnbCfgParam[idx1].plmn_length):
                val = enb_list[idx1][PLMNID_COL_IDX][idx2]
                req.multiEnbCfgParam[idx1].plmn_id[idx2] = int(val)

        print("***************** Sending Multiple Enb Config Request\n")
        assert (
            self._s1_util.issue_cmd(s1ap_types.tfwCmd.MULTIPLE_ENB_CONFIG_REQ, req) == 0
        )

    def sendActDedicatedBearerAccept(self, ue_id, bearerId):
        act_ded_bearer_acc = s1ap_types.UeActDedBearCtxtAcc_t()
        act_ded_bearer_acc.ue_Id = ue_id
        act_ded_bearer_acc.bearerId = bearerId
        self._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_ACT_DED_BER_ACC, act_ded_bearer_acc
        )
        print(
            "************** Sending activate dedicated EPS bearer " "context accept\n"
        )

    def sendDeactDedicatedBearerAccept(self, ue_id, bearerId):
        deact_ded_bearer_acc = s1ap_types.UeDeActvBearCtxtAcc_t()
        deact_ded_bearer_acc.ue_Id = ue_id
        deact_ded_bearer_acc.bearerId = bearerId
        self._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_DEACTIVATE_BER_ACC, deact_ded_bearer_acc
        )
        print("************* Sending deactivate EPS bearer context accept\n")

    def sendPdnConnectivityReq(self, ue_id, apn):
        req = s1ap_types.uepdnConReq_t()
        req.ue_Id = ue_id
        # Initial Request
        req.reqType = 1
        req.pdnType_pr.pres = 1
        # PDN Type = IPv4
        req.pdnType_pr.pdn_type = 1
        req.pdnAPN_pr.pres = 1
        req.pdnAPN_pr.len = len(apn)
        req.pdnAPN_pr.pdn_apn = (ctypes.c_ubyte * 100)(
            *[ctypes.c_ubyte(ord(c)) for c in apn[:100]]
        )
        self.s1_util.issue_cmd(s1ap_types.tfwCmd.UE_PDN_CONN_REQ, req)

        print("************* Sending Standalone PDN Connectivity Request\n")
