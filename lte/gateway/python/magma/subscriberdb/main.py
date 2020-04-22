"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""
import asyncio
import logging

from magma.common.service import MagmaService
from magma.common.streamer import StreamerClient
from .processor import Processor
from .protocols.diameter.application import base, s6a
from .protocols.diameter.server import S6aServer
from .rpc_servicer import SubscriberDBRpcServicer
from .subscription_profile import get_default_sub_profile
from .streamer_callback import SubscriberDBStreamerCallback
from .store.sqlite import SqliteStore
from .protocols.s6a_proxy_servicer import S6aProxyRpcServicer
from lte.protos.mconfig import mconfigs_pb2


def main():
    """ main() for subscriberdb """
    service = MagmaService('subscriberdb', mconfigs_pb2.SubscriberDB())

    # Initialize a store to keep all subscriber data.
    store = SqliteStore(service.config['db_path'], loop=service.loop)

    # Initialize the processor
    processor = Processor(store,
                          get_default_sub_profile(service),
                          service.mconfig.sub_profiles,
                          service.mconfig.lte_auth_op,
                          service.mconfig.lte_auth_amf)

    # Add all servicers to the server
    subscriberdb_servicer = SubscriberDBRpcServicer(store)
    subscriberdb_servicer.add_to_server(service.rpc_server)


    # Start a background thread to stream updates from the cloud
    if service.config['enable_streaming']:
        callback = SubscriberDBStreamerCallback(store, service.loop)
        stream = StreamerClient({"subscriberdb": callback}, service.loop)
        stream.start()
    else:
        logging.info('enable_streaming set to False. Streamer disabled!')

    # Wait until the datastore is populated by addition or resync before
    # listening for clients.
    def serve():
        # Waiting for subscribers to be added to store
        yield from store.on_ready()

        if service.config['s6a_over_grpc']:
            s6a_proxy_servicer = S6aProxyRpcServicer(processor)
            s6a_proxy_servicer.add_to_server(service.rpc_server)
        else:
            base_manager = base.BaseApplication(
                service.config['mme_realm'],
                service.config['mme_host_name'],
                service.config['mme_host_address'],
            )
            s6a_manager = _get_s6a_manager(service, processor)
            base_manager.register(s6a_manager)

            # Setup the Diameter/s6a MME
            s6a_server = service.loop.create_server(
                lambda: S6aServer(base_manager,
                              s6a_manager,
                              service.config['mme_realm'],
                              service.config['mme_host_name'],
                              loop=service.loop),
                service.config['host_address'], service.config['mme_port'])
            asyncio.ensure_future(s6a_server, loop=service.loop)
    asyncio.ensure_future(serve(), loop=service.loop)

    # Run the service loop
    service.run()

    # Cleanup the service
    service.close()


def _get_s6a_manager(service, processor):
    return s6a.S6AApplication(
        processor,
        service.config['mme_realm'],
        service.config['mme_host_name'],
        service.config['mme_host_address'],
        service.loop
    )


if __name__ == "__main__":
    main()
