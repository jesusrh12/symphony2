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
	"github.com/facebookincubator/symphony/pkg/ent/kqitarget"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
)

// KqiTargetDelete is the builder for deleting a KqiTarget entity.
type KqiTargetDelete struct {
	config
	hooks    []Hook
	mutation *KqiTargetMutation
}

// Where adds a new predicate to the delete builder.
func (ktd *KqiTargetDelete) Where(ps ...predicate.KqiTarget) *KqiTargetDelete {
	ktd.mutation.predicates = append(ktd.mutation.predicates, ps...)
	return ktd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ktd *KqiTargetDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ktd.hooks) == 0 {
		affected, err = ktd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KqiTargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ktd.mutation = mutation
			affected, err = ktd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ktd.hooks) - 1; i >= 0; i-- {
			mut = ktd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ktd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ktd *KqiTargetDelete) ExecX(ctx context.Context) int {
	n, err := ktd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ktd *KqiTargetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: kqitarget.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kqitarget.FieldID,
			},
		},
	}
	if ps := ktd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ktd.driver, _spec)
}

// KqiTargetDeleteOne is the builder for deleting a single KqiTarget entity.
type KqiTargetDeleteOne struct {
	ktd *KqiTargetDelete
}

// Exec executes the deletion query.
func (ktdo *KqiTargetDeleteOne) Exec(ctx context.Context) error {
	n, err := ktdo.ktd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{kqitarget.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ktdo *KqiTargetDeleteOne) ExecX(ctx context.Context) {
	ktdo.ktd.ExecX(ctx)
}
