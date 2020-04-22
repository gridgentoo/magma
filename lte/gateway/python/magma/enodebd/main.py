"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

from threading import Thread
from unittest import mock
from magma.enodebd.enodeb_status import get_service_status_old, \
    get_operational_states
from magma.enodebd.state_machines.enb_acs_manager import StateMachineManager
from magma.enodebd.logger import EnodebdLogger as logger
from .rpc_servicer import EnodebdRpcServicer
from .stats_manager import StatsManager
from .tr069.server import tr069_server
from .enodebd_iptables_rules import set_enodebd_iptables_rule
from magma.common.service import MagmaService
from orc8r.protos.service303_pb2 import State
from lte.protos.mconfig import mconfigs_pb2
from typing import List


def get_context(ip: str):
    with mock.patch('spyne.server.wsgi.WsgiApplication') as MockTransport:
        MockTransport.req_env = {"REMOTE_ADDR": ip}
        with mock.patch('spyne.server.wsgi.WsgiMethodContext') as MockContext:
            MockContext.transport = MockTransport
            return MockContext

def main():
    """
    Top-level function for enodebd
    """
    service = MagmaService('enodebd', mconfigs_pb2.EnodebD())
    logger.init()

    # State machine manager for tracking multiple connected eNB devices.
    state_machine_manager = StateMachineManager(service)

    # Statistics manager
    stats_mgr = StatsManager(state_machine_manager)
    stats_mgr.run()

    # Start TR-069 thread
    server_thread = Thread(target=tr069_server,
                           args=(state_machine_manager, ),
                           daemon=True)
    server_thread.start()

    # Add all servicers to the server
    enodebd_servicer = EnodebdRpcServicer(state_machine_manager)
    enodebd_servicer.add_to_server(service.rpc_server)

    # Register function to get service status
    def get_enodebd_status():
        return get_service_status_old(state_machine_manager)
    service.register_get_status_callback(get_enodebd_status)

    # Register a callback function for GetOperationalStates service303 function
    def get_enodeb_operational_states() -> List[State]:
        return get_operational_states(state_machine_manager)
    service.register_operational_states_callback(get_enodeb_operational_states)

    # Set eNodeBD iptables rules due to exposing public IP to eNodeB
    service.loop.create_task(set_enodebd_iptables_rule())

    # Run the service loop
    service.run()

    # Cleanup the service
    service.close()


def call_repeatedly(loop, interval, function, *args, **kwargs):
    """
    Wrapper function to schedule function periodically
    """
    # Schedule next call
    loop.call_later(interval, call_repeatedly, loop, interval, function,
                    *args, **kwargs)
    # Call function
    function(*args, **kwargs)

if __name__ == "__main__":
    main()
