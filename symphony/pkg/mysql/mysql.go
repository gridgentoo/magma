// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/telemetry"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// Open new connection and start stats recorder.
func Open(dsn string) *sql.DB {
	return sql.OpenDB(connector{dsn})
}

// RecordStats records database statistics for provided sql.DB.
func RecordStats(db *sql.DB) func() {
	return ocsql.RecordStats(db, 10*time.Second)
}

// SetLogger is used to set the logger for critical errors.
func SetLogger(logger log.Logger) {
	const lvl = zap.ErrorLevel
	l, _ := zap.NewStdLogAt(
		logger.Background().
			WithOptions(zap.AddStacktrace(lvl)).
			With(zap.String("pkg", "mysql")),
		lvl,
	)
	_ = mysql.SetLogger(l)
}

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
		ocsql.WithSampler(
			telemetry.WithoutNameSampler("sql:prepare"),
		),
	)
}

// DefaultViews are predefined views for opencensus metrics.
var DefaultViews = ocsql.DefaultViews
