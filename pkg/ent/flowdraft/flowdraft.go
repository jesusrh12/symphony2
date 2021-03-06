// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package flowdraft

import (
	"time"

	"github.com/facebook/ent"
)

const (
	// Label holds the string label denoting the flowdraft type in the database.
	Label = "flow_draft"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldEndParamDefinitions holds the string denoting the end_param_definitions field in the database.
	FieldEndParamDefinitions = "end_param_definitions"
	// FieldSameAsFlow holds the string denoting the sameasflow field in the database.
	FieldSameAsFlow = "same_as_flow"

	// EdgeBlocks holds the string denoting the blocks edge name in mutations.
	EdgeBlocks = "blocks"
	// EdgeFlow holds the string denoting the flow edge name in mutations.
	EdgeFlow = "flow"

	// Table holds the table name of the flowdraft in the database.
	Table = "flow_drafts"
	// BlocksTable is the table the holds the blocks relation/edge.
	BlocksTable = "blocks"
	// BlocksInverseTable is the table name for the Block entity.
	// It exists in this package in order to avoid circular dependency with the "block" package.
	BlocksInverseTable = "blocks"
	// BlocksColumn is the table column denoting the blocks relation/edge.
	BlocksColumn = "flow_draft_blocks"
	// FlowTable is the table the holds the flow relation/edge.
	FlowTable = "flow_drafts"
	// FlowInverseTable is the table name for the Flow entity.
	// It exists in this package in order to avoid circular dependency with the "flow" package.
	FlowInverseTable = "flows"
	// FlowColumn is the table column denoting the flow relation/edge.
	FlowColumn = "flow_draft"
)

// Columns holds all SQL columns for flowdraft fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldDescription,
	FieldEndParamDefinitions,
	FieldSameAsFlow,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the FlowDraft type.
var ForeignKeys = []string{
	"flow_draft",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	Hooks  [3]ent.Hook
	Policy ent.Policy
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultSameAsFlow holds the default value on creation for the sameAsFlow field.
	DefaultSameAsFlow bool
)
