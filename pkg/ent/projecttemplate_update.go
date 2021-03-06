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
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/projecttemplate"
	"github.com/facebookincubator/symphony/pkg/ent/projecttype"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/workorderdefinition"
)

// ProjectTemplateUpdate is the builder for updating ProjectTemplate entities.
type ProjectTemplateUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectTemplateMutation
}

// Where adds a new predicate for the builder.
func (ptu *ProjectTemplateUpdate) Where(ps ...predicate.ProjectTemplate) *ProjectTemplateUpdate {
	ptu.mutation.predicates = append(ptu.mutation.predicates, ps...)
	return ptu
}

// SetName sets the name field.
func (ptu *ProjectTemplateUpdate) SetName(s string) *ProjectTemplateUpdate {
	ptu.mutation.SetName(s)
	return ptu
}

// SetDescription sets the description field.
func (ptu *ProjectTemplateUpdate) SetDescription(s string) *ProjectTemplateUpdate {
	ptu.mutation.SetDescription(s)
	return ptu
}

// SetNillableDescription sets the description field if the given value is not nil.
func (ptu *ProjectTemplateUpdate) SetNillableDescription(s *string) *ProjectTemplateUpdate {
	if s != nil {
		ptu.SetDescription(*s)
	}
	return ptu
}

// ClearDescription clears the value of description.
func (ptu *ProjectTemplateUpdate) ClearDescription() *ProjectTemplateUpdate {
	ptu.mutation.ClearDescription()
	return ptu
}

// AddPropertyIDs adds the properties edge to PropertyType by ids.
func (ptu *ProjectTemplateUpdate) AddPropertyIDs(ids ...int) *ProjectTemplateUpdate {
	ptu.mutation.AddPropertyIDs(ids...)
	return ptu
}

// AddProperties adds the properties edges to PropertyType.
func (ptu *ProjectTemplateUpdate) AddProperties(p ...*PropertyType) *ProjectTemplateUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptu.AddPropertyIDs(ids...)
}

// AddWorkOrderIDs adds the work_orders edge to WorkOrderDefinition by ids.
func (ptu *ProjectTemplateUpdate) AddWorkOrderIDs(ids ...int) *ProjectTemplateUpdate {
	ptu.mutation.AddWorkOrderIDs(ids...)
	return ptu
}

// AddWorkOrders adds the work_orders edges to WorkOrderDefinition.
func (ptu *ProjectTemplateUpdate) AddWorkOrders(w ...*WorkOrderDefinition) *ProjectTemplateUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ptu.AddWorkOrderIDs(ids...)
}

// SetTypeID sets the type edge to ProjectType by id.
func (ptu *ProjectTemplateUpdate) SetTypeID(id int) *ProjectTemplateUpdate {
	ptu.mutation.SetTypeID(id)
	return ptu
}

// SetNillableTypeID sets the type edge to ProjectType by id if the given value is not nil.
func (ptu *ProjectTemplateUpdate) SetNillableTypeID(id *int) *ProjectTemplateUpdate {
	if id != nil {
		ptu = ptu.SetTypeID(*id)
	}
	return ptu
}

// SetType sets the type edge to ProjectType.
func (ptu *ProjectTemplateUpdate) SetType(p *ProjectType) *ProjectTemplateUpdate {
	return ptu.SetTypeID(p.ID)
}

// Mutation returns the ProjectTemplateMutation object of the builder.
func (ptu *ProjectTemplateUpdate) Mutation() *ProjectTemplateMutation {
	return ptu.mutation
}

// ClearProperties clears all "properties" edges to type PropertyType.
func (ptu *ProjectTemplateUpdate) ClearProperties() *ProjectTemplateUpdate {
	ptu.mutation.ClearProperties()
	return ptu
}

// RemovePropertyIDs removes the properties edge to PropertyType by ids.
func (ptu *ProjectTemplateUpdate) RemovePropertyIDs(ids ...int) *ProjectTemplateUpdate {
	ptu.mutation.RemovePropertyIDs(ids...)
	return ptu
}

// RemoveProperties removes properties edges to PropertyType.
func (ptu *ProjectTemplateUpdate) RemoveProperties(p ...*PropertyType) *ProjectTemplateUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptu.RemovePropertyIDs(ids...)
}

// ClearWorkOrders clears all "work_orders" edges to type WorkOrderDefinition.
func (ptu *ProjectTemplateUpdate) ClearWorkOrders() *ProjectTemplateUpdate {
	ptu.mutation.ClearWorkOrders()
	return ptu
}

// RemoveWorkOrderIDs removes the work_orders edge to WorkOrderDefinition by ids.
func (ptu *ProjectTemplateUpdate) RemoveWorkOrderIDs(ids ...int) *ProjectTemplateUpdate {
	ptu.mutation.RemoveWorkOrderIDs(ids...)
	return ptu
}

// RemoveWorkOrders removes work_orders edges to WorkOrderDefinition.
func (ptu *ProjectTemplateUpdate) RemoveWorkOrders(w ...*WorkOrderDefinition) *ProjectTemplateUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ptu.RemoveWorkOrderIDs(ids...)
}

// ClearType clears the "type" edge to type ProjectType.
func (ptu *ProjectTemplateUpdate) ClearType() *ProjectTemplateUpdate {
	ptu.mutation.ClearType()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *ProjectTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ptu.defaults()
	if len(ptu.hooks) == 0 {
		if err = ptu.check(); err != nil {
			return 0, err
		}
		affected, err = ptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptu.check(); err != nil {
				return 0, err
			}
			ptu.mutation = mutation
			affected, err = ptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptu.hooks) - 1; i >= 0; i-- {
			mut = ptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *ProjectTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *ProjectTemplateUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *ProjectTemplateUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptu *ProjectTemplateUpdate) defaults() {
	if _, ok := ptu.mutation.UpdateTime(); !ok {
		v := projecttemplate.UpdateDefaultUpdateTime()
		ptu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *ProjectTemplateUpdate) check() error {
	if v, ok := ptu.mutation.Name(); ok {
		if err := projecttemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ptu *ProjectTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttemplate.Table,
			Columns: projecttemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projecttemplate.FieldID,
			},
		},
	}
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecttemplate.FieldUpdateTime,
		})
	}
	if value, ok := ptu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttemplate.FieldName,
		})
	}
	if value, ok := ptu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttemplate.FieldDescription,
		})
	}
	if ptu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: projecttemplate.FieldDescription,
		})
	}
	if ptu.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.PropertiesTable,
			Columns: []string{projecttemplate.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !ptu.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.PropertiesTable,
			Columns: []string{projecttemplate.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.PropertiesTable,
			Columns: []string{projecttemplate.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptu.mutation.WorkOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.WorkOrdersTable,
			Columns: []string{projecttemplate.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.RemovedWorkOrdersIDs(); len(nodes) > 0 && !ptu.mutation.WorkOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.WorkOrdersTable,
			Columns: []string{projecttemplate.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.WorkOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.WorkOrdersTable,
			Columns: []string{projecttemplate.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projecttemplate.TypeTable,
			Columns: []string{projecttemplate.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projecttype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projecttemplate.TypeTable,
			Columns: []string{projecttemplate.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projecttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttemplate.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProjectTemplateUpdateOne is the builder for updating a single ProjectTemplate entity.
type ProjectTemplateUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProjectTemplateMutation
}

// SetName sets the name field.
func (ptuo *ProjectTemplateUpdateOne) SetName(s string) *ProjectTemplateUpdateOne {
	ptuo.mutation.SetName(s)
	return ptuo
}

// SetDescription sets the description field.
func (ptuo *ProjectTemplateUpdateOne) SetDescription(s string) *ProjectTemplateUpdateOne {
	ptuo.mutation.SetDescription(s)
	return ptuo
}

// SetNillableDescription sets the description field if the given value is not nil.
func (ptuo *ProjectTemplateUpdateOne) SetNillableDescription(s *string) *ProjectTemplateUpdateOne {
	if s != nil {
		ptuo.SetDescription(*s)
	}
	return ptuo
}

// ClearDescription clears the value of description.
func (ptuo *ProjectTemplateUpdateOne) ClearDescription() *ProjectTemplateUpdateOne {
	ptuo.mutation.ClearDescription()
	return ptuo
}

// AddPropertyIDs adds the properties edge to PropertyType by ids.
func (ptuo *ProjectTemplateUpdateOne) AddPropertyIDs(ids ...int) *ProjectTemplateUpdateOne {
	ptuo.mutation.AddPropertyIDs(ids...)
	return ptuo
}

// AddProperties adds the properties edges to PropertyType.
func (ptuo *ProjectTemplateUpdateOne) AddProperties(p ...*PropertyType) *ProjectTemplateUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptuo.AddPropertyIDs(ids...)
}

// AddWorkOrderIDs adds the work_orders edge to WorkOrderDefinition by ids.
func (ptuo *ProjectTemplateUpdateOne) AddWorkOrderIDs(ids ...int) *ProjectTemplateUpdateOne {
	ptuo.mutation.AddWorkOrderIDs(ids...)
	return ptuo
}

// AddWorkOrders adds the work_orders edges to WorkOrderDefinition.
func (ptuo *ProjectTemplateUpdateOne) AddWorkOrders(w ...*WorkOrderDefinition) *ProjectTemplateUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ptuo.AddWorkOrderIDs(ids...)
}

// SetTypeID sets the type edge to ProjectType by id.
func (ptuo *ProjectTemplateUpdateOne) SetTypeID(id int) *ProjectTemplateUpdateOne {
	ptuo.mutation.SetTypeID(id)
	return ptuo
}

// SetNillableTypeID sets the type edge to ProjectType by id if the given value is not nil.
func (ptuo *ProjectTemplateUpdateOne) SetNillableTypeID(id *int) *ProjectTemplateUpdateOne {
	if id != nil {
		ptuo = ptuo.SetTypeID(*id)
	}
	return ptuo
}

// SetType sets the type edge to ProjectType.
func (ptuo *ProjectTemplateUpdateOne) SetType(p *ProjectType) *ProjectTemplateUpdateOne {
	return ptuo.SetTypeID(p.ID)
}

// Mutation returns the ProjectTemplateMutation object of the builder.
func (ptuo *ProjectTemplateUpdateOne) Mutation() *ProjectTemplateMutation {
	return ptuo.mutation
}

// ClearProperties clears all "properties" edges to type PropertyType.
func (ptuo *ProjectTemplateUpdateOne) ClearProperties() *ProjectTemplateUpdateOne {
	ptuo.mutation.ClearProperties()
	return ptuo
}

// RemovePropertyIDs removes the properties edge to PropertyType by ids.
func (ptuo *ProjectTemplateUpdateOne) RemovePropertyIDs(ids ...int) *ProjectTemplateUpdateOne {
	ptuo.mutation.RemovePropertyIDs(ids...)
	return ptuo
}

// RemoveProperties removes properties edges to PropertyType.
func (ptuo *ProjectTemplateUpdateOne) RemoveProperties(p ...*PropertyType) *ProjectTemplateUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ptuo.RemovePropertyIDs(ids...)
}

// ClearWorkOrders clears all "work_orders" edges to type WorkOrderDefinition.
func (ptuo *ProjectTemplateUpdateOne) ClearWorkOrders() *ProjectTemplateUpdateOne {
	ptuo.mutation.ClearWorkOrders()
	return ptuo
}

// RemoveWorkOrderIDs removes the work_orders edge to WorkOrderDefinition by ids.
func (ptuo *ProjectTemplateUpdateOne) RemoveWorkOrderIDs(ids ...int) *ProjectTemplateUpdateOne {
	ptuo.mutation.RemoveWorkOrderIDs(ids...)
	return ptuo
}

// RemoveWorkOrders removes work_orders edges to WorkOrderDefinition.
func (ptuo *ProjectTemplateUpdateOne) RemoveWorkOrders(w ...*WorkOrderDefinition) *ProjectTemplateUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ptuo.RemoveWorkOrderIDs(ids...)
}

// ClearType clears the "type" edge to type ProjectType.
func (ptuo *ProjectTemplateUpdateOne) ClearType() *ProjectTemplateUpdateOne {
	ptuo.mutation.ClearType()
	return ptuo
}

// Save executes the query and returns the updated entity.
func (ptuo *ProjectTemplateUpdateOne) Save(ctx context.Context) (*ProjectTemplate, error) {
	var (
		err  error
		node *ProjectTemplate
	)
	ptuo.defaults()
	if len(ptuo.hooks) == 0 {
		if err = ptuo.check(); err != nil {
			return nil, err
		}
		node, err = ptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptuo.check(); err != nil {
				return nil, err
			}
			ptuo.mutation = mutation
			node, err = ptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptuo.hooks) - 1; i >= 0; i-- {
			mut = ptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *ProjectTemplateUpdateOne) SaveX(ctx context.Context) *ProjectTemplate {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *ProjectTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *ProjectTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptuo *ProjectTemplateUpdateOne) defaults() {
	if _, ok := ptuo.mutation.UpdateTime(); !ok {
		v := projecttemplate.UpdateDefaultUpdateTime()
		ptuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *ProjectTemplateUpdateOne) check() error {
	if v, ok := ptuo.mutation.Name(); ok {
		if err := projecttemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ptuo *ProjectTemplateUpdateOne) sqlSave(ctx context.Context) (_node *ProjectTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttemplate.Table,
			Columns: projecttemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projecttemplate.FieldID,
			},
		},
	}
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProjectTemplate.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ptuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projecttemplate.FieldUpdateTime,
		})
	}
	if value, ok := ptuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttemplate.FieldName,
		})
	}
	if value, ok := ptuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projecttemplate.FieldDescription,
		})
	}
	if ptuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: projecttemplate.FieldDescription,
		})
	}
	if ptuo.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.PropertiesTable,
			Columns: []string{projecttemplate.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !ptuo.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.PropertiesTable,
			Columns: []string{projecttemplate.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.PropertiesTable,
			Columns: []string{projecttemplate.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptuo.mutation.WorkOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.WorkOrdersTable,
			Columns: []string{projecttemplate.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.RemovedWorkOrdersIDs(); len(nodes) > 0 && !ptuo.mutation.WorkOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.WorkOrdersTable,
			Columns: []string{projecttemplate.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.WorkOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projecttemplate.WorkOrdersTable,
			Columns: []string{projecttemplate.WorkOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorderdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projecttemplate.TypeTable,
			Columns: []string{projecttemplate.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projecttype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projecttemplate.TypeTable,
			Columns: []string{projecttemplate.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projecttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectTemplate{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttemplate.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
