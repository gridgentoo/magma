/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMemorySubscriberStore(t *testing.T) {
	testSuite := new(SubscriberStoreTestSuite)
	testSuite.createStore = func() SubscriberStore {
		return NewMemorySubscriberStore()
	}
	suite.Run(t, testSuite)
}
