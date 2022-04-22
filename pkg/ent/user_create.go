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
	"github.com/facebookincubator/symphony/pkg/ent/appointment"
	"github.com/facebookincubator/symphony/pkg/ent/feature"
	"github.com/facebookincubator/symphony/pkg/ent/file"
	"github.com/facebookincubator/symphony/pkg/ent/organization"
	"github.com/facebookincubator/symphony/pkg/ent/project"
	"github.com/facebookincubator/symphony/pkg/ent/recommendations"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/ent/usersgroup"
	"github.com/facebookincubator/symphony/pkg/ent/workorder"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the update_time field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetAuthID sets the auth_id field.
func (uc *UserCreate) SetAuthID(s string) *UserCreate {
	uc.mutation.SetAuthID(s)
	return uc
}

// SetFirstName sets the first_name field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetNillableFirstName sets the first_name field if the given value is not nil.
func (uc *UserCreate) SetNillableFirstName(s *string) *UserCreate {
	if s != nil {
		uc.SetFirstName(*s)
	}
	return uc
}

// SetLastName sets the last_name field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetNillableLastName sets the last_name field if the given value is not nil.
func (uc *UserCreate) SetNillableLastName(s *string) *UserCreate {
	if s != nil {
		uc.SetLastName(*s)
	}
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the email field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetStatus sets the status field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(u *user.Status) *UserCreate {
	if u != nil {
		uc.SetStatus(*u)
	}
	return uc
}

// SetRole sets the role field.
func (uc *UserCreate) SetRole(u user.Role) *UserCreate {
	uc.mutation.SetRole(u)
	return uc
}

// SetNillableRole sets the role field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(u *user.Role) *UserCreate {
	if u != nil {
		uc.SetRole(*u)
	}
	return uc
}

// SetDistanceUnit sets the distance_unit field.
func (uc *UserCreate) SetDistanceUnit(uu user.DistanceUnit) *UserCreate {
	uc.mutation.SetDistanceUnit(uu)
	return uc
}

// SetNillableDistanceUnit sets the distance_unit field if the given value is not nil.
func (uc *UserCreate) SetNillableDistanceUnit(uu *user.DistanceUnit) *UserCreate {
	if uu != nil {
		uc.SetDistanceUnit(*uu)
	}
	return uc
}

// SetProfilePhotoID sets the profile_photo edge to File by id.
func (uc *UserCreate) SetProfilePhotoID(id int) *UserCreate {
	uc.mutation.SetProfilePhotoID(id)
	return uc
}

// SetNillableProfilePhotoID sets the profile_photo edge to File by id if the given value is not nil.
func (uc *UserCreate) SetNillableProfilePhotoID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetProfilePhotoID(*id)
	}
	return uc
}

// SetProfilePhoto sets the profile_photo edge to File.
func (uc *UserCreate) SetProfilePhoto(f *File) *UserCreate {
	return uc.SetProfilePhotoID(f.ID)
}

// AddUserCreateIDs adds the User_create edge to Recommendations by ids.
func (uc *UserCreate) AddUserCreateIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserCreateIDs(ids...)
	return uc
}

// AddUserCreate adds the User_create edges to Recommendations.
func (uc *UserCreate) AddUserCreate(r ...*Recommendations) *UserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddUserCreateIDs(ids...)
}

// AddUserApprovedIDs adds the User_approved edge to Recommendations by ids.
func (uc *UserCreate) AddUserApprovedIDs(ids ...int) *UserCreate {
	uc.mutation.AddUserApprovedIDs(ids...)
	return uc
}

// AddUserApproved adds the User_approved edges to Recommendations.
func (uc *UserCreate) AddUserApproved(r ...*Recommendations) *UserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddUserApprovedIDs(ids...)
}

// AddGroupIDs adds the groups edge to UsersGroup by ids.
func (uc *UserCreate) AddGroupIDs(ids ...int) *UserCreate {
	uc.mutation.AddGroupIDs(ids...)
	return uc
}

// AddGroups adds the groups edges to UsersGroup.
func (uc *UserCreate) AddGroups(u ...*UsersGroup) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// SetOrganizationID sets the organization edge to Organization by id.
func (uc *UserCreate) SetOrganizationID(id int) *UserCreate {
	uc.mutation.SetOrganizationID(id)
	return uc
}

// SetNillableOrganizationID sets the organization edge to Organization by id if the given value is not nil.
func (uc *UserCreate) SetNillableOrganizationID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetOrganizationID(*id)
	}
	return uc
}

// SetOrganization sets the organization edge to Organization.
func (uc *UserCreate) SetOrganization(o *Organization) *UserCreate {
	return uc.SetOrganizationID(o.ID)
}

// AddOwnedWorkOrderIDs adds the owned_work_orders edge to WorkOrder by ids.
func (uc *UserCreate) AddOwnedWorkOrderIDs(ids ...int) *UserCreate {
	uc.mutation.AddOwnedWorkOrderIDs(ids...)
	return uc
}

// AddOwnedWorkOrders adds the owned_work_orders edges to WorkOrder.
func (uc *UserCreate) AddOwnedWorkOrders(w ...*WorkOrder) *UserCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddOwnedWorkOrderIDs(ids...)
}

// AddAssignedWorkOrderIDs adds the assigned_work_orders edge to WorkOrder by ids.
func (uc *UserCreate) AddAssignedWorkOrderIDs(ids ...int) *UserCreate {
	uc.mutation.AddAssignedWorkOrderIDs(ids...)
	return uc
}

// AddAssignedWorkOrders adds the assigned_work_orders edges to WorkOrder.
func (uc *UserCreate) AddAssignedWorkOrders(w ...*WorkOrder) *UserCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddAssignedWorkOrderIDs(ids...)
}

// AddCreatedProjectIDs adds the created_projects edge to Project by ids.
func (uc *UserCreate) AddCreatedProjectIDs(ids ...int) *UserCreate {
	uc.mutation.AddCreatedProjectIDs(ids...)
	return uc
}

// AddCreatedProjects adds the created_projects edges to Project.
func (uc *UserCreate) AddCreatedProjects(p ...*Project) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddCreatedProjectIDs(ids...)
}

// AddFeatureIDs adds the features edge to Feature by ids.
func (uc *UserCreate) AddFeatureIDs(ids ...int) *UserCreate {
	uc.mutation.AddFeatureIDs(ids...)
	return uc
}

// AddFeatures adds the features edges to Feature.
func (uc *UserCreate) AddFeatures(f ...*Feature) *UserCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uc.AddFeatureIDs(ids...)
}

// AddAppointmentIDs adds the appointment edge to Appointment by ids.
func (uc *UserCreate) AddAppointmentIDs(ids ...int) *UserCreate {
	uc.mutation.AddAppointmentIDs(ids...)
	return uc
}

// AddAppointment adds the appointment edges to Appointment.
func (uc *UserCreate) AddAppointment(a ...*Appointment) *UserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAppointmentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.Role(); !ok {
		v := user.DefaultRole
		uc.mutation.SetRole(v)
	}
	if _, ok := uc.mutation.DistanceUnit(); !ok {
		v := user.DefaultDistanceUnit
		uc.mutation.SetDistanceUnit(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := uc.mutation.AuthID(); !ok {
		return &ValidationError{Name: "auth_id", err: errors.New("ent: missing required field \"auth_id\"")}
	}
	if v, ok := uc.mutation.AuthID(); ok {
		if err := user.AuthIDValidator(v); err != nil {
			return &ValidationError{Name: "auth_id", err: fmt.Errorf("ent: validator failed for field \"auth_id\": %w", err)}
		}
	}
	if v, ok := uc.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf("ent: validator failed for field \"first_name\": %w", err)}
		}
	}
	if v, ok := uc.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf("ent: validator failed for field \"last_name\": %w", err)}
		}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New("ent: missing required field \"role\"")}
	}
	if v, ok := uc.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	if _, ok := uc.mutation.DistanceUnit(); !ok {
		return &ValidationError{Name: "distance_unit", err: errors.New("ent: missing required field \"distance_unit\"")}
	}
	if v, ok := uc.mutation.DistanceUnit(); ok {
		if err := user.DistanceUnitValidator(v); err != nil {
			return &ValidationError{Name: "distance_unit", err: fmt.Errorf("ent: validator failed for field \"distance_unit\": %w", err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.AuthID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAuthID,
		})
		_node.AuthID = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := uc.mutation.DistanceUnit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldDistanceUnit,
		})
		_node.DistanceUnit = value
	}
	if nodes := uc.mutation.ProfilePhotoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ProfilePhotoTable,
			Columns: []string{user.ProfilePhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserCreateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCreateTable,
			Columns: []string{user.UserCreateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recommendations.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserApprovedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserApprovedTable,
			Columns: []string{user.UserApprovedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recommendations.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: usersgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.OrganizationTable,
			Columns: []string{user.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.OwnedWorkOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OwnedWorkOrdersTable,
			Columns: []string{user.OwnedWorkOrdersColumn},
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
	if nodes := uc.mutation.AssignedWorkOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.AssignedWorkOrdersTable,
			Columns: []string{user.AssignedWorkOrdersColumn},
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
	if nodes := uc.mutation.CreatedProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CreatedProjectsTable,
			Columns: []string{user.CreatedProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FeaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FeaturesTable,
			Columns: user.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AppointmentTable,
			Columns: []string{user.AppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointment.FieldID,
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

// UserCreateBulk is the builder for creating a bulk of User entities.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
