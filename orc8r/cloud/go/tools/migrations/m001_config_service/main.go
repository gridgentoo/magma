/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package main

import (
	"flag"
	"log"

	"magma/orc8r/cloud/go/tools/migrations"
	"magma/orc8r/cloud/go/tools/migrations/m001_config_service/migration"

	_ "github.com/lib/pq"
)

func main() {
	postValidate := flag.Bool("validate", false, "Set this flag to run validation after a completed migration.")
	flag.Parse()

	dbDriver := migrations.GetEnvWithDefault("SQL_DRIVER", "postgres")
	dbSource := migrations.GetEnvWithDefault("DATABASE_SOURCE", "dbname=magma user=magma password=magma host=192.168.80.20")

	if !*postValidate {
		err := migration.Migrate(dbDriver, dbSource)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := migration.Validate(dbDriver, dbSource)
		if err != nil {
			log.Fatal(err)
		}
	}
}
