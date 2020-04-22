/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Package diameter_test tests diameter calls within the magma setting
package gy_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	fegprotos "magma/feg/cloud/go/protos"
	"magma/feg/gateway/diameter"
	"magma/feg/gateway/services/session_proxy/credit_control"
	"magma/feg/gateway/services/session_proxy/credit_control/gy"
	"magma/feg/gateway/services/testcore/ocs/mock_ocs"
	"magma/lte/cloud/go/protos"
)

const (
	testIMSI1      = "000000000000001"
	testIMSI2      = "4321"
	returnedOctets = 1024
	validityTime   = 3600
)

var ocs *mock_ocs.OCSDiamServer

// TestGyClient tests CCR init, update, and terminate messages using a fake
// server
func TestGyClient(t *testing.T) {
	serverConfig := &diameter.DiameterServerConfig{DiameterServerConnConfig: diameter.DiameterServerConnConfig{
		Addr:     "127.0.0.1:0",
		Protocol: "tcp"},
	}
	clientConfig := getClientConfig()
	serverConfig, _ = startServer(clientConfig, serverConfig, gy.PerSessionInit)
	gyGlobalConfig := getGyGlobalConfig("")
	gyClient := gy.NewGyClient(
		clientConfig,
		serverConfig,
		getReAuthHandler(), nil, gyGlobalConfig,
	)

	// send init
	ccrInit := &gy.CreditControlRequest{
		SessionID:     "1",
		Type:          credit_control.CRTInit,
		IMSI:          testIMSI1,
		RequestNumber: 0,
		Credits:       nil,
		UeIPV4:        "192.168.1.1",
		SpgwIPV4:      "10.10.10.10",
	}
	done := make(chan interface{}, 1000)

	log.Printf("Sending CCR-Init")
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))
	answer := gy.GetAnswer(done)
	log.Printf("Received CCA-Init")
	assert.Equal(t, ccrInit.SessionID, answer.SessionID)
	assert.Equal(t, ccrInit.RequestNumber, answer.RequestNumber)
	assert.Equal(t, len(answer.Credits), 0)
	calledStationID, err := mock_ocs.GetAVP(ocs.LastMessageReceived, "Called-Station-Id")
	assert.NoError(t, err)
	assert.Equal(t, "", calledStationID)

	// send multiple updates
	ccrUpdates := []*gy.CreditControlRequest{
		{
			SessionID:     "1",
			Type:          credit_control.CRTUpdate,
			IMSI:          testIMSI1,
			RequestNumber: 1,
			Credits: []*gy.UsedCredits{{
				RatingGroup:  1,
				InputOctets:  1024,
				OutputOctets: 2048,
				TotalOctets:  3072,
			},
			}},
		{
			SessionID:     "2",
			Type:          credit_control.CRTUpdate,
			IMSI:          testIMSI2,
			RequestNumber: 1,
			Credits: []*gy.UsedCredits{{
				RatingGroup:  1,
				InputOctets:  1024,
				OutputOctets: 2048,
				TotalOctets:  3072,
			},
			}},
	}

	for _, update := range ccrUpdates {
		assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, update))
	}

	for i := 0; i < 2; i++ {
		update := gy.GetAnswer(done)
		assert.Equal(t, uint64(returnedOctets), *update.Credits[0].GrantedUnits.TotalOctets)
		assert.Equal(t, uint32(validityTime), update.Credits[0].ValidityTime)
		assert.Equal(t, ccrUpdates[i].SessionID, update.SessionID)
		assert.Equal(t, ccrUpdates[i].RequestNumber, update.RequestNumber)
	}

	// send terminates
	ccrTerminate := &gy.CreditControlRequest{
		SessionID:     "1",
		Type:          credit_control.CRTTerminate,
		IMSI:          testIMSI1,
		RequestNumber: 2,
		Credits: []*gy.UsedCredits{{
			RatingGroup:  1,
			InputOctets:  1024,
			OutputOctets: 2048,
			TotalOctets:  3072,
		}},
	}
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrTerminate))
	terminate := gy.GetAnswer(done)
	assert.Equal(t, len(terminate.Credits), 0)
	assert.Equal(t, ccrTerminate.SessionID, terminate.SessionID)
	assert.Equal(t, ccrTerminate.RequestNumber, terminate.RequestNumber)

	// Connection disabling should cause CCR to fail
	gyClient.DisableConnections(10 * time.Second)
	assert.Error(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))

	// CCR Success after enabling connections
	gyClient.EnableConnections()
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))
}

// TestGyClient test different options on global configuration
func TestGyClientWithGyGlobalConf(t *testing.T) {
	serverConfig := &diameter.DiameterServerConfig{DiameterServerConnConfig: diameter.DiameterServerConnConfig{
		Addr:     "127.0.0.1:0",
		Protocol: "tcp"},
	}

	clientConfig := getClientConfig()
	serverConfig, _ = startServer(clientConfig, serverConfig, gy.PerSessionInit)
	overWriteApn := "gy.Apn.magma.com"
	gyGlobalConfig := getGyGlobalConfig(overWriteApn)
	gyClient := gy.NewGyClient(
		clientConfig,
		serverConfig,
		getReAuthHandler(), nil, gyGlobalConfig,
	)

	// send init
	ccrInit := &gy.CreditControlRequest{
		SessionID:     "1",
		Type:          credit_control.CRTInit,
		IMSI:          testIMSI1,
		RequestNumber: 0,
		Credits:       nil,
		UeIPV4:        "192.168.1.1",
		SpgwIPV4:      "10.10.10.10",
	}
	done := make(chan interface{}, 1000)

	log.Printf("Sending CCR-Init with custom global parameters")
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))
	answer := gy.GetAnswer(done)
	log.Printf("Received CCA-Init")
	calledStationID, err := mock_ocs.GetAVP(ocs.LastMessageReceived, "Called-Station-Id")
	assert.NoError(t, err)
	assert.Equal(t, overWriteApn, calledStationID)
	assert.Equal(t, ccrInit.RequestNumber, answer.RequestNumber)
	assert.Equal(t, len(answer.Credits), 0)
}

func TestGyClientOutOfCredit(t *testing.T) {
	serverConfig := &diameter.DiameterServerConfig{DiameterServerConnConfig: diameter.DiameterServerConnConfig{
		Addr:     "127.0.0.1:0",
		Protocol: "tcp"},
	}
	clientConfig := getClientConfig()
	serverConfig, _ = startServer(clientConfig, serverConfig, gy.PerSessionInit)
	gyGlobalConfig := getGyGlobalConfig("")
	gyClient := gy.NewGyClient(
		clientConfig,
		serverConfig,
		getReAuthHandler(), nil, gyGlobalConfig,
	)

	// send init
	ccrInit := &gy.CreditControlRequest{
		SessionID:     "1",
		Type:          credit_control.CRTInit,
		IMSI:          testIMSI1,
		RequestNumber: 0,
		Credits:       nil,
		UeIPV4:        "192.168.1.1",
		SpgwIPV4:      "10.10.10.10",
	}
	done := make(chan interface{}, 1000)
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))
	gy.GetAnswer(done)

	// send request with (total credits - used credits) < max usage (final units)
	ccrUpdate := &gy.CreditControlRequest{
		SessionID:     "1",
		Type:          credit_control.CRTUpdate,
		IMSI:          testIMSI1,
		RequestNumber: 1,
		Credits: []*gy.UsedCredits{{
			RatingGroup:  1,
			InputOctets:  999990,
			OutputOctets: 0,
			TotalOctets:  999990,
		}},
	}

	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrUpdate))
	update := gy.GetAnswer(done)
	assert.Equal(t, uint64(10), *update.Credits[0].GrantedUnits.TotalOctets)
	assert.True(t, update.Credits[0].IsFinal)
	assert.Equal(t, gy.Terminate, update.Credits[0].FinalAction)
}

func TestGyClientPerKeyInit(t *testing.T) {
	serverConfig := &diameter.DiameterServerConfig{DiameterServerConnConfig: diameter.DiameterServerConnConfig{
		Addr:     "127.0.0.1:0",
		Protocol: "tcp"},
	}
	clientConfig := getClientConfig()
	serverConfig, _ = startServer(clientConfig, serverConfig, gy.PerKeyInit)
	gyGlobalConfig := getGyGlobalConfig("")
	gyClient := gy.NewGyClient(
		clientConfig,
		serverConfig,
		getReAuthHandler(), nil, gyGlobalConfig,
	)

	// send inits
	ccrInits := []*gy.CreditControlRequest{
		{
			SessionID:     "1",
			Type:          credit_control.CRTInit,
			IMSI:          testIMSI1,
			RequestNumber: 1,
			UeIPV4:        "192.168.1.1",
			SpgwIPV4:      "10.10.10.10",
			Credits: []*gy.UsedCredits{{
				RatingGroup: 1,
			},
			}},
		{
			SessionID:     "1",
			Type:          credit_control.CRTInit,
			IMSI:          testIMSI1,
			RequestNumber: 2,
			UeIPV4:        "192.168.1.1",
			SpgwIPV4:      "10.10.10.10",
			Credits: []*gy.UsedCredits{{
				RatingGroup: 2,
			}},
		},
	}
	done := make(chan interface{}, 1000)

	log.Printf("Sending CCR-Updates")
	for _, init := range ccrInits {
		assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, init))
	}

	for i := 0; i < 2; i++ {
		update := gy.GetAnswer(done)
		assert.Equal(t, uint64(returnedOctets), *update.Credits[0].GrantedUnits.TotalOctets)
		assert.Equal(t, uint32(validityTime), update.Credits[0].ValidityTime)
		assert.Equal(t, ccrInits[i].SessionID, update.SessionID)
		assert.Equal(t, ccrInits[i].RequestNumber, update.RequestNumber)
	}
}

func TestGyClientMultipleCredits(t *testing.T) {
	serverConfig := &diameter.DiameterServerConfig{DiameterServerConnConfig: diameter.DiameterServerConnConfig{
		Addr:     "127.0.0.1:0",
		Protocol: "tcp"},
	}
	clientConfig := getClientConfig()
	serverConfig, _ = startServer(clientConfig, serverConfig, gy.PerKeyInit)
	gyGlobalConfig := getGyGlobalConfig("")
	gyClient := gy.NewGyClient(
		clientConfig,
		serverConfig,
		getReAuthHandler(), nil, gyGlobalConfig,
	)

	// send inits
	ccrInit := &gy.CreditControlRequest{
		SessionID:     "1",
		Type:          credit_control.CRTInit,
		IMSI:          testIMSI1,
		RequestNumber: 1,
		UeIPV4:        "192.168.1.1",
		SpgwIPV4:      "10.10.10.10",
		Credits: []*gy.UsedCredits{
			{
				RatingGroup: 1,
			},
			{
				RatingGroup: 2,
			},
			{
				RatingGroup: 3,
			},
		},
	}
	done := make(chan interface{}, 1000)

	log.Printf("Sending CCR-Init")
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))

	ans := gy.GetAnswer(done)
	assert.Equal(t, ans.SessionID, ccrInit.SessionID)
	assert.Equal(t, ans.RequestNumber, ccrInit.RequestNumber)
	assert.Equal(t, len(ans.Credits), 3)
	for _, credit := range ans.Credits {
		assert.Contains(t, []uint32{1, 2, 3}, credit.RatingGroup)
		assert.Equal(t, uint64(returnedOctets), *credit.GrantedUnits.TotalOctets)
		assert.Equal(t, uint32(validityTime), credit.ValidityTime)
	}
}

func TestGyReAuth(t *testing.T) {
	serverConfig := &diameter.DiameterServerConfig{DiameterServerConnConfig: diameter.DiameterServerConnConfig{
		Addr:     "127.0.0.1:3874",
		Protocol: "tcp"},
	}
	clientConfig := getClientConfig()
	serverConfig, ocs := startServer(clientConfig, serverConfig, gy.PerKeyInit)
	gyGlobalConfig := getGyGlobalConfig("")
	gyClient := gy.NewGyClient(
		clientConfig,
		serverConfig,
		getReAuthHandler(), nil, gyGlobalConfig,
	)

	// send one init to set user context in OCS
	sessionID := fmt.Sprintf("IMSI%s-%d", testIMSI1, 1234)
	ccrInit := &gy.CreditControlRequest{
		SessionID:     sessionID,
		Type:          credit_control.CRTInit,
		IMSI:          testIMSI1,
		RequestNumber: 1,
		UeIPV4:        "192.168.1.1",
		SpgwIPV4:      "10.10.10.10",
		Credits: []*gy.UsedCredits{
			{
				RatingGroup: 1,
			},
		},
	}
	done := make(chan interface{}, 1000)

	log.Printf("Sending CCR-Init")
	assert.NoError(t, gyClient.SendCreditControlRequest(serverConfig, done, ccrInit))
	gy.GetAnswer(done)

	// success reauth
	var rg uint32 = 1
	raa, err := ocs.ReAuth(
		context.Background(),
		&fegprotos.ChargingReAuthTarget{Imsi: testIMSI1, RatingGroup: rg},
	)
	assert.NoError(t, err)
	assert.Equal(t, sessionID, raa.SessionId)
	assert.Equal(t, uint32(diam.Success), raa.ResultCode)
}

func getClientConfig() *diameter.DiameterClientConfig {
	return &diameter.DiameterClientConfig{
		Host:        "test.test.com",
		Realm:       "test.com",
		ProductName: "gy_test",
		AppID:       diam.CHARGING_CONTROL_APP_ID,
	}
}

func getGyGlobalConfig(ocsOverwriteApn string) *gy.GyGlobalConfig {
	return &gy.GyGlobalConfig{
		OCSOverwriteApn: ocsOverwriteApn,
	}
}

func startServer(
	client *diameter.DiameterClientConfig,
	server *diameter.DiameterServerConfig,
	initMethod gy.InitMethod,
) (*diameter.DiameterServerConfig, *mock_ocs.OCSDiamServer) {
	serverStarted := make(chan struct{})
	go func() {
		log.Printf("Starting server")
		ocs = mock_ocs.NewOCSDiamServer(
			client,
			&mock_ocs.OCSConfig{
				MaxUsageOctets: &fegprotos.Octets{TotalOctets: returnedOctets},
				MaxUsageTime:   1000,
				ValidityTime:   validityTime,
				ServerConfig:   server,
				GyInitMethod:   initMethod,
			},
		)
		ctx := context.Background()
		ocs.CreateAccount(ctx, &protos.SubscriberID{Id: testIMSI1})
		ocs.CreateAccount(ctx, &protos.SubscriberID{Id: testIMSI2})
		ocs.SetCredit(
			ctx,
			&fegprotos.CreditInfo{
				Imsi:        testIMSI1,
				ChargingKey: 1,
				Volume:      &fegprotos.Octets{TotalOctets: 1000000},
				UnitType:    fegprotos.CreditInfo_Bytes,
			},
		)
		ocs.SetCredit(
			ctx,
			&fegprotos.CreditInfo{
				Imsi:        testIMSI1,
				ChargingKey: 2,
				Volume:      &fegprotos.Octets{TotalOctets: 1000000},
				UnitType:    fegprotos.CreditInfo_Bytes,
			},
		)
		ocs.SetCredit(
			ctx,
			&fegprotos.CreditInfo{
				Imsi:        testIMSI1,
				ChargingKey: 3,
				Volume:      &fegprotos.Octets{TotalOctets: 1000000},
				UnitType:    fegprotos.CreditInfo_Bytes,
			},
		)
		ocs.SetCredit(
			ctx,
			&fegprotos.CreditInfo{
				Imsi:        testIMSI2,
				ChargingKey: 1,
				Volume:      &fegprotos.Octets{TotalOctets: 1000000},
				UnitType:    fegprotos.CreditInfo_Bytes,
			},
		)
		lis, err := ocs.StartListener()
		if err != nil {
			log.Fatalf("Could not start listener, %s", err.Error())
			return
		}
		server.Addr = lis.Addr().String()
		serverStarted <- struct{}{}
		err = ocs.Start(lis)
		if err != nil {
			log.Fatalf("Could not start server, %s", err.Error())
			return
		}
	}()
	<-serverStarted
	time.Sleep(time.Millisecond)
	return server, ocs
}

func getReAuthHandler() gy.ChargingReAuthHandler {
	return func(request *gy.ChargingReAuthRequest) *gy.ChargingReAuthAnswer {
		return &gy.ChargingReAuthAnswer{
			SessionID:  request.SessionID,
			ResultCode: diam.Success,
		}
	}
}
