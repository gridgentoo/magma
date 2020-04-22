/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/emakeev/snowflake"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"magma/gateway/config"
	bootstrap_client "magma/gateway/services/bootstrapper/service"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/pluginimpl/models"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/services/bootstrapper"
	"magma/orc8r/cloud/go/services/bootstrapper/servicers"
	certifierTestInit "magma/orc8r/cloud/go/services/certifier/test_init"
	"magma/orc8r/cloud/go/services/configurator"
	configuratorTestInit "magma/orc8r/cloud/go/services/configurator/test_init"
	configuratorTestUtils "magma/orc8r/cloud/go/services/configurator/test_utils"
	"magma/orc8r/cloud/go/services/device"
	deviceTestInit "magma/orc8r/cloud/go/services/device/test_init"
	"magma/orc8r/cloud/go/test_utils"
	"magma/orc8r/lib/go/protos"
	"magma/orc8r/lib/go/security/csr"
	"magma/orc8r/lib/go/security/key"
)

const (
	echoType  = "ECHO"
	rsaType   = "SOFTWARE_RSA_SHA256"
	ecdsaType = "SOFTWARE_ECDSA_SHA256"
)

func testWithECHO(
	t *testing.T, networkId string, srv *servicers.BootstrapperServer, ctx context.Context) {

	testAgHwId := "test_ag_echo"

	configuratorTestUtils.RegisterGateway(
		t,
		networkId,
		testAgHwId,
		&models.GatewayDevice{
			HardwareID: testAgHwId,
			Key:        &models.ChallengeKey{KeyType: echoType},
		},
	)

	// check challenge type
	challenge, err := srv.GetChallenge(ctx, &protos.AccessGatewayID{Id: testAgHwId})
	assert.NoError(t, err)
	assert.Equal(t, challenge.KeyType, protos.ChallengeKey_ECHO)

	// create response
	response := &protos.Response_EchoResponse{
		EchoResponse: &protos.Response_Echo{Response: challenge.Challenge},
	}
	csr, err := csr.CreateCSR(time.Duration(time.Hour*24*10), "cn", "cn")
	assert.NoError(t, err)
	resp := protos.Response{
		HwId:      &protos.AccessGatewayID{Id: testAgHwId},
		Challenge: challenge.Challenge,
		Response:  response,
		Csr:       csr,
	}
	cert, err := srv.RequestSign(ctx, &resp)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
}

func testWithRSA(
	t *testing.T, networkId string, srv *servicers.BootstrapperServer, ctx context.Context) {

	testAgHwId := "test_ag_rsa"
	privateKey, err := key.GenerateKey("", 1024)
	assert.NoError(t, err)
	marshaledPubKey, err := x509.MarshalPKIXPublicKey(key.PublicKey(privateKey))
	assert.NoError(t, err)

	pubKey := strfmt.Base64(marshaledPubKey)
	configuratorTestUtils.RegisterGateway(
		t,
		networkId,
		testAgHwId,
		&models.GatewayDevice{
			HardwareID: testAgHwId,
			Key: &models.ChallengeKey{
				KeyType: rsaType,
				Key:     &pubKey,
			},
		})

	challenge, err := srv.GetChallenge(ctx, &protos.AccessGatewayID{Id: testAgHwId})
	assert.NoError(t, err)
	assert.Equal(t, challenge.KeyType, protos.ChallengeKey_SOFTWARE_RSA_SHA256)

	// sign challenge with private key
	hashed := sha256.Sum256(challenge.Challenge)
	signature, err := rsa.SignPKCS1v15(
		rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashed[:])
	assert.NoError(t, err)

	// create response
	response := &protos.Response_RsaResponse{
		RsaResponse: &protos.Response_RSA{Signature: signature},
	}
	csr, err := csr.CreateCSR(time.Duration(time.Hour*24*10), "cn", "cn")
	assert.NoError(t, err)
	resp := protos.Response{
		HwId:      &protos.AccessGatewayID{Id: testAgHwId},
		Challenge: challenge.Challenge,
		Response:  response,
		Csr:       csr,
	}
	cert, err := srv.RequestSign(ctx, &resp)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
}

func testWithECDSA(
	t *testing.T, networkId string, srv *servicers.BootstrapperServer, ctx context.Context) {

	testAgHwId := "test_ag_ecdsa"
	privateKey, err := key.GenerateKey("P256", 0)
	assert.NoError(t, err)
	marshaledPubKey, err := x509.MarshalPKIXPublicKey(key.PublicKey(privateKey))
	assert.NoError(t, err)

	pubKey := strfmt.Base64(marshaledPubKey)
	configuratorTestUtils.RegisterGateway(
		t,
		networkId,
		testAgHwId,
		&models.GatewayDevice{
			HardwareID: testAgHwId,
			Key: &models.ChallengeKey{
				KeyType: ecdsaType,
				Key:     &pubKey,
			},
		},
	)

	challenge, err := srv.GetChallenge(ctx, &protos.AccessGatewayID{Id: testAgHwId})
	assert.NoError(t, err)
	assert.Equal(t, challenge.KeyType, protos.ChallengeKey_SOFTWARE_ECDSA_SHA256)

	hashed := sha256.Sum256(challenge.Challenge)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey.(*ecdsa.PrivateKey), hashed[:])
	assert.NoError(t, err)

	// create response
	response := &protos.Response_EcdsaResponse{
		EcdsaResponse: &protos.Response_ECDSA{R: r.Bytes(), S: s.Bytes()},
	}
	csr, err := csr.CreateCSR(time.Duration(time.Hour*24*10), "cn", "cn")
	assert.NoError(t, err)
	resp := protos.Response{
		HwId:      &protos.AccessGatewayID{Id: testAgHwId},
		Challenge: challenge.Challenge,
		Response:  response,
		Csr:       csr,
	}
	cert, err := srv.RequestSign(ctx, &resp)
	assert.NoError(t, err)
	assert.NotNil(t, cert)
}

// Test with real GW bootstrapper
func testWithGatewayBootstrapper(t *testing.T, networkId string) {
	srv, lis := test_utils.NewTestService(t, orc8r.ModuleName, bootstrapper.ServiceName)
	assert.Equal(t, protos.ServiceInfo_STARTING, srv.State)
	assert.Equal(t, protos.ServiceInfo_APP_UNHEALTHY, srv.Health)

	privateKey, err := key.GenerateKey("", 2048)
	assert.NoError(t, err)

	bootstrServer, err := servicers.NewBootstrapperServer(privateKey.(*rsa.PrivateKey))
	assert.NoError(t, err)

	protos.RegisterBootstrapperServer(srv.GrpcServer, bootstrServer)
	srv.GrpcServer.RegisterService(protos.GetLegacyBootstrapperDesc(), bootstrServer)

	go srv.RunTest(lis)

	srvIp, srvPort, err := net.SplitHostPort(lis.Addr().String())
	assert.NoError(t, err)
	srvPortInt, err := strconv.Atoi(srvPort)
	assert.NoError(t, err)
	dir, err := ioutil.TempDir("", "magma_bst")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	completed := false
	completedPtr := &completed
	completeChan := make(chan interface{})
	go func(t *testing.T) {
		for i := range completeChan {
			*completedPtr = true
			switch u := i.(type) {
			case bootstrap_client.BootstrapCompletion:
				t.Logf("bootstrap comnpleted with result: %v", u.Result)
				assert.NoError(t, err)
			case struct{}:
			default:
				t.Errorf("unknown completion type: %T", u)
			}
		}
	}(t)

	config.OverwriteControlProxyConfigs(&config.ControlProxyCfg{
		RootCaFile:           "",
		GwCertFile:           dir + "/gateway.crt",
		GwCertKeyFile:        dir + "/gateway.key",
		BootstrapAddr:        srvIp,
		BootstrapPort:        srvPortInt,
		ProxyCloudConnection: true,
	})

	mdc := &config.MagmadCfg{}
	mdc.BootstrapConfig.ChallengeKey = dir + "/gw_challenge.key"
	config.OverwriteMagmadConfigs(mdc)

	b := bootstrap_client.NewBootstrapper(completeChan)
	err = b.Initialize()
	assert.NoError(t, err)

	uuid, err := snowflake.Get()
	assert.NoError(t, err)
	gwHwId := uuid.String()

	ck, err := key.ReadKey(mdc.BootstrapConfig.ChallengeKey)
	assert.NoError(t, err)
	pubKey, err := x509.MarshalPKIXPublicKey(key.PublicKey(ck))
	assert.NoError(t, err)
	encodedPubKey := strfmt.Base64(pubKey)

	configuratorTestUtils.RegisterGateway(
		t,
		networkId,
		gwHwId,
		&models.GatewayDevice{
			HardwareID: gwHwId,
			Key: &models.ChallengeKey{
				KeyType: ecdsaType,
				Key:     &encodedPubKey,
			},
		},
	)

	err = b.PeriodicCheck(time.Now())
	assert.NoError(t, err)
	completeChan <- struct{}{} // 'flush' the chan
	assert.True(t, completed)

	// reset configs
	config.OverwriteMagmadConfigs(nil)
	config.OverwriteControlProxyConfigs(nil)
}

func testNegative(
	t *testing.T, networkId string, srv *servicers.BootstrapperServer, ctx context.Context) {

	testAgHwId := "test_ag_negative"
	privateKey, err := key.GenerateKey("P256", 0)
	assert.NoError(t, err)
	marshaledPubKey, err := x509.MarshalPKIXPublicKey(key.PublicKey(privateKey))
	assert.NoError(t, err)

	pubKey := strfmt.Base64(marshaledPubKey)
	configuratorTestUtils.RegisterGateway(
		t,
		networkId,
		testAgHwId,
		&models.GatewayDevice{
			HardwareID: testAgHwId,
			Key: &models.ChallengeKey{
				KeyType: "10",
				Key:     &pubKey,
			},
		},
	)

	// cannot get challenge because of unsupported key type
	_, err = srv.GetChallenge(ctx, &protos.AccessGatewayID{Id: testAgHwId})
	assert.Error(t, err)

	configuratorTestUtils.RemoveGateway(t, networkId, testAgHwId)

	configuratorTestUtils.RegisterGateway(
		t,
		networkId,
		testAgHwId,
		&models.GatewayDevice{
			HardwareID: testAgHwId,
			Key: &models.ChallengeKey{
				KeyType: rsaType,
				Key:     &pubKey,
			},
		},
	)

	challenge, err := srv.GetChallenge(ctx, &protos.AccessGatewayID{Id: testAgHwId})
	assert.NoError(t, err)

	// compute response
	hashed := sha256.Sum256(challenge.Challenge)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey.(*ecdsa.PrivateKey), hashed[:])
	assert.NoError(t, err)

	csr, err := csr.CreateCSR(time.Duration(time.Hour*24*10), "cn", "cn")
	assert.NoError(t, err)

	// create response
	response := &protos.Response_EcdsaResponse{
		EcdsaResponse: &protos.Response_ECDSA{R: r.Bytes(), S: s.Bytes()},
	}

	// mess up challenge
	resp := protos.Response{
		HwId:      &protos.AccessGatewayID{Id: testAgHwId},
		Challenge: []byte("mess up challenge"),
		Response:  response,
		Csr:       csr,
	}
	_, err = srv.RequestSign(ctx, &resp)
	assert.Error(t, err)

	// mess up csr
	resp = protos.Response{
		HwId:      &protos.AccessGatewayID{Id: testAgHwId},
		Challenge: challenge.Challenge,
		Response:  response,
		Csr:       nil,
	}
	_, err = srv.RequestSign(ctx, &resp)
	assert.Error(t, err)

	// mess up response
	response = &protos.Response_EcdsaResponse{
		EcdsaResponse: &protos.Response_ECDSA{R: []byte("12344"), S: s.Bytes()},
	}
	resp = protos.Response{
		HwId:      &protos.AccessGatewayID{Id: testAgHwId},
		Challenge: challenge.Challenge,
		Response:  response,
		Csr:       csr,
	}
	_, err = srv.RequestSign(ctx, &resp)
	assert.Error(t, err)

	// mess up hw_id
	resp = protos.Response{
		HwId:      &protos.AccessGatewayID{Id: "mess up hw_id"},
		Challenge: challenge.Challenge,
		Response:  response,
		Csr:       csr,
	}
	_, err = srv.RequestSign(ctx, &resp)
	assert.Error(t, err)
}

func TestBootstrapperServer(t *testing.T) {
	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	_ = serde.RegisterSerdes(serde.NewBinarySerde(device.SerdeDomain, orc8r.AccessGatewayRecordType, &models.GatewayDevice{}))

	testNetworkID := "bootstrapper_test_network"
	err := configurator.CreateNetwork(configurator.Network{
		ID:   testNetworkID,
		Name: "Test Network Name",
	})
	assert.NoError(t, err)
	exists, err := configurator.DoesNetworkExist(testNetworkID)
	assert.True(t, exists)

	ctx := context.Background()

	// create bootstrapper with short key
	privateKey, err := key.GenerateKey("", 512)
	assert.NoError(t, err)
	_, err = servicers.NewBootstrapperServer(privateKey.(*rsa.PrivateKey))
	assert.Error(t, err)

	// create bootstrapper server
	privateKey, err = key.GenerateKey("", 2048)
	assert.NoError(t, err)
	srv, err := servicers.NewBootstrapperServer(privateKey.(*rsa.PrivateKey))

	// for signing csr
	certifierTestInit.StartTestService(t)

	testWithECHO(t, testNetworkID, srv, ctx)
	ctx = metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-serial", "bla"))
	testWithRSA(t, testNetworkID, srv, ctx)
	ctx = metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-serial", ""))
	testWithECDSA(t, testNetworkID, srv, ctx)
	ctx = metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-cn", "bla"))
	testNegative(t, testNetworkID, srv, ctx)
	testWithGatewayBootstrapper(t, testNetworkID)
}
