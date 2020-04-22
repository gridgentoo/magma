/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#pragma once

#include "CreditKey.h"
#include "RuleStore.h"
#include "SessionCredit.h"
#include "StoredState.h"
#include <unordered_map>

namespace magma {

/**
 * CreditPool is an interface that defines a group of credits to track. It is
 * keyed by some type and requires some update response type to receive credit.
 */
template <typename KeyType, typename UpdateResponseType,
          typename UpdateRequestType>
class CreditPool {
public:
  /**
   * add_used_credit adds usage to a specific credit
   */
  virtual bool add_used_credit(const KeyType &key, uint64_t used_tx,
                               uint64_t used_rx,
                               SessionStateUpdateCriteria &update_criteria) = 0;

  /**
   * reset_reporting_credit resets the credit state machine by clearing any
   * credit that was in the reporting state
   */
  virtual bool
  reset_reporting_credit(const KeyType &key,
                         SessionStateUpdateCriteria &update_criteria) = 0;

  /**
   * get_updates gets any usage updates required by the credits in the pool
   */
  virtual void
  get_updates(std::string imsi, std::string ip_addr,
              StaticRuleStore &static_rules, DynamicRuleStore *dynamic_rules,
              std::vector<UpdateRequestType> *updates_out,
              std::vector<std::unique_ptr<ServiceAction>> *actions_out,
              SessionStateUpdateCriteria &update_criteria,
              const bool force_update) = 0;

  /**
   * get_termination_updates gets updates from all credits in the pool at the
   * time of termination
   */
  virtual bool
  get_termination_updates(SessionTerminateRequest *termination_out,
                          SessionStateUpdateCriteria &update_criteria) = 0;

  /**
   * receive_credit adds allowed credit from the cloud
   */
  virtual bool receive_credit(const UpdateResponseType &update,
                              SessionStateUpdateCriteria &update_criteria) = 0;

  /**
   * get_credit is a helper function to return the bytes in a credit bucket
   */
  virtual uint64_t get_credit(const KeyType &key, Bucket bucket) const = 0;

  virtual SessionCreditUpdateCriteria *
  get_credit_update(const KeyType &key,
                    SessionStateUpdateCriteria &update_criteria) = 0;

  /**
   * Updates either the Monitor or SessionCredit using the update criteria
   */
  virtual void
  merge_credit_update(const KeyType &key,
                      SessionCreditUpdateCriteria &credit_update) = 0;

  virtual uint32_t get_credit_key_count() const = 0;
};

/**
 * ChargingCreditPool manages a pool of credits for OCS-based charging. It is
 * keyed by rating groups & service Identity (uint32, [uint32]) and receives
 * CreditUpdateResponses to update
 * credit
 */
class ChargingCreditPool
    : public CreditPool<CreditKey, CreditUpdateResponse, CreditUsage> {
public:
  static std::unique_ptr<ChargingCreditPool>
  unmarshal(const StoredChargingCreditPool &marshaled);

  StoredChargingCreditPool marshal();

  ChargingCreditPool(const std::string &imsi);

  bool add_used_credit(const CreditKey &key, uint64_t used_tx, uint64_t used_rx,
                       SessionStateUpdateCriteria &update_criteria) override;

  bool
  reset_reporting_credit(const CreditKey &key,
                         SessionStateUpdateCriteria &update_criteria) override;

  void get_updates(std::string imsi, std::string ip_addr,
                   StaticRuleStore &static_rules,
                   DynamicRuleStore *dynamic_rules,
                   std::vector<CreditUsage> *updates_out,
                   std::vector<std::unique_ptr<ServiceAction>> *actions_out,
                   SessionStateUpdateCriteria &update_criteria,
                   const bool force_update = false) override;

  bool
  get_termination_updates(SessionTerminateRequest *termination_out,
                          SessionStateUpdateCriteria &update_criteria) override;

  bool receive_credit(const CreditUpdateResponse &update,
                      SessionStateUpdateCriteria &update_criteria) override;

  uint64_t get_credit(const CreditKey &key, Bucket bucket) const override;

  void add_credit(const CreditKey &key, std::unique_ptr<SessionCredit> credit,
                  SessionStateUpdateCriteria &update_criteria);

  SessionCreditUpdateCriteria *
  get_credit_update(const CreditKey &key,
                    SessionStateUpdateCriteria &update_criteria) override;

  void merge_credit_update(const CreditKey &key,
                           SessionCreditUpdateCriteria &credit_update) override;

  uint32_t get_credit_key_count() const override;

  ChargingReAuthAnswer::Result
  reauth_key(const CreditKey &charging_key,
             SessionStateUpdateCriteria &update_criteria);

  ChargingReAuthAnswer::Result
  reauth_all(SessionStateUpdateCriteria &update_criteria);

private:
  std::unordered_map<CreditKey, std::unique_ptr<SessionCredit>,
                     decltype(&ccHash), decltype(&ccEqual)>
      credit_map_;
  std::string imsi_;

private:
  bool init_new_credit(const CreditUpdateResponse &update,
                       SessionStateUpdateCriteria &update_criteria);

  void populate_output_actions(
      std::string imsi, std::string ip_addr, CreditKey key,
      StaticRuleStore &static_rules, DynamicRuleStore *dynamic_rules,
      std::unique_ptr<ServiceAction> &action,
      std::vector<std::unique_ptr<ServiceAction>> *actions_out) const;
};

/**
 * UsageMonitoringCreditPool manages a pool of credits for PCRF-based usage
 * monitoring. It is keyed by monitoring keys (string) and receives
 * UsageMonitoringUpdateResponse to update credit
 */
class UsageMonitoringCreditPool
    : public CreditPool<std::string, UsageMonitoringUpdateResponse,
                        UsageMonitorUpdate> {
public:
  struct Monitor {
    SessionCredit credit;
    MonitoringLevel level;

    Monitor() : credit(CreditType::MONITORING) {}
  };

  static StoredMonitor
  marshal_monitor(std::unique_ptr<UsageMonitoringCreditPool::Monitor> &monitor);

  static std::unique_ptr<Monitor>
  unmarshal_monitor(const StoredMonitor &marshaled);

  static std::unique_ptr<UsageMonitoringCreditPool>
  unmarshal(const StoredUsageMonitoringCreditPool &marshaled);

  StoredUsageMonitoringCreditPool marshal();

  UsageMonitoringCreditPool(const std::string &imsi);

  bool add_used_credit(const std::string &key, uint64_t used_tx,
                       uint64_t used_rx,
                       SessionStateUpdateCriteria &update_criteria) override;

  bool
  reset_reporting_credit(const std::string &key,
                         SessionStateUpdateCriteria &update_criteria) override;

  void get_updates(std::string imsi, std::string ip_addr,
                   StaticRuleStore &static_rules,
                   DynamicRuleStore *dynamic_rules,
                   std::vector<UsageMonitorUpdate> *updates_out,
                   std::vector<std::unique_ptr<ServiceAction>> *actions_out,
                   SessionStateUpdateCriteria &update_criteria,
                   const bool force_update = false) override;

  bool
  get_termination_updates(SessionTerminateRequest *termination_out,
                          SessionStateUpdateCriteria &update_criteria) override;

  bool receive_credit(const UsageMonitoringUpdateResponse &update,
                      SessionStateUpdateCriteria &update_criteria) override;

  uint64_t get_credit(const std::string &key, Bucket bucket) const override;

  void add_monitor(const std::string &key,
                   std::unique_ptr<UsageMonitoringCreditPool::Monitor> monitor,
                   SessionStateUpdateCriteria &update_criteria);

  SessionCreditUpdateCriteria *
  get_credit_update(const std::string &key,
                    SessionStateUpdateCriteria &update_criteria) override;

  void merge_credit_update(const std::string &key,
                           SessionCreditUpdateCriteria &credit_update) override;

  uint32_t get_credit_key_count() const override;

  std::unique_ptr<std::string> get_session_level_key();

private:
  std::unordered_map<std::string, std::unique_ptr<Monitor>> monitor_map_;
  std::string imsi_;
  std::unique_ptr<std::string> session_level_key_;

private:
  void update_session_level_key(const UsageMonitoringUpdateResponse &update,
                                SessionStateUpdateCriteria &update_criteria);
  bool init_new_credit(const UsageMonitoringUpdateResponse &update,
                       SessionStateUpdateCriteria &update_criteria);
  void populate_output_actions(
      std::string imsi, std::string ip_addr, std::string key,
      StaticRuleStore &static_rules, DynamicRuleStore *dynamic_rules,
      std::unique_ptr<ServiceAction> &action,
      std::vector<std::unique_ptr<ServiceAction>> *actions_out) const;
};

} // namespace magma
