/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package main

import (
	"flag"
	"time"

	"magma/feg/gateway/registry"
	"magma/feg/gateway/services/gateway_health/health_manager"
	"magma/orc8r/lib/go/service"

	"github.com/golang/glog"
)

func init() {
	flag.Parse()
}

func main() {
	// Create the service
	srv, err := service.NewServiceWithOptions(registry.ModuleName, registry.HEALTH)
	if err != nil {
		glog.Fatalf("Error creating HEALTH service: %s", err)
	}

	cloudReg := registry.Get()
	healthCfg := health_manager.GetHealthConfig()
	healthManager := health_manager.NewHealthManager(cloudReg, healthCfg)
	// Run Health Collection Loop
	go func() {
		for {
			<-time.After(time.Duration(healthCfg.UpdateIntervalSecs) * time.Second)
			healthManager.SendHealthUpdate()
		}
	}()

	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running service: %s", err)
	}
}
