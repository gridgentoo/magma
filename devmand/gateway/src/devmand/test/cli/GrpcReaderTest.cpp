// Copyright (c) 2020-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#define LOG_WITH_GLOG

#include <magma_logging.h>

#include <devmand/channels/cli/plugin/protocpp/ReaderPlugin.grpc.pb.h>
#include <devmand/devices/cli/translation/GrpcReader.h>
#include <devmand/test/cli/utils/Log.h>
#include <devmand/test/cli/utils/MockCli.h>
#include <folly/executors/CPUThreadPoolExecutor.h>
#include <folly/json.h>
#include <grpc++/grpc++.h>
#include <gtest/gtest.h>
#include <devmand/devices/cli/UbntInterfacePlugin.cpp>
#include <thread>

namespace devmand {
namespace test {
namespace cli {

using namespace std;
using namespace folly;
using namespace grpc;
using namespace devmand::channels::cli;
using namespace devmand::devices::cli;
using namespace devmand::channels::cli::plugin;
using namespace devmand::test::utils::cli;

static void sendCommandRequest(
    string cmd,
    ServerReaderWriter<ReadResponse, ReadRequest>* stream) {
  MLOG(MDEBUG) << "Sending cli request :" << cmd;
  ReadResponse readResponse;
  CliRequest* cliRequest = new CliRequest();
  cliRequest->set_cmd(cmd);
  readResponse.set_allocated_clirequest(cliRequest);
  stream->Write(readResponse);
}

static void sendActualResponse(
    string json,
    ServerReaderWriter<ReadResponse, ReadRequest>* stream) {
  ReadResponse readResponse;
  ActualReadResponse* actualReadResponse = new ActualReadResponse();
  actualReadResponse->set_json(json);
  readResponse.set_allocated_actualreadresponse(actualReadResponse);
  stream->Write(readResponse);
}

static std::unique_ptr<Server> startServer(
    const string& address,
    ReaderPlugin::Service& service) {
  std::string server_address(address);
  ServerBuilder builder;
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<Server> server(builder.BuildAndStart());
  return server;
}

// request 3 commands, then send final response
class DummyReader : public ReaderPlugin::Service {
  Status Read(
      ServerContext* context,
      ServerReaderWriter<ReadResponse, ReadRequest>* stream) {
    (void)context;
    MLOG(MDEBUG) << "Got Read";
    ReadRequest readRequest;
    int remainingCommands = 3;
    bool gotActualReadRequest = false;
    string path;
    while (stream->Read(&readRequest)) {
      if (not gotActualReadRequest) {
        if (readRequest.has_actualreadrequest()) {
          MLOG(MDEBUG) << "Got actualreadrequest";
          gotActualReadRequest = true;
          path = readRequest.actualreadrequest().path();
          // handle request
          sendCommandRequest(path + to_string(remainingCommands), stream);
        } else {
          MLOG(MWARNING) << "First request must be ActualReadRequest";
          return Status(
              StatusCode::INVALID_ARGUMENT,
              "First request must be ActualReadRequest");
        }
      } else {
        // handle command output
        // output should be same output as input
        string output = readRequest.cliresponse().output();
        MLOG(MDEBUG) << "Got cliresponse " << output;
        if (path + to_string(remainingCommands) != output) {
          throw runtime_error("Equality fail");
        }
        remainingCommands -= 1;
        if (remainingCommands > 0) {
          sendCommandRequest(path + to_string(remainingCommands), stream);
        } else {
          // send final response
          MLOG(MDEBUG) << "Sending actualReadResponse:" << path;
          sendActualResponse("{\"path\":\"" + path + "\"}", stream);
          break;
        }
      }
    }
    return Status::OK;
  }
};

class GrpcReaderTest : public ::testing::Test {
 protected:
  shared_ptr<CPUThreadPoolExecutor> testExec;
  shared_ptr<MockedCli> mockedCli;
  Path somePath = "/somepath";
  unique_ptr<Server> server;
  unsigned int port = 50052;
  string address = "127.0.0.1:" + to_string(port);
  DummyReader service;

  void SetUp() override {
    devmand::test::utils::log::initLog();
    testExec = make_shared<CPUThreadPoolExecutor>(5);
    map<string, string> mockedResponses;
    for (int i = 3; i > 0; i--) {
      const string& s = somePath.str() + to_string(i);
      mockedResponses[s] = s;
    }
    mockedResponses.insert(make_pair(
        "show running-config interface 4/1",
        "foo\nmtu 99\ndescription descr\nshutdown\ninterface '4/1'\n"));
    mockedCli = make_shared<MockedCli>(mockedResponses);

    MLOG(MDEBUG) << "Starting server at " << address;

    server = startServer(address, service);
  }

  void TearDown() override {
    server->Shutdown();
    server->Wait();
    testExec->join();
  }
};

TEST_F(GrpcReaderTest, testDummyReader) {
  auto grpcClientChannel =
      grpc::CreateChannel(address, grpc::InsecureChannelCredentials());
  GrpcReader tested(grpcClientChannel, "tested", testExec);
  DeviceAccess deviceAccess = DeviceAccess(mockedCli, "test", testExec);
  dynamic result = tested.read(somePath, deviceAccess).get();
  EXPECT_EQ(result["path"], somePath.str());
}

} // namespace cli
} // namespace test
} // namespace devmand
