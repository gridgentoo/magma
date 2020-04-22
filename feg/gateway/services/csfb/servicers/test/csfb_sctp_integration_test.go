/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package test

import (
	"testing"
	"unsafe"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/services/csfb/servicers"
	"magma/feg/gateway/services/csfb/servicers/encode/message"
	"magma/feg/gateway/services/csfb/test_init"
	orcprotos "magma/orc8r/lib/go/protos"

	"github.com/ishidawataru/sctp"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestCsfbServer_EPSDetach_Integration(t *testing.T) {
	req := &protos.EPSDetachIndication{
		Imsi:                         "111111",
		MmeName:                      "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcde",
		ImsiDetachFromEpsServiceType: []byte{byte(0x11)},
	}

	// serverUpFlag := make(chan bool)
	msgSentFlag := make(chan bool)
	connCloseFlag := make(chan bool)
	port := make(chan int)
	go func() {
		ln, portNumber := test_init.GetMockVLRListenerAndPort(t)
		port <- portNumber

		netConn, err := ln.Accept()
		assert.NoError(t, err)

		// wait for messages before closing the listener
		<-msgSentFlag

		wconn := sctp.NewSCTPSndRcvInfoWrappedConn(netConn.(*sctp.SCTPConn))
		defer wconn.Close()

		buf := make([]byte, 254)
		n, err := wconn.Read(buf)
		assert.NoError(t, err)

		encodedMsg, err := message.EncodeSGsAPEPSDetachIndication(req)
		assert.NoError(t, err)
		assert.Equal(t, encodedMsg, buf[unsafe.Sizeof(sctp.SndRcvInfo{}):n])

		ln.Close()
		netConn.Close()

		connCloseFlag <- true
	}()

	// wait for initialization of mock listener
	vlrSCTPAddr := servicers.ConstructSCTPAddr(
		servicers.DefaultVLRIPAddress,
		<-port,
	)
	vlrConn, err := servicers.NewSCTPClientConnection(vlrSCTPAddr, nil)
	assert.NoError(t, err)
	err = vlrConn.EstablishConn()
	defer vlrConn.CloseConn()
	assert.NoError(t, err)
	conn := test_init.GetConnToTestFedGWServiceServer(t, vlrConn)
	defer conn.Close()

	client := protos.NewCSFBFedGWServiceClient(conn)
	reply, err := client.EPSDetachInd(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, &orcprotos.Void{}, reply)

	msgSentFlag <- true
	// close the mock listener first before moving on to the next test
	<-connCloseFlag
}
