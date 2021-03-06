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
	"github.com/facebookincubator/symphony/pkg/ent/equipment"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentcategory"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentpositiondefinition"
	"github.com/facebookincubator/symphony/pkg/ent/equipmenttype"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/serviceendpointdefinition"
)

// EquipmentTypeCreate is the builder for creating a EquipmentType entity.
type EquipmentTypeCreate struct {
	config
	mutation *EquipmentTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (etc *EquipmentTypeCreate) SetCreateTime(t time.Time) *EquipmentTypeCreate {
	etc.mutation.SetCreateTime(t)
	return etc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (etc *EquipmentTypeCreate) SetNillableCreateTime(t *time.Time) *EquipmentTypeCreate {
	if t != nil {
		etc.SetCreateTime(*t)
	}
	return etc
}

// SetUpdateTime sets the update_time field.
func (etc *EquipmentTypeCreate) SetUpdateTime(t time.Time) *EquipmentTypeCreate {
	etc.mutation.SetUpdateTime(t)
	return etc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (etc *EquipmentTypeCreate) SetNillableUpdateTime(t *time.Time) *EquipmentTypeCreate {
	if t != nil {
		etc.SetUpdateTime(*t)
	}
	return etc
}

// SetName sets the name field.
func (etc *EquipmentTypeCreate) SetName(s string) *EquipmentTypeCreate {
	etc.mutation.SetName(s)
	return etc
}

// AddPortDefinitionIDs adds the port_definitions edge to EquipmentPortDefinition by ids.
func (etc *EquipmentTypeCreate) AddPortDefinitionIDs(ids ...int) *EquipmentTypeCreate {
	etc.mutation.AddPortDefinitionIDs(ids...)
	return etc
}

// AddPortDefinitions adds the port_definitions edges to EquipmentPortDefinition.
func (etc *EquipmentTypeCreate) AddPortDefinitions(e ...*EquipmentPortDefinition) *EquipmentTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddPortDefinitionIDs(ids...)
}

// AddPositionDefinitionIDs adds the position_definitions edge to EquipmentPositionDefinition by ids.
func (etc *EquipmentTypeCreate) AddPositionDefinitionIDs(ids ...int) *EquipmentTypeCreate {
	etc.mutation.AddPositionDefinitionIDs(ids...)
	return etc
}

// AddPositionDefinitions adds the position_definitions edges to EquipmentPositionDefinition.
func (etc *EquipmentTypeCreate) AddPositionDefinitions(e ...*EquipmentPositionDefinition) *EquipmentTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddPositionDefinitionIDs(ids...)
}

// AddPropertyTypeIDs adds the property_types edge to PropertyType by ids.
func (etc *EquipmentTypeCreate) AddPropertyTypeIDs(ids ...int) *EquipmentTypeCreate {
	etc.mutation.AddPropertyTypeIDs(ids...)
	return etc
}

// AddPropertyTypes adds the property_types edges to PropertyType.
func (etc *EquipmentTypeCreate) AddPropertyTypes(p ...*PropertyType) *EquipmentTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return etc.AddPropertyTypeIDs(ids...)
}

// AddEquipmentIDs adds the equipment edge to Equipment by ids.
func (etc *EquipmentTypeCreate) AddEquipmentIDs(ids ...int) *EquipmentTypeCreate {
	etc.mutation.AddEquipmentIDs(ids...)
	return etc
}

// AddEquipment adds the equipment edges to Equipment.
func (etc *EquipmentTypeCreate) AddEquipment(e ...*Equipment) *EquipmentTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddEquipmentIDs(ids...)
}

// SetCategoryID sets the category edge to EquipmentCategory by id.
func (etc *EquipmentTypeCreate) SetCategoryID(id int) *EquipmentTypeCreate {
	etc.mutation.SetCategoryID(id)
	return etc
}

// SetNillableCategoryID sets the category edge to EquipmentCategory by id if the given value is not nil.
func (etc *EquipmentTypeCreate) SetNillableCategoryID(id *int) *EquipmentTypeCreate {
	if id != nil {
		etc = etc.SetCategoryID(*id)
	}
	return etc
}

// SetCategory sets the category edge to EquipmentCategory.
func (etc *EquipmentTypeCreate) SetCategory(e *EquipmentCategory) *EquipmentTypeCreate {
	return etc.SetCategoryID(e.ID)
}

// AddServiceEndpointDefinitionIDs adds the service_endpoint_definitions edge to ServiceEndpointDefinition by ids.
func (etc *EquipmentTypeCreate) AddServiceEndpointDefinitionIDs(ids ...int) *EquipmentTypeCreate {
	etc.mutation.AddServiceEndpointDefinitionIDs(ids...)
	return etc
}

// AddServiceEndpointDefinitions adds the service_endpoint_definitions edges to ServiceEndpointDefinition.
func (etc *EquipmentTypeCreate) AddServiceEndpointDefinitions(s ...*ServiceEndpointDefinition) *EquipmentTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return etc.AddServiceEndpointDefinitionIDs(ids...)
}

// Mutation returns the EquipmentTypeMutation object of the builder.
func (etc *EquipmentTypeCreate) Mutation() *EquipmentTypeMutation {
	return etc.mutation
}

// Save creates the EquipmentType in the database.
func (etc *EquipmentTypeCreate) Save(ctx context.Context) (*EquipmentType, error) {
	var (
		err  error
		node *EquipmentType
	)
	etc.defaults()
	if len(etc.hooks) == 0 {
		if err = etc.check(); err != nil {
			return nil, err
		}
		node, err = etc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = etc.check(); err != nil {
				return nil, err
			}
			etc.mutation = mutation
			node, err = etc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(etc.hooks) - 1; i >= 0; i-- {
			mut = etc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EquipmentTypeCreate) SaveX(ctx context.Context) *EquipmentType {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (etc *EquipmentTypeCreate) defaults() {
	if _, ok := etc.mutation.CreateTime(); !ok {
		v := equipmenttype.DefaultCreateTime()
		etc.mutation.SetCreateTime(v)
	}
	if _, ok := etc.mutation.UpdateTime(); !ok {
		v := equipmenttype.DefaultUpdateTime()
		etc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EquipmentTypeCreate) check() error {
	if _, ok := etc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := etc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := etc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (etc *EquipmentTypeCreate) sqlSave(ctx context.Context) (*EquipmentType, error) {
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (etc *EquipmentTypeCreate) createSpec() (*EquipmentType, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentType{config: etc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipmenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmenttype.FieldID,
			},
		}
	)
	if value, ok := etc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: equipmenttype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := etc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: equipmenttype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := etc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmenttype.FieldName,
		})
		_node.Name = value
	}
	if nodes := etc.mutation.PortDefinitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipmenttype.PortDefinitionsTable,
			Columns: []string{equipmenttype.PortDefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentportdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.PositionDefinitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipmenttype.PositionDefinitionsTable,
			Columns: []string{equipmenttype.PositionDefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentpositiondefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.PropertyTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipmenttype.PropertyTypesTable,
			Columns: []string{equipmenttype.PropertyTypesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipmenttype.EquipmentTable,
			Columns: []string{equipmenttype.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmenttype.CategoryTable,
			Columns: []string{equipmenttype.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.ServiceEndpointDefinitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   equipmenttype.ServiceEndpointDefinitionsTable,
			Columns: []string{equipmenttype.ServiceEndpointDefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: serviceendpointdefinition.FieldID,
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

// EquipmentTypeCreateBulk is the builder for creating a bulk of EquipmentType entities.
type EquipmentTypeCreateBulk struct {
	config
	builders []*EquipmentTypeCreate
}

// Save creates the EquipmentType entities in the database.
func (etcb *EquipmentTypeCreateBulk) Save(ctx context.Context) ([]*EquipmentType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EquipmentType, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (etcb *EquipmentTypeCreateBulk) SaveX(ctx context.Context) []*EquipmentType {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
