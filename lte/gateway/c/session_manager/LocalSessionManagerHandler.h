/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */
#pragma once

#include <functional>

#include <grpc++/grpc++.h>
#include <lte/protos/session_manager.grpc.pb.h>

#include "LocalEnforcer.h"
#include "SessionID.h"
#include "SessionReporter.h"
#include "SessionStore.h"

using grpc::Server;
using grpc::ServerContext;
using grpc::Status;

namespace magma {
using namespace orc8r;

class LocalSessionManagerHandler {
 public:
  virtual ~LocalSessionManagerHandler() {}

  /**
   * Report flow stats from pipelined and track the usage per rule
   */
  virtual void ReportRuleStats(
      ServerContext* context, const RuleRecordTable* request,
      std::function<void(Status, Void)> response_callback) = 0;

  /**
   * Create a new session, initializing credit monitoring and requesting credit
   * from the cloud
   */
  virtual void CreateSession(
      ServerContext* context, const LocalCreateSessionRequest* request,
      std::function<void(Status, LocalCreateSessionResponse)>
          response_callback) = 0;

  /**
   * Terminate a session, untracking credit and terminating in the cloud
   */
  virtual void EndSession(
      ServerContext* context, const LocalEndSessionRequest* request,
      std::function<void(Status, LocalEndSessionResponse)>
          response_callback) = 0;
};

/**
 * LocalSessionManagerHandler processes proxied gRPC requests to the session
 * manager. The handler uses a monitor and reporter to keep track of state
 * and report to the cloud, respectively
 */
class LocalSessionManagerHandlerImpl : public LocalSessionManagerHandler {
 public:
  LocalSessionManagerHandlerImpl(
      std::shared_ptr<LocalEnforcer> monitor, SessionReporter* reporter,
      std::shared_ptr<AsyncDirectorydClient> directoryd_client,
      SessionStore& session_store);
  ~LocalSessionManagerHandlerImpl() {}
  /**
   * Report flow stats from pipelined and track the usage per rule
   */
  void ReportRuleStats(
      ServerContext* context, const RuleRecordTable* request,
      std::function<void(Status, Void)> response_callback);

  /**
   * Create a new session, initializing credit monitoring and requesting credit
   * from the cloud
   */
  void CreateSession(
      ServerContext* context, const LocalCreateSessionRequest* request,
      std::function<void(Status, LocalCreateSessionResponse)>
          response_callback);

  /**
   * Terminate a session, untracking credit and terminating in the cloud
   */
  void EndSession(
      ServerContext* context, const LocalEndSessionRequest* request,
      std::function<void(Status, LocalEndSessionResponse)> response_callback);

 private:
  SessionStore& session_store_;
  std::shared_ptr<LocalEnforcer> enforcer_;
  SessionReporter* reporter_;
  std::shared_ptr<AsyncDirectorydClient> directoryd_client_;
  SessionIDGenerator id_gen_;
  uint64_t current_epoch_;
  uint64_t reported_epoch_;
  std::chrono::seconds retry_timeout_;
  static const std::string hex_digit_;

 private:
  void check_usage_for_reporting(
      SessionMap session_map, SessionUpdate& session_update);
  bool is_pipelined_restarted();
  bool restart_pipelined(const std::uint64_t& epoch);

  void end_session(
      SessionMap& session_map, const LocalEndSessionRequest& request,
      std::function<void(Status, LocalEndSessionResponse)> response_callback);

  std::string convert_mac_addr_to_str(const std::string& mac_addr);

  void add_session_to_directory_record(
      const std::string& imsi, const std::string& session_id);

  /**
   * Send session creation request to the CentralSessionController.
   * If it is successful, create a session in session_map, and respond to
   * gRPC caller.
   */
  void send_create_session(
      SessionMap& session_map, const CreateSessionRequest& request,
      const std::string& imsi, const std::string& sid, const SessionConfig& cfg,
      std::function<void(grpc::Status, LocalCreateSessionResponse)>
          response_callback);

  void handle_setup_callback(
      const std::uint64_t& epoch, Status status, SetupFlowsResult resp);

  SessionConfig build_session_config(const LocalCreateSessionRequest& request);

  void recycle_session(
      SessionMap& session_map, const LocalCreateSessionRequest& request,
      const std::string& imsi, const std::string& sid,
      const std::string& core_sid, SessionConfig cfg, const bool is_wifi,
      std::function<void(Status, LocalCreateSessionResponse)>
          response_callback);

  /**
   * Get the most recently written state of sessions for Creation
   * Does not get any other sessions.
   *
   * NOTE: Call only from the main EventBase thread, otherwise there will
   *       be undefined behavior.
   */
  SessionMap get_sessions_for_creation(
      const LocalCreateSessionRequest& request);

  /**
   * Get the most recently written state of sessions for reporting usage.
   * Does not get sessions that are not required for reporting.
   *
   * NOTE: Call only from the main EventBase thread, otherwise there will
   *       be undefined behavior.
   */
  SessionMap get_sessions_for_reporting(const RuleRecordTable& request);

  /**
   * Get the most recently written state of the session that is to be deleted.
   * Does not get any other sessions.
   *
   * NOTE: Call only from the main EventBase thread, otherwise there will
   *       be undefined behavior.
   */
  SessionMap get_sessions_for_deletion(const LocalEndSessionRequest& request);
};

}  // namespace magma
