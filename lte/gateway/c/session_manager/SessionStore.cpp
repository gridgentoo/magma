/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#include "SessionState.h"
#include "SessionStore.h"
#include "StoredState.h"
#include "magma_logging.h"

namespace magma {
namespace lte {

SessionStore::SessionStore(std::shared_ptr<StaticRuleStore> rule_store)
    : rule_store_(rule_store),
      store_client_(std::make_shared<MemoryStoreClient>(rule_store)) {}

SessionStore::SessionStore(
    std::shared_ptr<StaticRuleStore> rule_store,
    std::shared_ptr<RedisStoreClient> store_client)
    : rule_store_(rule_store), store_client_(store_client) {}

SessionMap SessionStore::read_sessions(const SessionRead& req) {
  return store_client_->read_sessions(req);
}

SessionMap SessionStore::read_all_sessions() {
  return store_client_->read_all_sessions();
}

SessionMap SessionStore::read_sessions_for_reporting(const SessionRead& req) {
  auto session_map   = store_client_->read_sessions(req);
  auto session_map_2 = store_client_->read_sessions(req);
  // For all sessions of the subscriber, increment the request numbers
  for (const std::string& imsi : req) {
    for (auto& session : session_map_2[imsi]) {
      session->increment_request_number(session->get_credit_key_count());
    }
  }
  store_client_->write_sessions(std::move(session_map_2));
  return session_map;
}

SessionMap SessionStore::read_sessions_for_deletion(const SessionRead& req) {
  auto session_map   = store_client_->read_sessions(req);
  auto session_map_2 = store_client_->read_sessions(req);
  // For all sessions of the subscriber, increment the request numbers
  for (const std::string& imsi : req) {
    for (auto& session : session_map_2[imsi]) {
      session->increment_request_number(1);
    }
  }
  store_client_->write_sessions(std::move(session_map_2));
  return session_map;
}

bool SessionStore::create_sessions(
    const std::string& subscriber_id,
    std::vector<std::unique_ptr<SessionState>> sessions) {
  auto session_map           = SessionMap{};
  session_map[subscriber_id] = std::move(sessions);
  store_client_->write_sessions(std::move(session_map));
  return true;
}

bool SessionStore::update_sessions(const SessionUpdate& update_criteria) {
  MLOG(MDEBUG) << "Running update_sessions";
  // Read the current state
  auto subscriber_ids = std::set<std::string>{};
  for (const auto& it : update_criteria) {
    subscriber_ids.insert(it.first);
  }
  auto session_map = store_client_->read_sessions(subscriber_ids);

  // Now attempt to modify the state
  for (auto& it : session_map) {
    auto it2 = it.second.begin();
    while (it2 != it.second.end()) {
      auto updates = update_criteria.find(it.first)->second;
      if (updates.find((*it2)->get_session_id()) != updates.end()) {
        auto update = updates[(*it2)->get_session_id()];
        if (!merge_into_session(*it2, update)) {
          return false;
        }
        if (update.is_session_ended) {
          // TODO: Instead of deleting from session_map, mark as ended and
          //       no longer mark on read
          it2 = it.second.erase(it2);
          continue;
        }
      }
      ++it2;
    }
  }
  return store_client_->write_sessions(std::move(session_map));
}

bool SessionStore::merge_into_session(
    std::unique_ptr<SessionState>& session,
    const SessionStateUpdateCriteria& update_criteria) {
  // Config
  if (update_criteria.is_config_updated) {
    session->unmarshal_config(update_criteria.updated_config);
  }

  // Static rules
  auto uc = get_default_update_criteria();
  for (const auto& rule_id : update_criteria.static_rules_to_install) {
    if (session->is_static_rule_installed(rule_id)) {
      MLOG(MERROR) << "Failed to merge: " << session->get_session_id()
                   << " because static rule already installed: " << rule_id
                   << std::endl;
      return false;
    }
    session->activate_static_rule(rule_id, uc);
  }
  for (const auto& rule_id : update_criteria.static_rules_to_uninstall) {
    if (!session->is_static_rule_installed(rule_id)) {
      MLOG(MERROR) << "Failed to merge: " << session->get_session_id()
                   << " because static rule already uninstalled: " << rule_id
                   << std::endl;
      return false;
    }
    session->deactivate_static_rule(rule_id, uc);
  }

  // Dynamic rules
  for (const auto& rule : update_criteria.dynamic_rules_to_install) {
    if (session->is_dynamic_rule_installed(rule.id())) {
      MLOG(MERROR) << "Failed to merge: " << session->get_session_id()
                   << " because dynamic rule already installed: " << rule.id()
                   << std::endl;
      return false;
    }
    session->insert_dynamic_rule(rule, uc);
  }
  PolicyRule* _ = {};
  for (const auto& rule_id : update_criteria.dynamic_rules_to_uninstall) {
    if (!session->is_dynamic_rule_installed(rule_id)) {
      MLOG(MERROR) << "Failed to merge: " << session->get_session_id()
                   << " because dynamic rule already uninstalled: " << rule_id
                   << std::endl;
      return false;
    }
    session->remove_dynamic_rule(rule_id, _, uc);
  }

  // Charging credit
  for (const auto& it : update_criteria.charging_credit_map) {
    auto key           = it.first;
    auto credit_update = it.second;
    session->get_charging_pool().merge_credit_update(key, credit_update);
  }
  for (const auto& it : update_criteria.charging_credit_to_install) {
    auto key           = it.first;
    auto stored_credit = it.second;
    auto uc            = get_default_update_criteria();
    session->get_charging_pool().add_credit(
        key, SessionCredit::unmarshal(stored_credit, CHARGING), uc);
  }

  // Monitoring credit
  for (const auto& it : update_criteria.monitor_credit_map) {
    auto key           = it.first;
    auto credit_update = it.second;
    session->get_monitor_pool().merge_credit_update(key, credit_update);
  }
  for (const auto& it : update_criteria.monitor_credit_to_install) {
    auto key            = it.first;
    auto stored_monitor = it.second;
    auto uc             = get_default_update_criteria();
    session->get_monitor_pool().add_monitor(
        key, UsageMonitoringCreditPool::unmarshal_monitor(stored_monitor), uc);
  }
  return true;
}

SessionUpdate SessionStore::get_default_session_update(
    SessionMap& session_map) {
  SessionUpdate update = {};
  for (const auto& session_pair : session_map) {
    for (const auto& session : session_pair.second) {
      update[session_pair.first][session->get_session_id()] =
          get_default_update_criteria();
    }
  }
  return update;
}

}  // namespace lte
}  // namespace magma
