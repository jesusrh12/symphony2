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
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/recommendations"
	"github.com/facebookincubator/symphony/pkg/ent/recommendationssources"
)

// RecommendationsSourcesQuery is the builder for querying RecommendationsSources entities.
type RecommendationsSourcesQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.RecommendationsSources
	// eager-loading edges.
	withRecommendations *RecommendationsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rsq *RecommendationsSourcesQuery) Where(ps ...predicate.RecommendationsSources) *RecommendationsSourcesQuery {
	rsq.predicates = append(rsq.predicates, ps...)
	return rsq
}

// Limit adds a limit step to the query.
func (rsq *RecommendationsSourcesQuery) Limit(limit int) *RecommendationsSourcesQuery {
	rsq.limit = &limit
	return rsq
}

// Offset adds an offset step to the query.
func (rsq *RecommendationsSourcesQuery) Offset(offset int) *RecommendationsSourcesQuery {
	rsq.offset = &offset
	return rsq
}

// Order adds an order step to the query.
func (rsq *RecommendationsSourcesQuery) Order(o ...OrderFunc) *RecommendationsSourcesQuery {
	rsq.order = append(rsq.order, o...)
	return rsq
}

// QueryRecommendations chains the current query on the recommendations edge.
func (rsq *RecommendationsSourcesQuery) QueryRecommendations() *RecommendationsQuery {
	query := &RecommendationsQuery{config: rsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rsq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recommendationssources.Table, recommendationssources.FieldID, selector),
			sqlgraph.To(recommendations.Table, recommendations.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, recommendationssources.RecommendationsTable, recommendationssources.RecommendationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RecommendationsSources entity in the query. Returns *NotFoundError when no recommendationssources was found.
func (rsq *RecommendationsSourcesQuery) First(ctx context.Context) (*RecommendationsSources, error) {
	nodes, err := rsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recommendationssources.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) FirstX(ctx context.Context) *RecommendationsSources {
	node, err := rsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecommendationsSources id in the query. Returns *NotFoundError when no id was found.
func (rsq *RecommendationsSourcesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recommendationssources.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) FirstIDX(ctx context.Context) int {
	id, err := rsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only RecommendationsSources entity in the query, returns an error if not exactly one entity was returned.
func (rsq *RecommendationsSourcesQuery) Only(ctx context.Context) (*RecommendationsSources, error) {
	nodes, err := rsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recommendationssources.Label}
	default:
		return nil, &NotSingularError{recommendationssources.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) OnlyX(ctx context.Context) *RecommendationsSources {
	node, err := rsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only RecommendationsSources id in the query, returns an error if not exactly one id was returned.
func (rsq *RecommendationsSourcesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = &NotSingularError{recommendationssources.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) OnlyIDX(ctx context.Context) int {
	id, err := rsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecommendationsSourcesSlice.
func (rsq *RecommendationsSourcesQuery) All(ctx context.Context) ([]*RecommendationsSources, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) AllX(ctx context.Context) []*RecommendationsSources {
	nodes, err := rsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecommendationsSources ids.
func (rsq *RecommendationsSourcesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rsq.Select(recommendationssources.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) IDsX(ctx context.Context) []int {
	ids, err := rsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rsq *RecommendationsSourcesQuery) Count(ctx context.Context) (int, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) CountX(ctx context.Context) int {
	count, err := rsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rsq *RecommendationsSourcesQuery) Exist(ctx context.Context) (bool, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rsq *RecommendationsSourcesQuery) ExistX(ctx context.Context) bool {
	exist, err := rsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rsq *RecommendationsSourcesQuery) Clone() *RecommendationsSourcesQuery {
	if rsq == nil {
		return nil
	}
	return &RecommendationsSourcesQuery{
		config:              rsq.config,
		limit:               rsq.limit,
		offset:              rsq.offset,
		order:               append([]OrderFunc{}, rsq.order...),
		unique:              append([]string{}, rsq.unique...),
		predicates:          append([]predicate.RecommendationsSources{}, rsq.predicates...),
		withRecommendations: rsq.withRecommendations.Clone(),
		// clone intermediate query.
		sql:  rsq.sql.Clone(),
		path: rsq.path,
	}
}

//  WithRecommendations tells the query-builder to eager-loads the nodes that are connected to
// the "recommendations" edge. The optional arguments used to configure the query builder of the edge.
func (rsq *RecommendationsSourcesQuery) WithRecommendations(opts ...func(*RecommendationsQuery)) *RecommendationsSourcesQuery {
	query := &RecommendationsQuery{config: rsq.config}
	for _, opt := range opts {
		opt(query)
	}
	rsq.withRecommendations = query
	return rsq
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
//	client.RecommendationsSources.Query().
//		GroupBy(recommendationssources.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rsq *RecommendationsSourcesQuery) GroupBy(field string, fields ...string) *RecommendationsSourcesGroupBy {
	group := &RecommendationsSourcesGroupBy{config: rsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(), nil
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
//	client.RecommendationsSources.Query().
//		Select(recommendationssources.FieldCreateTime).
//		Scan(ctx, &v)
//
func (rsq *RecommendationsSourcesQuery) Select(field string, fields ...string) *RecommendationsSourcesSelect {
	selector := &RecommendationsSourcesSelect{config: rsq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(), nil
	}
	return selector
}

func (rsq *RecommendationsSourcesQuery) prepareQuery(ctx context.Context) error {
	if rsq.path != nil {
		prev, err := rsq.path(ctx)
		if err != nil {
			return err
		}
		rsq.sql = prev
	}
	if err := recommendationssources.Policy.EvalQuery(ctx, rsq); err != nil {
		return err
	}
	return nil
}

func (rsq *RecommendationsSourcesQuery) sqlAll(ctx context.Context) ([]*RecommendationsSources, error) {
	var (
		nodes       = []*RecommendationsSources{}
		_spec       = rsq.querySpec()
		loadedTypes = [1]bool{
			rsq.withRecommendations != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &RecommendationsSources{config: rsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rsq.withRecommendations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*RecommendationsSources)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Recommendations = []*Recommendations{}
		}
		query.withFKs = true
		query.Where(predicate.Recommendations(func(s *sql.Selector) {
			s.Where(sql.InValues(recommendationssources.RecommendationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.recommendations_sources_recommendations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "recommendations_sources_recommendations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "recommendations_sources_recommendations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Recommendations = append(node.Edges.Recommendations, n)
		}
	}

	return nodes, nil
}

func (rsq *RecommendationsSourcesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rsq.querySpec()
	return sqlgraph.CountNodes(ctx, rsq.driver, _spec)
}

func (rsq *RecommendationsSourcesQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rsq *RecommendationsSourcesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recommendationssources.Table,
			Columns: recommendationssources.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recommendationssources.FieldID,
			},
		},
		From:   rsq.sql,
		Unique: true,
	}
	if ps := rsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, recommendationssources.ValidColumn)
			}
		}
	}
	return _spec
}

func (rsq *RecommendationsSourcesQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rsq.driver.Dialect())
	t1 := builder.Table(recommendationssources.Table)
	selector := builder.Select(t1.Columns(recommendationssources.Columns...)...).From(t1)
	if rsq.sql != nil {
		selector = rsq.sql
		selector.Select(selector.Columns(recommendationssources.Columns...)...)
	}
	for _, p := range rsq.predicates {
		p(selector)
	}
	for _, p := range rsq.order {
		p(selector, recommendationssources.ValidColumn)
	}
	if offset := rsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecommendationsSourcesGroupBy is the builder for group-by RecommendationsSources entities.
type RecommendationsSourcesGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rsgb *RecommendationsSourcesGroupBy) Aggregate(fns ...AggregateFunc) *RecommendationsSourcesGroupBy {
	rsgb.fns = append(rsgb.fns, fns...)
	return rsgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rsgb *RecommendationsSourcesGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rsgb.path(ctx)
	if err != nil {
		return err
	}
	rsgb.sql = query
	return rsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) StringsX(ctx context.Context) []string {
	v, err := rsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) StringX(ctx context.Context) string {
	v, err := rsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) IntsX(ctx context.Context) []int {
	v, err := rsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) IntX(ctx context.Context) int {
	v, err := rsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RecommendationsSourcesGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rsgb *RecommendationsSourcesGroupBy) BoolX(ctx context.Context) bool {
	v, err := rsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rsgb *RecommendationsSourcesGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rsgb.fields {
		if !recommendationssources.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rsgb *RecommendationsSourcesGroupBy) sqlQuery() *sql.Selector {
	selector := rsgb.sql
	columns := make([]string, 0, len(rsgb.fields)+len(rsgb.fns))
	columns = append(columns, rsgb.fields...)
	for _, fn := range rsgb.fns {
		columns = append(columns, fn(selector, recommendationssources.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(rsgb.fields...)
}

// RecommendationsSourcesSelect is the builder for select fields of RecommendationsSources entities.
type RecommendationsSourcesSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rss *RecommendationsSourcesSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rss.path(ctx)
	if err != nil {
		return err
	}
	rss.sql = query
	return rss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) StringsX(ctx context.Context) []string {
	v, err := rss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) StringX(ctx context.Context) string {
	v, err := rss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) IntsX(ctx context.Context) []int {
	v, err := rss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) IntX(ctx context.Context) int {
	v, err := rss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) Float64X(ctx context.Context) float64 {
	v, err := rss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RecommendationsSourcesSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) BoolsX(ctx context.Context) []bool {
	v, err := rss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rss *RecommendationsSourcesSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recommendationssources.Label}
	default:
		err = fmt.Errorf("ent: RecommendationsSourcesSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rss *RecommendationsSourcesSelect) BoolX(ctx context.Context) bool {
	v, err := rss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rss *RecommendationsSourcesSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rss.fields {
		if !recommendationssources.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := rss.sqlQuery().Query()
	if err := rss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rss *RecommendationsSourcesSelect) sqlQuery() sql.Querier {
	selector := rss.sql
	selector.Select(selector.Columns(rss.fields...)...)
	return selector
}
