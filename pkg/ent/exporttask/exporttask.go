// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package exporttask

import (
	"fmt"
	"io"
	"strconv"

	"github.com/facebook/ent"
)

const (
	// Label holds the string label denoting the exporttask type in the database.
	Label = "export_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldProgress holds the string denoting the progress field in the database.
	FieldProgress = "progress"
	// FieldFilters holds the string denoting the filters field in the database.
	FieldFilters = "filters"
	// FieldStoreKey holds the string denoting the store_key field in the database.
	FieldStoreKey = "store_key"
	// FieldWoIDToExport holds the string denoting the wo_id_to_export field in the database.
	FieldWoIDToExport = "wo_id_to_export"

	// Table holds the table name of the exporttask in the database.
	Table = "export_tasks"
)

// Columns holds all SQL columns for exporttask fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldStatus,
	FieldProgress,
	FieldFilters,
	FieldStoreKey,
	FieldWoIDToExport,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/facebookincubator/symphony/pkg/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultProgress holds the default value on creation for the progress field.
	DefaultProgress float64
	// ProgressValidator is a validator for the "progress" field. It is called by the builders before save.
	ProgressValidator func(float64) error
	// DefaultFilters holds the default value on creation for the filters field.
	DefaultFilters string
)

// Type defines the type for the type enum field.
type Type string

// Type values.
const (
	TypeEquipment       Type = "EQUIPMENT"
	TypeLocation        Type = "LOCATION"
	TypePort            Type = "PORT"
	TypeLink            Type = "LINK"
	TypeService         Type = "SERVICE"
	TypeWorkOrder       Type = "WORK_ORDER"
	TypeSingleWorkOrder Type = "SINGLE_WORK_ORDER"
	TypeProject         Type = "PROJECT"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeEquipment, TypeLocation, TypePort, TypeLink, TypeService, TypeWorkOrder, TypeSingleWorkOrder, TypeProject:
		return nil
	default:
		return fmt.Errorf("exporttask: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusPending    Status = "PENDING"
	StatusInProgress Status = "IN_PROGRESS"
	StatusSucceeded  Status = "SUCCEEDED"
	StatusFailed     Status = "FAILED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusInProgress, StatusSucceeded, StatusFailed:
		return nil
	default:
		return fmt.Errorf("exporttask: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (_type Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_type.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_type *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_type = Type(str)
	if err := TypeValidator(*_type); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (s Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(s.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (s *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*s = Status(str)
	if err := StatusValidator(*s); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}