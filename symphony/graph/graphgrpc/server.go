// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphgrpc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/facebookincubator/symphony/graph/graphactions"
	"github.com/facebookincubator/symphony/graph/viewer"
	"github.com/facebookincubator/symphony/pkg/actions"
	"github.com/facebookincubator/symphony/pkg/actions/executor"
	"github.com/facebookincubator/symphony/pkg/grpc-middleware/sqltx"
	"github.com/facebookincubator/symphony/pkg/log"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newServer(tenancy viewer.Tenancy, db *sql.DB, logger log.Logger, registry *executor.Registry) (*grpc.Server, func(), error) {
	grpc_zap.ReplaceGrpcLoggerV2(logger.Background())
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger.Background()),
			grpc_recovery.UnaryServerInterceptor(),
			sqltx.UnaryServerInterceptor(db),
		)),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)
	RegisterTenantServiceServer(s,
		NewTenantService(func(ctx context.Context) ExecQueryer {
			return sqltx.FromContext(ctx)
		}),
	)
	RegisterActionsAlertServiceServer(s,
		NewActionsAlertService(func(ctx context.Context, tenantID string) (*actions.Client, error) {
			entClient, err := tenancy.ClientFor(ctx, tenantID)
			if err != nil {
				return nil, err
			}
			dataLoader := graphactions.EntDataLoader{
				Client: entClient,
			}
			onError := func(ctx context.Context, err error) {
				logger.For(ctx).Error("error executing action", zap.Error(err))
			}
			exc := &executor.Executor{
				Registry:   registry,
				DataLoader: dataLoader,
				OnError:    onError,
			}
			return actions.NewClient(exc), nil
		}),
	)
	RegisterUserServiceServer(s, NewUserService(tenancy.ClientFor))

	reflection.Register(s)
	err := view.Register(ocgrpc.DefaultServerViews...)
	if err != nil {
		return nil, nil, fmt.Errorf("registering grpc views: %w", err)
	}
	return s, func() { view.Unregister(ocgrpc.DefaultServerViews...) }, nil
}
