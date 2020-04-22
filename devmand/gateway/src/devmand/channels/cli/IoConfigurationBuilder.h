// Copyright (c) 2019-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#pragma once

#include <devmand/cartography/DeviceConfig.h>
#include <devmand/channels/cli/Cli.h>
#include <devmand/channels/cli/CliFlavour.h>
#include <devmand/channels/cli/CliThreadWheelTimekeeper.h>
#include <devmand/channels/cli/PromptAwareCli.h>
#include <devmand/channels/cli/ReadCachingCli.h>
#include <devmand/channels/cli/SshSessionAsync.h>
#include <devmand/channels/cli/TreeCache.h>
#include <devmand/channels/cli/engine/Engine.h>
#include <folly/Executor.h>

namespace devmand::channels::cli {

using devmand::cartography::DeviceConfig;
using devmand::channels::cli::Cli;
using devmand::channels::cli::CliFlavour;
using devmand::channels::cli::sshsession::SshSessionAsync;
using namespace std;

using folly::Executor;
using folly::SemiFuture;
using folly::Timekeeper;

static constexpr auto configKeepAliveIntervalSeconds =
    "keepAliveIntervalSeconds";
static constexpr auto configMaxCommandTimeoutSeconds =
    "maxCommandTimeoutSeconds";
static constexpr auto reconnectingQuietPeriodConfig = "reconnectingQuietPeriod";
static constexpr auto sshConnectionTimeoutConfig = "sshConnectionTimeout";

/*
 * Responsible for creation of cli stack for single ssh connection.
 * CLIs can be separated to two groups:
 * - persistent CLIs, which provide high level functions like TCP connection
 * error detection.
 *   - KeepaliveCli - sends keepalive commands periodically
 *   - ReconnectingCli - reestablishes inner layers upon connection error
 * - inner CLIs. They are destroyed and recreated whenever TCP connection is
 * reestablished.
 *   - QueuedCli - orders commands so that only one command can be executed at a
 * time
 *   - LoggingCli - logs aggregated view of device and caching layers
 *   - TimeoutTrackingCli - errors on exceeded timeout
 *   - TreeCacheCli - maintain high level cache
 *   - ReadCachingCli - maintains low level cache
 *   - LoggingCli - logs communication with device
 *   - PromptAwareCli + SshSessionAsync - lowest level, deals with ssh, prompt
 * resolution
 */
class IoConfigurationBuilder {
 public:
  struct ConnectionParameters {
    string username;
    string password;
    string ip;
    string id;
    int port;
    shared_ptr<CliFlavour> flavour;
    chrono::seconds kaTimeout;
    chrono::seconds cmdTimeout;
    chrono::seconds reconnectingQuietPeriod;
    long sshConnectionTimeout; /* in seconds */
    shared_ptr<CliThreadWheelTimekeeper> timekeeper;
    shared_ptr<Executor> sshExecutor;
    shared_ptr<Executor> paExecutor;
    shared_ptr<Executor> rcExecutor;
    shared_ptr<Executor> tcExecutor;
    shared_ptr<Executor> ttExecutor;
    shared_ptr<Executor> lExecutor;
    shared_ptr<Executor> qExecutor;
    shared_ptr<Executor> rExecutor;
    shared_ptr<Executor> kaExecutor;
  };

  IoConfigurationBuilder(
      const DeviceConfig& deviceConfig,
      channels::cli::Engine& engine,
      shared_ptr<CliFlavour> cliFlavour);
  IoConfigurationBuilder(shared_ptr<ConnectionParameters> _connectionParams);

  ~IoConfigurationBuilder();

  shared_ptr<Cli> createAll(
      shared_ptr<CliCache> commandCache,
      shared_ptr<TreeCache> treeCache);

  static Future<shared_ptr<Cli>> createPromptAwareCli(
      shared_ptr<ConnectionParameters> params);

  static shared_ptr<ConnectionParameters> makeConnectionParameters(
      string id,
      string hostname,
      string username,
      string password,
      shared_ptr<CliFlavour> cliFlavour,
      int port,
      chrono::seconds kaTimeout,
      chrono::seconds cmdTimeout,
      chrono::seconds reconnectingQuietPeriod,
      long sshConnectionTimeout,
      channels::cli::Engine& engine);

  shared_ptr<ConnectionParameters> getConnectionParameters();

 private:
  shared_ptr<ConnectionParameters> connectionParameters;

  shared_ptr<Cli> createAllUsingFactory(
      shared_ptr<CliCache> commandCache,
      shared_ptr<TreeCache> treeCache);

  static chrono::seconds toSeconds(const string& value);

  static string loadConfigValue(
      const std::map<std::string, std::string>& plaintextCliKv,
      const string& key,
      const string& defaultValue);

  static Future<shared_ptr<Cli>> configurePromptAwareCli(
      shared_ptr<PromptAwareCli> cli,
      shared_ptr<SshSessionAsync> session,
      shared_ptr<ConnectionParameters> params,
      shared_ptr<Executor> executor);
};
} // namespace devmand::channels::cli
