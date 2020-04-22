/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/plugin"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/dispatcher/gateway_registry"
	"magma/orc8r/lib/go/registry"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "service303_cli",
	Short: "Management CLI for Service303",
}

var services []string
var gwServices []gateway_registry.GwServiceType

var isGatewayServiceQuery bool
var hardwareID string
var networkID string
var gatewayID string

func main() {
	flag.Parse()
	plugin.LoadAllPluginsFatalOnError(&plugin.DefaultOrchestratorPluginLoader{})

	services = registry.ListAllServices()
	gwServices = gateway_registry.ListAllGwServices()

	rootCmd.PersistentFlags().BoolVar(&isGatewayServiceQuery, "gateway-service", false, "query a gateway service")
	rootCmd.PersistentFlags().StringVar(&hardwareID, "hwid", "", "the hardware id of the gateway to send command to")
	rootCmd.PersistentFlags().StringVar(&networkID, "network", "", "the network id")
	rootCmd.PersistentFlags().StringVar(&gatewayID, "gateway", "", "the gateway id")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(2)
	}
}

func validateGlobalFlags() error {
	if !isGatewayServiceQuery && hardwareID == "" && networkID == "" && gatewayID == "" {
		return nil
	}
	if isGatewayServiceQuery {
		if hardwareID != "" && networkID == "" && gatewayID == "" {
			return nil
		}
		if hardwareID == "" && networkID != "" && gatewayID != "" {
			return nil
		}
	}
	return fmt.Errorf("invalid flag combination")
}

func setHwIdFlag() error {
	if networkID == "" || gatewayID == "" {
		return nil
	}
	var err error
	hardwareID, err = configurator.GetPhysicalIDOfEntity(networkID, orc8r.MagmadGatewayType, gatewayID)
	if err != nil {
		return err
	}
	return nil
}

func isValidService(service string, services []string) bool {
	for _, serv := range services {
		if serv == service {
			return true
		}
	}
	return false
}

func isValidGwService(service gateway_registry.GwServiceType, services []gateway_registry.GwServiceType) bool {
	for _, serv := range services {
		if serv == service {
			return true
		}
	}
	return false
}
