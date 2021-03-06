// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebookincubator/symphony/pkg/ent/organization"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges OrganizationEdges `json:"edges"`
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// UserFk holds the value of the user_fk edge.
	UserFk []*User
	// WorkOrderFk holds the value of the work_order_fk edge.
	WorkOrderFk []*WorkOrder
	// Policies holds the value of the policies edge.
	Policies []*PermissionsPolicy
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserFkOrErr returns the UserFk value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) UserFkOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.UserFk, nil
	}
	return nil, &NotLoadedError{edge: "user_fk"}
}

// WorkOrderFkOrErr returns the WorkOrderFk value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) WorkOrderFkOrErr() ([]*WorkOrder, error) {
	if e.loadedTypes[1] {
		return e.WorkOrderFk, nil
	}
	return nil, &NotLoadedError{edge: "work_order_fk"}
}

// PoliciesOrErr returns the Policies value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PoliciesOrErr() ([]*PermissionsPolicy, error) {
	if e.loadedTypes[2] {
		return e.Policies, nil
	}
	return nil, &NotLoadedError{edge: "policies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // name
		&sql.NullString{}, // description
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(values ...interface{}) error {
	if m, n := len(values), len(organization.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		o.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		o.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		o.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[3])
	} else if value.Valid {
		o.Description = value.String
	}
	return nil
}

// QueryUserFk queries the user_fk edge of the Organization.
func (o *Organization) QueryUserFk() *UserQuery {
	return (&OrganizationClient{config: o.config}).QueryUserFk(o)
}

// QueryWorkOrderFk queries the work_order_fk edge of the Organization.
func (o *Organization) QueryWorkOrderFk() *WorkOrderQuery {
	return (&OrganizationClient{config: o.config}).QueryWorkOrderFk(o)
}

// QueryPolicies queries the policies edge of the Organization.
func (o *Organization) QueryPolicies() *PermissionsPolicyQuery {
	return (&OrganizationClient{config: o.config}).QueryPolicies(o)
}

// Update returns a builder for updating this Organization.
// Note that, you need to call Organization.Unwrap() before calling this method, if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return (&OrganizationClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(o.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(o.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(o.Name)
	builder.WriteString(", description=")
	builder.WriteString(o.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization

func (o Organizations) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
