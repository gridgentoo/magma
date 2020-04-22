"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

from prometheus_client import Counter


# Counters for Diameter/S6a application
S6A_AUTH_SUCCESS_TOTAL = Counter('s6a_auth_success',
                                  'Total successful S6a auth requests')
S6A_AUTH_FAILURE_TOTAL = Counter('s6a_auth_failure',
                                  'Total failed S6a auth requests with reason', ['code'])
S6A_LUR_TOTAL = Counter('s6a_location_update',
                         'Total S6a location update requests')

DIAMETER_AUTHENTICATION_REJECTED = 4001
DIAMETER_ERROR_USER_UNKNOWN = 5001
DIAMETER_AUTHORIZATION_REJECTED = 5003
# Counters for Diameter base application
DIAMETER_CEX_TOTAL = Counter('diameter_capabilities_exchange',
                             'Total Diameter capabilities exchange requests')
DIAMETER_WATCHDOG_TOTAL = Counter('diameter_watchdog',
                                  'Total Diameter watchdog requests')
DIAMETER_DISCONECT_TOTAL = Counter('diameter_disconnect',
                                   'Total Diameter disconnect requests')
