/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package integration

import (
	"context"
	"fmt"

	"magma/cwf/gateway/registry"
	fegprotos "magma/feg/cloud/go/protos"
	"magma/feg/gateway/object_store"
	"magma/feg/gateway/policydb"
	"magma/feg/gateway/services/testcore/hss"
	lteprotos "magma/lte/cloud/go/protos"
	registryTestUtils "magma/orc8r/cloud/go/test_utils"
	"magma/orc8r/lib/go/protos"

	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Wrapper for GRPC Client functionality
type hssClient struct {
	fegprotos.HSSConfiguratorClient
	cc *grpc.ClientConn
}

// Wrapper for GRPC Client functionality
type pcrfClient struct {
	fegprotos.MockPCRFClient
	cc *grpc.ClientConn
}

// Wrapper for GRPC Client functionality
type ocsClient struct {
	fegprotos.MockOCSClient
	cc *grpc.ClientConn
}

// Wrapper for GRPC Client functionality
type pipelinedClient struct {
	lteprotos.PipelinedClient
	cc *grpc.ClientConn
}

// Wrapper for PolicyDB objects
type policyDBWrapper struct {
	redisClient      object_store.RedisClient
	policyMap        object_store.ObjectMap
	baseNameMap      object_store.ObjectMap
	omniPresentRules object_store.ObjectMap
}

/**  ========== HSS Helpers ========== **/
// getHSSClient is a utility function to getHSSClient a RPC connection to a
// remote HSS service.
func getHSSClient() (*hssClient, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = registry.GetConnection(MockHSSRemote)
	if err != nil {
		errMsg := fmt.Sprintf("HSS client initialization error: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return &hssClient{
		fegprotos.NewHSSConfiguratorClient(conn),
		conn,
	}, err
}

// addSubscriber tries to add this subscriber to the HSS server.
// This function returns an AlreadyExists error if the subscriber has already
// been added.
// Input: The subscriber data which will be added.
func addSubscriberToHSS(sub *lteprotos.SubscriberData) error {
	err := hss.VerifySubscriberData(sub)
	if err != nil {
		errMsg := fmt.Errorf("Invalid AddSubscriberRequest provided: %s", err)
		return errors.New(errMsg.Error())
	}
	cli, err := getHSSClient()
	if err != nil {
		return err
	}
	_, err = cli.AddSubscriber(context.Background(), sub)
	return err
}

func deleteSubscribersFromHSS(subscriberID string) error {
	cli, err := getHSSClient()
	if err != nil {
		return err
	}
	_, err = cli.DeleteSubscriber(context.Background(), &lteprotos.SubscriberID{Id: subscriberID})
	return err
}

/**  ========== PCRF Helpers ========== **/
// getPCRFClient is a utility function to get an RPC connection to a
// remote PCRF service.
func getPCRFClient() (*pcrfClient, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = registry.GetConnection(MockPCRFRemote)
	if err != nil {
		errMsg := fmt.Sprintf("PCRF client initialization error: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return &pcrfClient{
		MockPCRFClient: fegprotos.NewMockPCRFClient(conn),
		cc:             conn,
	}, err
}

func sendPolicyReAuthRequest(target *fegprotos.PolicyReAuthTarget) (*fegprotos.PolicyReAuthAnswer, error) {
	cli, err := getPCRFClient()
	if err != nil {
		return nil, err
	}
	raa, err := cli.ReAuth(context.Background(), target)
	return raa, err
}

func sendPolicyAbortSession(target *fegprotos.PolicyAbortSessionRequest) (*fegprotos.PolicyAbortSessionResponse, error) {
	cli, err := getPCRFClient()
	if err != nil {
		return nil, err
	}
	raa, err := cli.AbortSession(context.Background(), target)
	return raa, err
}

// addSubscriber tries to add this subscriber to the PCRF server.
// Input: The subscriber data which will be added.
func addSubscriberToPCRF(sub *lteprotos.SubscriberID) error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	_, err = cli.CreateAccount(context.Background(), sub)
	return err
}

func clearSubscribersFromPCRF() error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	_, err = cli.ClearSubscribers(context.Background(), &protos.Void{})
	return err
}

func addPCRFRules(rules *fegprotos.AccountRules) error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	_, err = cli.SetRules(context.Background(), rules)
	return err
}

func addPCRFUsageMonitors(monitorInfo *fegprotos.UsageMonitorConfiguration) error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	_, err = cli.SetUsageMonitors(context.Background(), monitorInfo)
	return err
}

func usePCRFMockDriver() error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	_, err = cli.SetPCRFConfigs(context.Background(), &fegprotos.PCRFConfigs{UseMockDriver: true})
	return err
}

func clearPCRFMockDriver() error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	_, err = cli.SetPCRFConfigs(context.Background(), &fegprotos.PCRFConfigs{UseMockDriver: false})
	return err
}

func setPCRFExpectations(expectations []*fegprotos.GxCreditControlExpectation, defaultAnswer *fegprotos.GxCreditControlAnswer) error {
	cli, err := getPCRFClient()
	if err != nil {
		return err
	}
	request := &fegprotos.GxCreditControlExpectations{
		Expectations: expectations,
		GxDefaultCca: defaultAnswer,
	}
	if defaultAnswer != nil {
		request.UnexpectedRequestBehavior = fegprotos.UnexpectedRequestBehavior_CONTINUE_WITH_DEFAULT_ANSWER
	}
	_, err = cli.SetExpectations(context.Background(), request)
	return err
}

func getAssertExpectationsResult() ([]*fegprotos.ExpectationResult, []*fegprotos.ErrorByIndex, error) {
	cli, err := getPCRFClient()
	if err != nil {
		return nil, nil, nil
	}
	res, err := cli.AssertExpectations(context.Background(), &protos.Void{})
	if err != nil {
		return nil, nil, err
	}
	return res.Results, res.Errors, nil
}

/**  ========== OCS Helpers ========== **/
// getOCSClient is a utility function to an RPC connection to a
// remote OCS service.
func getOCSClient() (*ocsClient, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = registry.GetConnection(MockOCSRemote)
	if err != nil {
		errMsg := fmt.Sprintf("PCRF client initialization error: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return &ocsClient{
		fegprotos.NewMockOCSClient(conn),
		conn,
	}, err
}

// setNewOCSConfig tries to override the default ocs settings
// Input: ocsConfig data
func setNewOCSConfig(ocsConfig *fegprotos.OCSConfig) error {
	cli, err := getOCSClient()
	if err != nil {
		return err
	}
	_, err = cli.SetOCSSettings(context.Background(), ocsConfig)
	return err
}

// addSubscriber tries to add this subscriber to the OCS server.
// Input: The subscriber data which will be added.
func addSubscriberToOCS(sub *lteprotos.SubscriberID) error {
	cli, err := getOCSClient()
	if err != nil {
		return err
	}
	_, err = cli.CreateAccount(context.Background(), sub)
	return err
}

func clearSubscribersFromOCS() error {
	cli, err := getOCSClient()
	if err != nil {
		return err
	}
	_, err = cli.ClearSubscribers(context.Background(), &protos.Void{})
	return err
}

// setCreditOCS tries to set a credit for this subscriber to the OCS server
// Input: The credit info data which will be set
func setCreditOnOCS(creditInfo *fegprotos.CreditInfo) error {
	cli, err := getOCSClient()
	if err != nil {
		return err
	}
	_, err = cli.SetCredit(context.Background(), creditInfo)
	return err
}

// sendChargingReAuthRequest triggers a RAR from OCS to Sessiond
// Input: ChargingReAuthTarget
func sendChargingReAuthRequest(imsi string, ratingGroup uint32) (*fegprotos.ChargingReAuthAnswer, error) {
	cli, err := getOCSClient()
	if err != nil {
		return nil, err
	}
	raa, err := cli.ReAuth(context.Background(),
		&fegprotos.ChargingReAuthTarget{Imsi: imsi, RatingGroup: ratingGroup})
	return raa, err
}

/**  ========== Pipelined Helpers ========== **/
// getPipelinedClient is a utility function to an RPC connection to a
// remote OCS service.
func getPipelinedClient() (*pipelinedClient, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = registryTestUtils.GetConnectionWithAuthority(PipelinedRemote)
	if err != nil {
		errMsg := fmt.Sprintf("Pipelined client initialization error: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return &pipelinedClient{
		lteprotos.NewPipelinedClient(conn),
		conn,
	}, err
}

func deactivateSubscriberFlows(imsi string, ruleIDs []string) error {
	cli, err := getPipelinedClient()
	if err == nil && cli != nil {
		_, err = cli.DeactivateFlows(context.Background(), &lteprotos.DeactivateFlowsRequest{
			Sid:     &lteprotos.SubscriberID{Id: imsi},
			RuleIds: ruleIDs,
		})
	}
	return err
}

func getPolicyUsage() (*lteprotos.RuleRecordTable, error) {
	client, _ := getPipelinedClient()
	return client.GetPolicyUsage(context.Background(), &protos.Void{})
}

/**  ========== PolicyDB related Helpers ========== **/
// In the actual CWAG setup, PolicyRules and BaseNames managed by PolicyDB are
// streamed down from the cloud. Since this integration test setup does not
// include the cloud, we will get around this by directly modifying the redis
// DB.
func initializePolicyDBWrapper() (*policyDBWrapper, error) {
	address, err := registry.GetServiceAddress(RedisRemote)
	if err != nil {
		return nil, err
	}
	redisClientImpl := &object_store.RedisClientImpl{
		RawClient: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf(address),
			Password: "",
			DB:       0,
		}),
	}
	policyMap := object_store.NewRedisMap(
		redisClientImpl,
		"policydb:rules",
		policydb.GetPolicySerializer(),
		policydb.GetPolicyDeserializer(),
	)
	baseNameMap := object_store.NewRedisMap(
		redisClientImpl,
		"policydb:base_names",
		policydb.GetBaseNameSerializer(),
		policydb.GetBaseNameDeserializer(),
	)
	omniPresentRules := object_store.NewRedisMap(
		redisClientImpl,
		"policydb:omnipresent_rules",
		policydb.GetRuleMappingSerializer(),
		policydb.GetRuleMappingDeserializer(),
	)
	return &policyDBWrapper{
		redisClient:      redisClientImpl,
		policyMap:        policyMap,
		baseNameMap:      baseNameMap,
		omniPresentRules: omniPresentRules,
	}, nil
}
