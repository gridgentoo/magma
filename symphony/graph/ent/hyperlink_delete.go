// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/hyperlink"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
)

// HyperlinkDelete is the builder for deleting a Hyperlink entity.
type HyperlinkDelete struct {
	config
	hooks      []Hook
	mutation   *HyperlinkMutation
	predicates []predicate.Hyperlink
}

// Where adds a new predicate to the delete builder.
func (hd *HyperlinkDelete) Where(ps ...predicate.Hyperlink) *HyperlinkDelete {
	hd.predicates = append(hd.predicates, ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HyperlinkDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hd.hooks) == 0 {
		affected, err = hd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HyperlinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hd.mutation = mutation
			affected, err = hd.sqlExec(ctx)
			return affected, err
		})
		for i := len(hd.hooks) - 1; i >= 0; i-- {
			mut = hd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HyperlinkDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HyperlinkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hyperlink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hyperlink.FieldID,
			},
		},
	}
	if ps := hd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
}

// HyperlinkDeleteOne is the builder for deleting a single Hyperlink entity.
type HyperlinkDeleteOne struct {
	hd *HyperlinkDelete
}

// Exec executes the deletion query.
func (hdo *HyperlinkDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hyperlink.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HyperlinkDeleteOne) ExecX(ctx context.Context) {
	hdo.hd.ExecX(ctx)
}
