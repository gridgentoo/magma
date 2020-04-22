/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */
/*! \file gtpv1u.h
* \brief
* \author Sebastien ROUX, Lionel Gauthier
* \company Eurecom
* \email: lionel.gauthier@eurecom.fr
*/

#ifndef FILE_GTPV1_U_SEEN
#define FILE_GTPV1_U_SEEN

#include <arpa/inet.h>
#include <net/if.h>
#include "sgw_ie_defs.h"

#define GTPU_HEADER_OVERHEAD_MAX 64

/*
 * downlink flow description for a dedicated bearer
 */

#define SRC_IPV4 0x1
#define DST_IPV4 0x2
#define TCP_SRC_PORT 0x4
#define TCP_DST_PORT 0x8
#define UDP_SRC_PORT 0x10
#define UDP_DST_PORT 0x20
#define IP_PROTO 0x40

// This is the default precedence value for flow rules.
// A flow rule with precedence value 0 takes precedence over
// all other flows with higher value. Flow rules that use
// the default precedence have the lowest priority.
#define DEFAULT_PRECEDENCE 65535
// This is the maximum priority an OVS rule can take.
// A high priority value takes precedence over the lower value.
// For equal priority values, the behavior is undefined, but
// it is not atypical to see a previously installed rule to take
// precedence over latter rules.
#define MAX_PRIORITY 65535

struct ipv4flow_dl {
  struct in_addr dst_ip;
  struct in_addr src_ip;
  uint32_t set_params;
  uint16_t tcp_dst_port;
  uint16_t tcp_src_port;
  uint16_t udp_dst_port;
  uint16_t udp_src_port;
  uint8_t ip_proto;
};

/*
 * This structure defines the management hooks for GTP tunnels and paging support.
 * The following hooks can be defined; unless noted otherwise, they are
 * optional and can be filled with a null pointer.
 *
 * int (*init)(struct in_addr *ue_net, uint32_t mask,
 *             int mtu, int *fd0, int *fd1u);
 *     This function is called when initializing GTP network device. How to use
 *     these input parameters are defined by the actual function implementations.
 *         @ue_net: subnet assigned to UEs
 *         @mask: network mask for the UE subnet
 *         @mtu: MTU for the GTP network device.
 *         @fd0: socket file descriptor for GTPv0.
 *         @fd1u: socket file descriptor for GTPv1u.
 *
 * int (*uninit)(void);
 *     This function is called to destroy GTP network device.
 *
 * int (*reset)(void);
 *     This function is called to reset the GTP network device to clean state.
 *
 * int (*add_tunnel)(struct in_addr ue, struct in_addr enb, uint32_t i_tei, uint32_t o_tei, Imsi_t imsi);
 *     Add a gtp tunnel.
 *         @ue: UE IP address
 *         @enb: eNB IP address
 *         @i_tei: RX GTP Tunnel ID
 *         @o_tei: TX GTP Tunnel ID.
 *         @imsi: UE IMSI
 *         @flow_dl: Downlink flow rule
 *         @flow_precedence_dl: Downlink flow rule precedence
 *
 * int (*del_tunnel)(uint32_t i_tei, uint32_t o_tei);
 *     Delete a gtp tunnel.
 *         @ue: UE IP address
 *         @i_tei: RX GTP Tunnel ID
 *         @o_tei: TX GTP Tunnel ID.
 *
 * int (*discard_data_on_tunnel)(struct in_addr ue, uint32_t i_tei);
 *         @ue: UE IP address
 *         @i_tei: RX GTP Tunnel ID
 *
 * int (*forward_data_on_tunnel)(struct in_addr ue, uint32_t i_tei);
 *         @ue: UE IP address
 *         @i_tei: RX GTP Tunnel ID
 *         @flow_dl: Downlink flow rule
 *         @flow_precedence_dl: Downlink flow rule precedence
 *
 * int (*add_paging_rule)(struct in_addr ue);
 *        @ue: UE IP address
 */
struct gtp_tunnel_ops {
  int (*init)(
    struct in_addr* ue_net,
    uint32_t mask,
    int mtu,
    int* fd0,
    int* fd1u,
    bool persist_state);
  int (*uninit)(void);
  int (*reset)(void);
  int (*add_tunnel)(
    struct in_addr ue,
    struct in_addr enb,
    uint32_t i_tei,
    uint32_t o_tei,
    Imsi_t imsi,
    struct ipv4flow_dl *flow_dl,
    uint32_t flow_precedence_dl);
  int (*del_tunnel)(struct in_addr ue, uint32_t i_tei,
      uint32_t o_tei, struct ipv4flow_dl *flow_dl);
  int (*discard_data_on_tunnel)(struct in_addr ue,
      uint32_t i_tei, struct ipv4flow_dl *flow_dl);
  int (*forward_data_on_tunnel)(
    struct in_addr ue,
    uint32_t i_tei,
    struct ipv4flow_dl* flow_dl,
    uint32_t flow_precedence_dl);
  int (*add_paging_rule)(struct in_addr ue);
  int (*delete_paging_rule)(struct in_addr ue);
};

#if ENABLE_OPENFLOW
const struct gtp_tunnel_ops *gtp_tunnel_ops_init_openflow(void);
#else
const struct gtp_tunnel_ops *gtp_tunnel_ops_init_libgtpnl(void);
#endif

#endif /* FILE_GTPV1_U_SEEN */
