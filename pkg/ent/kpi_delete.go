// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/symphony/pkg/ent/kpi"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
)

// KpiDelete is the builder for deleting a Kpi entity.
type KpiDelete struct {
	config
	hooks    []Hook
	mutation *KpiMutation
}

// Where adds a new predicate to the delete builder.
func (kd *KpiDelete) Where(ps ...predicate.Kpi) *KpiDelete {
	kd.mutation.predicates = append(kd.mutation.predicates, ps...)
	return kd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kd *KpiDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kd.hooks) == 0 {
		affected, err = kd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KpiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kd.mutation = mutation
			affected, err = kd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kd.hooks) - 1; i >= 0; i-- {
			mut = kd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (kd *KpiDelete) ExecX(ctx context.Context) int {
	n, err := kd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kd *KpiDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: kpi.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kpi.FieldID,
			},
		},
	}
	if ps := kd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, kd.driver, _spec)
}

// KpiDeleteOne is the builder for deleting a single Kpi entity.
type KpiDeleteOne struct {
	kd *KpiDelete
}

// Exec executes the deletion query.
func (kdo *KpiDeleteOne) Exec(ctx context.Context) error {
	n, err := kdo.kd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{kpi.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kdo *KpiDeleteOne) ExecX(ctx context.Context) {
	kdo.kd.ExecX(ctx)
}
