// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/symphony/pkg/ent/kqi"
	"github.com/facebookincubator/symphony/pkg/ent/kqiperspective"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
)

// KqiPerspectiveQuery is the builder for querying KqiPerspective entities.
type KqiPerspectiveQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.KqiPerspective
	// eager-loading edges.
	withKqiPerspectiveFk *KqiQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (kpq *KqiPerspectiveQuery) Where(ps ...predicate.KqiPerspective) *KqiPerspectiveQuery {
	kpq.predicates = append(kpq.predicates, ps...)
	return kpq
}

// Limit adds a limit step to the query.
func (kpq *KqiPerspectiveQuery) Limit(limit int) *KqiPerspectiveQuery {
	kpq.limit = &limit
	return kpq
}

// Offset adds an offset step to the query.
func (kpq *KqiPerspectiveQuery) Offset(offset int) *KqiPerspectiveQuery {
	kpq.offset = &offset
	return kpq
}

// Order adds an order step to the query.
func (kpq *KqiPerspectiveQuery) Order(o ...OrderFunc) *KqiPerspectiveQuery {
	kpq.order = append(kpq.order, o...)
	return kpq
}

// QueryKqiPerspectiveFk chains the current query on the kqiPerspectiveFk edge.
func (kpq *KqiPerspectiveQuery) QueryKqiPerspectiveFk() *KqiQuery {
	query := &KqiQuery{config: kpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kqiperspective.Table, kqiperspective.FieldID, selector),
			sqlgraph.To(kqi.Table, kqi.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, kqiperspective.KqiPerspectiveFkTable, kqiperspective.KqiPerspectiveFkColumn),
		)
		fromU = sqlgraph.SetNeighbors(kpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KqiPerspective entity in the query. Returns *NotFoundError when no kqiperspective was found.
func (kpq *KqiPerspectiveQuery) First(ctx context.Context) (*KqiPerspective, error) {
	nodes, err := kpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kqiperspective.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) FirstX(ctx context.Context) *KqiPerspective {
	node, err := kpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KqiPerspective id in the query. Returns *NotFoundError when no id was found.
func (kpq *KqiPerspectiveQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kqiperspective.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) FirstIDX(ctx context.Context) int {
	id, err := kpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only KqiPerspective entity in the query, returns an error if not exactly one entity was returned.
func (kpq *KqiPerspectiveQuery) Only(ctx context.Context) (*KqiPerspective, error) {
	nodes, err := kpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kqiperspective.Label}
	default:
		return nil, &NotSingularError{kqiperspective.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) OnlyX(ctx context.Context) *KqiPerspective {
	node, err := kpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only KqiPerspective id in the query, returns an error if not exactly one id was returned.
func (kpq *KqiPerspectiveQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = &NotSingularError{kqiperspective.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) OnlyIDX(ctx context.Context) int {
	id, err := kpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KqiPerspectives.
func (kpq *KqiPerspectiveQuery) All(ctx context.Context) ([]*KqiPerspective, error) {
	if err := kpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) AllX(ctx context.Context) []*KqiPerspective {
	nodes, err := kpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KqiPerspective ids.
func (kpq *KqiPerspectiveQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := kpq.Select(kqiperspective.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) IDsX(ctx context.Context) []int {
	ids, err := kpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kpq *KqiPerspectiveQuery) Count(ctx context.Context) (int, error) {
	if err := kpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) CountX(ctx context.Context) int {
	count, err := kpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kpq *KqiPerspectiveQuery) Exist(ctx context.Context) (bool, error) {
	if err := kpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kpq *KqiPerspectiveQuery) ExistX(ctx context.Context) bool {
	exist, err := kpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kpq *KqiPerspectiveQuery) Clone() *KqiPerspectiveQuery {
	if kpq == nil {
		return nil
	}
	return &KqiPerspectiveQuery{
		config:               kpq.config,
		limit:                kpq.limit,
		offset:               kpq.offset,
		order:                append([]OrderFunc{}, kpq.order...),
		unique:               append([]string{}, kpq.unique...),
		predicates:           append([]predicate.KqiPerspective{}, kpq.predicates...),
		withKqiPerspectiveFk: kpq.withKqiPerspectiveFk.Clone(),
		// clone intermediate query.
		sql:  kpq.sql.Clone(),
		path: kpq.path,
	}
}

//  WithKqiPerspectiveFk tells the query-builder to eager-loads the nodes that are connected to
// the "kqiPerspectiveFk" edge. The optional arguments used to configure the query builder of the edge.
func (kpq *KqiPerspectiveQuery) WithKqiPerspectiveFk(opts ...func(*KqiQuery)) *KqiPerspectiveQuery {
	query := &KqiQuery{config: kpq.config}
	for _, opt := range opts {
		opt(query)
	}
	kpq.withKqiPerspectiveFk = query
	return kpq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.KqiPerspective.Query().
//		GroupBy(kqiperspective.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (kpq *KqiPerspectiveQuery) GroupBy(field string, fields ...string) *KqiPerspectiveGroupBy {
	group := &KqiPerspectiveGroupBy{config: kpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kpq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.KqiPerspective.Query().
//		Select(kqiperspective.FieldCreateTime).
//		Scan(ctx, &v)
//
func (kpq *KqiPerspectiveQuery) Select(field string, fields ...string) *KqiPerspectiveSelect {
	selector := &KqiPerspectiveSelect{config: kpq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kpq.sqlQuery(), nil
	}
	return selector
}

func (kpq *KqiPerspectiveQuery) prepareQuery(ctx context.Context) error {
	if kpq.path != nil {
		prev, err := kpq.path(ctx)
		if err != nil {
			return err
		}
		kpq.sql = prev
	}
	if err := kqiperspective.Policy.EvalQuery(ctx, kpq); err != nil {
		return err
	}
	return nil
}

func (kpq *KqiPerspectiveQuery) sqlAll(ctx context.Context) ([]*KqiPerspective, error) {
	var (
		nodes       = []*KqiPerspective{}
		_spec       = kpq.querySpec()
		loadedTypes = [1]bool{
			kpq.withKqiPerspectiveFk != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &KqiPerspective{config: kpq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, kpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := kpq.withKqiPerspectiveFk; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*KqiPerspective)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.KqiPerspectiveFk = []*Kqi{}
		}
		query.withFKs = true
		query.Where(predicate.Kqi(func(s *sql.Selector) {
			s.Where(sql.InValues(kqiperspective.KqiPerspectiveFkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.kqi_perspective_kqi_perspective_fk
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "kqi_perspective_kqi_perspective_fk" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "kqi_perspective_kqi_perspective_fk" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.KqiPerspectiveFk = append(node.Edges.KqiPerspectiveFk, n)
		}
	}

	return nodes, nil
}

func (kpq *KqiPerspectiveQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kpq.querySpec()
	return sqlgraph.CountNodes(ctx, kpq.driver, _spec)
}

func (kpq *KqiPerspectiveQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (kpq *KqiPerspectiveQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kqiperspective.Table,
			Columns: kqiperspective.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: kqiperspective.FieldID,
			},
		},
		From:   kpq.sql,
		Unique: true,
	}
	if ps := kpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, kqiperspective.ValidColumn)
			}
		}
	}
	return _spec
}

func (kpq *KqiPerspectiveQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(kpq.driver.Dialect())
	t1 := builder.Table(kqiperspective.Table)
	selector := builder.Select(t1.Columns(kqiperspective.Columns...)...).From(t1)
	if kpq.sql != nil {
		selector = kpq.sql
		selector.Select(selector.Columns(kqiperspective.Columns...)...)
	}
	for _, p := range kpq.predicates {
		p(selector)
	}
	for _, p := range kpq.order {
		p(selector, kqiperspective.ValidColumn)
	}
	if offset := kpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KqiPerspectiveGroupBy is the builder for group-by KqiPerspective entities.
type KqiPerspectiveGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kpgb *KqiPerspectiveGroupBy) Aggregate(fns ...AggregateFunc) *KqiPerspectiveGroupBy {
	kpgb.fns = append(kpgb.fns, fns...)
	return kpgb
}

// Scan applies the group-by query and scan the result into the given value.
func (kpgb *KqiPerspectiveGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kpgb.path(ctx)
	if err != nil {
		return err
	}
	kpgb.sql = query
	return kpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := kpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(kpgb.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := kpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) StringsX(ctx context.Context) []string {
	v, err := kpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) StringX(ctx context.Context) string {
	v, err := kpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(kpgb.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := kpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) IntsX(ctx context.Context) []int {
	v, err := kpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) IntX(ctx context.Context) int {
	v, err := kpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(kpgb.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := kpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := kpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) Float64X(ctx context.Context) float64 {
	v, err := kpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(kpgb.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := kpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := kpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (kpgb *KqiPerspectiveGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kpgb *KqiPerspectiveGroupBy) BoolX(ctx context.Context) bool {
	v, err := kpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kpgb *KqiPerspectiveGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kpgb.fields {
		if !kqiperspective.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kpgb *KqiPerspectiveGroupBy) sqlQuery() *sql.Selector {
	selector := kpgb.sql
	columns := make([]string, 0, len(kpgb.fields)+len(kpgb.fns))
	columns = append(columns, kpgb.fields...)
	for _, fn := range kpgb.fns {
		columns = append(columns, fn(selector, kqiperspective.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(kpgb.fields...)
}

// KqiPerspectiveSelect is the builder for select fields of KqiPerspective entities.
type KqiPerspectiveSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (kps *KqiPerspectiveSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := kps.path(ctx)
	if err != nil {
		return err
	}
	kps.sql = query
	return kps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kps.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) StringsX(ctx context.Context) []string {
	v, err := kps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) StringX(ctx context.Context) string {
	v, err := kps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kps.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) IntsX(ctx context.Context) []int {
	v, err := kps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) IntX(ctx context.Context) int {
	v, err := kps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kps.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) Float64X(ctx context.Context) float64 {
	v, err := kps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kps.fields) > 1 {
		return nil, errors.New("ent: KqiPerspectiveSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) BoolsX(ctx context.Context) []bool {
	v, err := kps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (kps *KqiPerspectiveSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{kqiperspective.Label}
	default:
		err = fmt.Errorf("ent: KqiPerspectiveSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kps *KqiPerspectiveSelect) BoolX(ctx context.Context) bool {
	v, err := kps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kps *KqiPerspectiveSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kps.fields {
		if !kqiperspective.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := kps.sqlQuery().Query()
	if err := kps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kps *KqiPerspectiveSelect) sqlQuery() sql.Querier {
	selector := kps.sql
	selector.Select(selector.Columns(kps.fields...)...)
	return selector
}
