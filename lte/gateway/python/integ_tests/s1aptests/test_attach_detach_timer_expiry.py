"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest
import time

import gpp_types
import s1ap_types
import s1ap_wrapper


class TestDetachTimerExpiryHandle(unittest.TestCase):

    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_detach_timer_expiry_handle(self):
        """
        Test Detach timer expiry handling
        currenly manually need to verify through mme log
        """
        self._s1ap_wrapper.configUEDevice(1)
        req = self._s1ap_wrapper.ue_req
        ue_id = req.ue_id
        print("************************* Running End to End attach for UE id ",
              ue_id)
        # Now actually complete the attach
        self._s1ap_wrapper._s1_util.attach(
            ue_id, s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
            s1ap_types.ueAttachAccept_t)

        # Wait on EMM Information from MME
        self._s1ap_wrapper._s1_util.receive_emm_info()

        # Delay to ensure S1APTester sends attach complete before sending UE
        # context release
        time.sleep(0.5)

        print("************************* Sending UE context release request ",
              "for UE id ", ue_id)
        # Send UE context release request to move UE to idle mode
        req = s1ap_types.ueCntxtRelReq_t()
        req.ue_Id = ue_id
        req.cause.causeVal = gpp_types.CauseRadioNetwork.USER_INACTIVITY.value
        self._s1ap_wrapper.s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_CNTXT_REL_REQUEST, req)
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_CTX_REL_IND.value)

        print("************************* Sending Service request for UE id ",
              ue_id)
        # To expires detach timer, first ensure
        # mobile reachability timer expired or not and then
        # delay sending initial ue message(service req) by detach timer value.
        # For testing purpose manually reduce the mobile reachability time to
        # 1minute in mme.conf.template file
        # DETACH TIMER VALUE = mobile reachability timer value + delta value

        time.sleep(740)

        # Send service request to reconnect UE
        req = s1ap_types.ueserviceReq_t()
        req.ue_Id = ue_id
        req.ueMtmsi = s1ap_types.ueMtmsi_t()
        req.ueMtmsi.pres = False
        req.rrcCause = s1ap_types.Rrc_Cause.TFW_MO_DATA.value
        self._s1ap_wrapper.s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_SERVICE_REQUEST, req)
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_SERVICE_REJECT_IND.value)
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_CTX_REL_IND.value)

        time.sleep(0.5)


if __name__ == "__main__":
    unittest.main()
