/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package service

import (
	"context"
	"fmt"
	"log"
	"magma/gateway/service_registry"
	"magma/orc8r/lib/go/protos"
	"magma/orc8r/lib/go/registry"

	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type testSyncRpcService struct {
	syncRpcReqCh  chan *protos.SyncRPCRequest
	syncRpcRespCh chan *protos.SyncRPCResponse
}

func (svc *testSyncRpcService) EstablishSyncRPCStream(stream protos.SyncRPCService_EstablishSyncRPCStreamServer) error {
	go func() {
		for {
			resp, _ := stream.Recv()
			svc.syncRpcRespCh <- resp
		}
	}()

	for req := range svc.syncRpcReqCh {
		stream.Send(req)
	}
	return nil
}

func (svc *testSyncRpcService) SyncRPC(stream protos.SyncRPCService_SyncRPCServer) error {
	return nil
}

// run instance of the test grpc service
func runTestSyncRpcService(server *testSyncRpcService, grpcPortCh chan string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":0"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	v := strings.Split(lis.Addr().String(), ":")
	grpcPortCh <- v[len(v)-1]
	grpcServer := grpc.NewServer()
	protos.RegisterSyncRPCServiceServer(grpcServer, server)
	grpcServer.Serve(lis)
}

type testBrokerImpl struct {
	testRespCh chan *protos.SyncRPCResponse
}

func (t *testBrokerImpl) send(_ context.Context, _ string, _ *protos.SyncRPCRequest, respCh chan *protos.SyncRPCResponse) {
	resp := <-t.testRespCh
	respCh <- resp
}

func TestSyncRpcClient(t *testing.T) {
	BrokerRespCh := make(chan *protos.SyncRPCResponse)
	testBrokerImpl := &testBrokerImpl{testRespCh: BrokerRespCh}
	reg := service_registry.Get()
	cfg := Config{SyncRpcHeartbeatInterval: 100 * time.Second}
	client := SyncRpcClient{
		serviceRegistry: reg,
		respCh:          make(chan *protos.SyncRPCResponse),
		terminatedReqs:  make(map[uint32]bool),
		outstandingReqs: make(map[uint32]context.CancelFunc),
		cfg:             cfg,
		broker:          testBrokerImpl,
	}
	ctx := context.Background()

	grpcPortCh := make(chan string)
	svcSyncRpcReqCh := make(chan *protos.SyncRPCRequest)
	svcSyncRpcRespCh := make(chan *protos.SyncRPCResponse)
	svc := &testSyncRpcService{
		syncRpcReqCh:  svcSyncRpcReqCh,
		syncRpcRespCh: svcSyncRpcRespCh,
	}
	go runTestSyncRpcService(svc, grpcPortCh)
	grpcPort := <-grpcPortCh
	go func() {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", grpcPort),
			grpc.WithInsecure())
		if err != nil {
			t.Fatal("Failed creating a test client")
			return
		}
		defer conn.Close()

		grpcClient := protos.NewSyncRPCServiceClient(conn)
		client.runSyncRpcClient(ctx, grpcClient)
	}()
	// send a syncRpcRequest and verify if we receive a proper syncRpcResponse
	reg.AddService(registry.ServiceLocation{
		Name: "testService",
		Host: "localhost",
		Port: 9999,
	})
	svcSyncRpcReqCh <- &protos.SyncRPCRequest{ReqId: 1, ReqBody: &protos.GatewayRequest{Authority: "testService"}}
	BrokerRespCh <- &protos.SyncRPCResponse{ReqId: 1}

	select {
	case resp := <-svcSyncRpcRespCh:
		assert.Equal(t, resp.ReqId, uint32(1))
	case <-time.After(3 * time.Second):
		t.Fatal("Timeout 1")
	}

	// send a SyncRpcRequest terminating a request
	svcSyncRpcReqCh <- &protos.SyncRPCRequest{ReqId: 2, ReqBody: &protos.GatewayRequest{Authority: "testService"}}
	svcSyncRpcReqCh <- &protos.SyncRPCRequest{ReqId: 2, ConnClosed: true}
	BrokerRespCh <- &protos.SyncRPCResponse{ReqId: 2}
	timer := time.NewTimer(time.Second)

	select {
	case resp := <-svcSyncRpcRespCh:
		t.Fatalf("no response was expected. recd %v", resp)
	case <-timer.C:
		break
	}

	// send a syncRpcRequest which is already being handled
	svcSyncRpcReqCh <- &protos.SyncRPCRequest{ReqId: 3, ReqBody: &protos.GatewayRequest{Authority: "testService"}}
	svcSyncRpcReqCh <- &protos.SyncRPCRequest{ReqId: 3, ReqBody: &protos.GatewayRequest{Authority: "testService"}}

	select {
	case resp := <-svcSyncRpcRespCh:
		assert.Contains(t, resp.RespBody.Err, "already being handled")
	case <-time.After(3 * time.Second):
		t.Fatal("Timeout 2")
	}

	// finally check if we receive periodic heartbeats
	// run new client with short heartbeat interval
	cfg.SyncRpcHeartbeatInterval = 1 * time.Second
	client2 := SyncRpcClient{serviceRegistry: reg, cfg: cfg}
	go func() {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", grpcPort),
			grpc.WithInsecure())
		if err != nil {
			t.Fatal("Failed creating a test client")
			return
		}
		defer conn.Close()

		grpcClient := protos.NewSyncRPCServiceClient(conn)
		client2.runSyncRpcClient(ctx, grpcClient)
	}()
	select {
	case resp := <-svcSyncRpcRespCh:
		assert.Equal(t, resp.HeartBeat, true)
	case <-time.After(3 * time.Second):
		t.Fatal("Timeout 3")
	}
}
