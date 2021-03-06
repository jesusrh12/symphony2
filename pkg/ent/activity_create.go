// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/symphony/pkg/ent/activity"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/ent/workorder"
)

// ActivityCreate is the builder for creating a Activity entity.
type ActivityCreate struct {
	config
	mutation *ActivityMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (ac *ActivityCreate) SetCreateTime(t time.Time) *ActivityCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (ac *ActivityCreate) SetNillableCreateTime(t *time.Time) *ActivityCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the update_time field.
func (ac *ActivityCreate) SetUpdateTime(t time.Time) *ActivityCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (ac *ActivityCreate) SetNillableUpdateTime(t *time.Time) *ActivityCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetActivityType sets the activity_type field.
func (ac *ActivityCreate) SetActivityType(at activity.ActivityType) *ActivityCreate {
	ac.mutation.SetActivityType(at)
	return ac
}

// SetIsCreate sets the is_create field.
func (ac *ActivityCreate) SetIsCreate(b bool) *ActivityCreate {
	ac.mutation.SetIsCreate(b)
	return ac
}

// SetNillableIsCreate sets the is_create field if the given value is not nil.
func (ac *ActivityCreate) SetNillableIsCreate(b *bool) *ActivityCreate {
	if b != nil {
		ac.SetIsCreate(*b)
	}
	return ac
}

// SetOldValue sets the old_value field.
func (ac *ActivityCreate) SetOldValue(s string) *ActivityCreate {
	ac.mutation.SetOldValue(s)
	return ac
}

// SetNillableOldValue sets the old_value field if the given value is not nil.
func (ac *ActivityCreate) SetNillableOldValue(s *string) *ActivityCreate {
	if s != nil {
		ac.SetOldValue(*s)
	}
	return ac
}

// SetNewValue sets the new_value field.
func (ac *ActivityCreate) SetNewValue(s string) *ActivityCreate {
	ac.mutation.SetNewValue(s)
	return ac
}

// SetNillableNewValue sets the new_value field if the given value is not nil.
func (ac *ActivityCreate) SetNillableNewValue(s *string) *ActivityCreate {
	if s != nil {
		ac.SetNewValue(*s)
	}
	return ac
}

// SetClockDetails sets the clock_details field.
func (ac *ActivityCreate) SetClockDetails(ad activity.ClockDetails) *ActivityCreate {
	ac.mutation.SetClockDetails(ad)
	return ac
}

// SetNillableClockDetails sets the clock_details field if the given value is not nil.
func (ac *ActivityCreate) SetNillableClockDetails(ad *activity.ClockDetails) *ActivityCreate {
	if ad != nil {
		ac.SetClockDetails(*ad)
	}
	return ac
}

// SetAuthorID sets the author edge to User by id.
func (ac *ActivityCreate) SetAuthorID(id int) *ActivityCreate {
	ac.mutation.SetAuthorID(id)
	return ac
}

// SetNillableAuthorID sets the author edge to User by id if the given value is not nil.
func (ac *ActivityCreate) SetNillableAuthorID(id *int) *ActivityCreate {
	if id != nil {
		ac = ac.SetAuthorID(*id)
	}
	return ac
}

// SetAuthor sets the author edge to User.
func (ac *ActivityCreate) SetAuthor(u *User) *ActivityCreate {
	return ac.SetAuthorID(u.ID)
}

// SetWorkOrderID sets the work_order edge to WorkOrder by id.
func (ac *ActivityCreate) SetWorkOrderID(id int) *ActivityCreate {
	ac.mutation.SetWorkOrderID(id)
	return ac
}

// SetNillableWorkOrderID sets the work_order edge to WorkOrder by id if the given value is not nil.
func (ac *ActivityCreate) SetNillableWorkOrderID(id *int) *ActivityCreate {
	if id != nil {
		ac = ac.SetWorkOrderID(*id)
	}
	return ac
}

// SetWorkOrder sets the work_order edge to WorkOrder.
func (ac *ActivityCreate) SetWorkOrder(w *WorkOrder) *ActivityCreate {
	return ac.SetWorkOrderID(w.ID)
}

// Mutation returns the ActivityMutation object of the builder.
func (ac *ActivityCreate) Mutation() *ActivityMutation {
	return ac.mutation
}

// Save creates the Activity in the database.
func (ac *ActivityCreate) Save(ctx context.Context) (*Activity, error) {
	var (
		err  error
		node *Activity
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ActivityCreate) SaveX(ctx context.Context) *Activity {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ac *ActivityCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := activity.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := activity.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
	if _, ok := ac.mutation.IsCreate(); !ok {
		v := activity.DefaultIsCreate
		ac.mutation.SetIsCreate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ActivityCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := ac.mutation.ActivityType(); !ok {
		return &ValidationError{Name: "activity_type", err: errors.New("ent: missing required field \"activity_type\"")}
	}
	if v, ok := ac.mutation.ActivityType(); ok {
		if err := activity.ActivityTypeValidator(v); err != nil {
			return &ValidationError{Name: "activity_type", err: fmt.Errorf("ent: validator failed for field \"activity_type\": %w", err)}
		}
	}
	if _, ok := ac.mutation.IsCreate(); !ok {
		return &ValidationError{Name: "is_create", err: errors.New("ent: missing required field \"is_create\"")}
	}
	return nil
}

func (ac *ActivityCreate) sqlSave(ctx context.Context) (*Activity, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *ActivityCreate) createSpec() (*Activity, *sqlgraph.CreateSpec) {
	var (
		_node = &Activity{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: activity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activity.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activity.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activity.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.ActivityType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: activity.FieldActivityType,
		})
		_node.ActivityType = value
	}
	if value, ok := ac.mutation.IsCreate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: activity.FieldIsCreate,
		})
		_node.IsCreate = value
	}
	if value, ok := ac.mutation.OldValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldOldValue,
		})
		_node.OldValue = value
	}
	if value, ok := ac.mutation.NewValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldNewValue,
		})
		_node.NewValue = value
	}
	if value, ok := ac.mutation.ClockDetails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: activity.FieldClockDetails,
		})
		_node.ClockDetails = value
	}
	if nodes := ac.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   activity.AuthorTable,
			Columns: []string{activity.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.WorkOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activity.WorkOrderTable,
			Columns: []string{activity.WorkOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActivityCreateBulk is the builder for creating a bulk of Activity entities.
type ActivityCreateBulk struct {
	config
	builders []*ActivityCreate
}

// Save creates the Activity entities in the database.
func (acb *ActivityCreateBulk) Save(ctx context.Context) ([]*Activity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Activity, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActivityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (acb *ActivityCreateBulk) SaveX(ctx context.Context) []*Activity {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
