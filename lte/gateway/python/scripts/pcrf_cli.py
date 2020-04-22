#!/usr/bin/env python3

"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import argparse
import grpc

from magma.common.rpc_utils import cloud_grpc_wrapper
from orc8r.protos.common_pb2 import Void
from feg.protos.mock_core_pb2_grpc import MockCoreConfiguratorStub


@cloud_grpc_wrapper
def send_reset(client, args):
    print("Sending reset")
    try:
        client.Reset(Void())
    except grpc.RpcError as e:
        print("gRPC failed with %s: %s" % (e.code(), e.details()))


def create_parser():
    """
    Creates the argparse parser with all the arguments.
    """
    parser = argparse.ArgumentParser(
        description='Management CLI for mock PCRF',
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)

    # Add subcommands
    subparsers = parser.add_subparsers(title='subcommands', dest='cmd')

    # Reset
    alert_ack_parser = subparsers.add_parser(
        'reset', help='Send Reset to mock PCRF hosted in FeG')
    alert_ack_parser.set_defaults(func=send_reset)

    return parser


def main():
    parser = create_parser()

    # Parse the args
    args = parser.parse_args()
    if not args.cmd:
        parser.print_usage()
        exit(1)

    # Execute the subcommand function
    args.func(args, MockCoreConfiguratorStub, 'pcrf')


if __name__ == "__main__":
    main()
