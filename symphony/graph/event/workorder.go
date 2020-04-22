// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/hook"
	"github.com/facebookincubator/symphony/graph/ent/workorder"
	"github.com/facebookincubator/symphony/graph/graphql/models"
)

// Work order events.
const (
	WorkOrderAdded = "work_order/added"
	WorkOrderDone  = "work_order/done"
)

// Hook returns the hook which generates events from mutations.
func (e *Eventer) workOrderHook() ent.Hook {
	chain := hook.NewChain(
		e.workOrderCreateHook(),
		e.workOrderUpdateHook(),
		e.workOrderUpdateOneHook(),
	)
	return chain.Hook()
}

func (e *Eventer) workOrderCreateHook() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.WorkOrderFunc(func(ctx context.Context, m *ent.WorkOrderMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return value, err
			}
			e.emit(ctx, WorkOrderAdded, value)
			if value.(*ent.WorkOrder).Status == models.WorkOrderStatusDone.String() {
				e.emit(ctx, WorkOrderDone, value)
			}
			return value, nil
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func (e *Eventer) workOrderUpdateHook() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.WorkOrderFunc(func(ctx context.Context, m *ent.WorkOrderMutation) (ent.Value, error) {
			if status, exists := m.Status(); exists && status == models.WorkOrderStatusDone.String() {
				return nil, errors.New("work order status update to done by predicate not allowed")
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate)
}

func (e *Eventer) workOrderUpdateOneHook() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.WorkOrderFunc(func(ctx context.Context, m *ent.WorkOrderMutation) (ent.Value, error) {
			status, exists := m.Status()
			if !exists || status != models.WorkOrderStatusDone.String() {
				return next.Mutate(ctx, m)
			}
			id, _ := m.ID()
			statusChanged, err := m.Client().WorkOrder.
				Query().
				Where(
					workorder.ID(id),
					workorder.StatusNEQ(
						models.WorkOrderStatusDone.String(),
					),
				).
				Exist(ctx)
			if err != nil {
				return nil, fmt.Errorf("querying work order %d current status: %w", id, err)
			}
			value, err := next.Mutate(ctx, m)
			if err == nil && statusChanged {
				e.emit(ctx, WorkOrderDone, value)
			}
			return value, err
		})
	}
	return hook.On(hk, ent.OpUpdateOne)
}
