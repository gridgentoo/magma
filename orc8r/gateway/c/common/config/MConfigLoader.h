/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */
#pragma once

namespace magma {

/**
 * MConfigLoader is used to load mconfig files for service configurations
 */
class MConfigLoader {
public:
  /**
   * load_service_mconfig loads an mconfig from the statically defined files.
   * @param service_name - name of service to load
   * @param message - pointer to protobuf message to load file to. Note that
   *                  this should match the message defined in mconfigs.proto.
   * @returns true if message was parsed successfully, false if the file reading
   *          failed, the service configuration is missing, or the message
   *          passed is not the right type
   */
  bool load_service_mconfig(
    const std::string& service_name,
    google::protobuf::Message* message);

private:
  static constexpr const char* DYNAMIC_MCONFIG_PATH
    = "/var/opt/magma/configs/gateway.mconfig";
  static constexpr const char* CONFIG_DIR = "/etc/magma";
  static constexpr const char* MCONFIG_FILE_NAME = "gateway.mconfig";

private:
  void get_mconfig_file(std::ifstream* file);
};

}
