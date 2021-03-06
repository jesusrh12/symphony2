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
	"github.com/facebookincubator/symphony/pkg/ent/recommendationscategory"
)

// RecommendationsCategory is the model entity for the RecommendationsCategory schema.
type RecommendationsCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecommendationsCategoryQuery when eager-loading is set.
	Edges RecommendationsCategoryEdges `json:"edges"`
}

// RecommendationsCategoryEdges holds the relations/edges for other nodes in the graph.
type RecommendationsCategoryEdges struct {
	// Recommendations holds the value of the recommendations edge.
	Recommendations []*Recommendations
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecommendationsOrErr returns the Recommendations value or an error if the edge
// was not loaded in eager-loading.
func (e RecommendationsCategoryEdges) RecommendationsOrErr() ([]*Recommendations, error) {
	if e.loadedTypes[0] {
		return e.Recommendations, nil
	}
	return nil, &NotLoadedError{edge: "recommendations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RecommendationsCategory) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RecommendationsCategory fields.
func (rc *RecommendationsCategory) assignValues(values ...interface{}) error {
	if m, n := len(values), len(recommendationscategory.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	rc.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		rc.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		rc.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		rc.Name = value.String
	}
	return nil
}

// QueryRecommendations queries the recommendations edge of the RecommendationsCategory.
func (rc *RecommendationsCategory) QueryRecommendations() *RecommendationsQuery {
	return (&RecommendationsCategoryClient{config: rc.config}).QueryRecommendations(rc)
}

// Update returns a builder for updating this RecommendationsCategory.
// Note that, you need to call RecommendationsCategory.Unwrap() before calling this method, if this RecommendationsCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (rc *RecommendationsCategory) Update() *RecommendationsCategoryUpdateOne {
	return (&RecommendationsCategoryClient{config: rc.config}).UpdateOne(rc)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (rc *RecommendationsCategory) Unwrap() *RecommendationsCategory {
	tx, ok := rc.config.driver.(*txDriver)
	if !ok {
		panic("ent: RecommendationsCategory is not a transactional entity")
	}
	rc.config.driver = tx.drv
	return rc
}

// String implements the fmt.Stringer.
func (rc *RecommendationsCategory) String() string {
	var builder strings.Builder
	builder.WriteString("RecommendationsCategory(")
	builder.WriteString(fmt.Sprintf("id=%v", rc.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(rc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(rc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(rc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// RecommendationsCategories is a parsable slice of RecommendationsCategory.
type RecommendationsCategories []*RecommendationsCategory

func (rc RecommendationsCategories) config(cfg config) {
	for _i := range rc {
		rc[_i].config = cfg
	}
}
