// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package viewertest

import (
	"context"
	"testing"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/enttest"
	"github.com/facebookincubator/symphony/graph/ent/migrate"
	"github.com/facebookincubator/symphony/graph/viewer"
	"github.com/facebookincubator/symphony/pkg/testdb"
	"github.com/stretchr/testify/require"
)

type options struct {
	tenant   string
	user     string
	role     string
	features []string
}

// Option enables viewer customization.
type Option func(*options)

// WithTenant overrides default tenant name.
func WithTenant(tenant string) Option {
	return func(o *options) {
		o.tenant = tenant
	}
}

// WithUser overrides default user name.
func WithUser(user string) Option {
	return func(o *options) {
		o.user = user
	}
}

// WithRole overrides default role.
func WithRole(role string) Option {
	return func(o *options) {
		o.role = role
	}
}

// WithFeatures overrides default feature set.
func WithFeatures(features ...string) Option {
	return func(o *options) {
		o.features = features
	}
}

// NewContext returns viewer context for tests.
func NewContext(parent context.Context, c *ent.Client, opts ...Option) context.Context {
	o := &options{
		tenant:   DefaultTenant,
		user:     DefaultUser,
		role:     DefaultRole,
		features: []string{viewer.FeatureReadOnly},
	}
	for _, opt := range opts {
		opt(o)
	}
	ctx := ent.NewContext(parent, c)
	u := viewer.MustGetOrCreateUser(ctx, o.user, o.role)
	v := viewer.New(o.tenant, u, viewer.WithFeatures(o.features...))
	return viewer.NewContext(ctx, v)
}

// NewTestClient creates an ent test client
func NewTestClient(t *testing.T) *ent.Client {
	db, name, err := testdb.Open()
	require.NoError(t, err)
	db.SetMaxOpenConns(1)
	drv := sql.OpenDB(name, db)
	return enttest.NewClient(t,
		enttest.WithOptions(ent.Driver(drv)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	)
}
