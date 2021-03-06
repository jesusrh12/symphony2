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
	"github.com/facebookincubator/symphony/pkg/ent/customer"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/service"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where adds a new predicate for the builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetExternalID sets the external_id field.
func (cu *CustomerUpdate) SetExternalID(s string) *CustomerUpdate {
	cu.mutation.SetExternalID(s)
	return cu
}

// SetNillableExternalID sets the external_id field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableExternalID(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetExternalID(*s)
	}
	return cu
}

// ClearExternalID clears the value of external_id.
func (cu *CustomerUpdate) ClearExternalID() *CustomerUpdate {
	cu.mutation.ClearExternalID()
	return cu
}

// AddServiceIDs adds the services edge to Service by ids.
func (cu *CustomerUpdate) AddServiceIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddServiceIDs(ids...)
	return cu
}

// AddServices adds the services edges to Service.
func (cu *CustomerUpdate) AddServices(s ...*Service) *CustomerUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddServiceIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearServices clears all "services" edges to type Service.
func (cu *CustomerUpdate) ClearServices() *CustomerUpdate {
	cu.mutation.ClearServices()
	return cu
}

// RemoveServiceIDs removes the services edge to Service by ids.
func (cu *CustomerUpdate) RemoveServiceIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveServiceIDs(ids...)
	return cu
}

// RemoveServices removes services edges to Service.
func (cu *CustomerUpdate) RemoveServices(s ...*Service) *CustomerUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveServiceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CustomerUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := customer.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CustomerUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cu.mutation.ExternalID(); ok {
		if err := customer.ExternalIDValidator(v); err != nil {
			return &ValidationError{Name: "external_id", err: fmt.Errorf("ent: validator failed for field \"external_id\": %w", err)}
		}
	}
	return nil
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if value, ok := cu.mutation.ExternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldExternalID,
		})
	}
	if cu.mutation.ExternalIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldExternalID,
		})
	}
	if cu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   customer.ServicesTable,
			Columns: customer.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedServicesIDs(); len(nodes) > 0 && !cu.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   customer.ServicesTable,
			Columns: customer.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   customer.ServicesTable,
			Columns: customer.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// SetName sets the name field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetExternalID sets the external_id field.
func (cuo *CustomerUpdateOne) SetExternalID(s string) *CustomerUpdateOne {
	cuo.mutation.SetExternalID(s)
	return cuo
}

// SetNillableExternalID sets the external_id field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableExternalID(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetExternalID(*s)
	}
	return cuo
}

// ClearExternalID clears the value of external_id.
func (cuo *CustomerUpdateOne) ClearExternalID() *CustomerUpdateOne {
	cuo.mutation.ClearExternalID()
	return cuo
}

// AddServiceIDs adds the services edge to Service by ids.
func (cuo *CustomerUpdateOne) AddServiceIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddServiceIDs(ids...)
	return cuo
}

// AddServices adds the services edges to Service.
func (cuo *CustomerUpdateOne) AddServices(s ...*Service) *CustomerUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddServiceIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearServices clears all "services" edges to type Service.
func (cuo *CustomerUpdateOne) ClearServices() *CustomerUpdateOne {
	cuo.mutation.ClearServices()
	return cuo
}

// RemoveServiceIDs removes the services edge to Service by ids.
func (cuo *CustomerUpdateOne) RemoveServiceIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveServiceIDs(ids...)
	return cuo
}

// RemoveServices removes services edges to Service.
func (cuo *CustomerUpdateOne) RemoveServices(s ...*Service) *CustomerUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveServiceIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CustomerUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := customer.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CustomerUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.ExternalID(); ok {
		if err := customer.ExternalIDValidator(v); err != nil {
			return &ValidationError{Name: "external_id", err: fmt.Errorf("ent: validator failed for field \"external_id\": %w", err)}
		}
	}
	return nil
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Customer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if value, ok := cuo.mutation.ExternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldExternalID,
		})
	}
	if cuo.mutation.ExternalIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldExternalID,
		})
	}
	if cuo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   customer.ServicesTable,
			Columns: customer.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedServicesIDs(); len(nodes) > 0 && !cuo.mutation.ServicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   customer.ServicesTable,
			Columns: customer.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   customer.ServicesTable,
			Columns: customer.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
