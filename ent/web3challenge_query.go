// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/predicate"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3challenge"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3user"
)

// Web3ChallengeQuery is the builder for querying Web3Challenge entities.
type Web3ChallengeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Web3Challenge
	// eager-loading edges.
	withWeb3User *Web3UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the Web3ChallengeQuery builder.
func (wq *Web3ChallengeQuery) Where(ps ...predicate.Web3Challenge) *Web3ChallengeQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *Web3ChallengeQuery) Limit(limit int) *Web3ChallengeQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *Web3ChallengeQuery) Offset(offset int) *Web3ChallengeQuery {
	wq.offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *Web3ChallengeQuery) Unique(unique bool) *Web3ChallengeQuery {
	wq.unique = &unique
	return wq
}

// Order adds an order step to the query.
func (wq *Web3ChallengeQuery) Order(o ...OrderFunc) *Web3ChallengeQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryWeb3User chains the current query on the "web3_user" edge.
func (wq *Web3ChallengeQuery) QueryWeb3User() *Web3UserQuery {
	query := &Web3UserQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(web3challenge.Table, web3challenge.FieldID, selector),
			sqlgraph.To(web3user.Table, web3user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, web3challenge.Web3UserTable, web3challenge.Web3UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Web3Challenge entity from the query.
// Returns a *NotFoundError when no Web3Challenge was found.
func (wq *Web3ChallengeQuery) First(ctx context.Context) (*Web3Challenge, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{web3challenge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *Web3ChallengeQuery) FirstX(ctx context.Context) *Web3Challenge {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Web3Challenge ID from the query.
// Returns a *NotFoundError when no Web3Challenge ID was found.
func (wq *Web3ChallengeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{web3challenge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *Web3ChallengeQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Web3Challenge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Web3Challenge entity is found.
// Returns a *NotFoundError when no Web3Challenge entities are found.
func (wq *Web3ChallengeQuery) Only(ctx context.Context) (*Web3Challenge, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{web3challenge.Label}
	default:
		return nil, &NotSingularError{web3challenge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *Web3ChallengeQuery) OnlyX(ctx context.Context) *Web3Challenge {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Web3Challenge ID in the query.
// Returns a *NotSingularError when more than one Web3Challenge ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *Web3ChallengeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = &NotSingularError{web3challenge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *Web3ChallengeQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Web3Challenges.
func (wq *Web3ChallengeQuery) All(ctx context.Context) ([]*Web3Challenge, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *Web3ChallengeQuery) AllX(ctx context.Context) []*Web3Challenge {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Web3Challenge IDs.
func (wq *Web3ChallengeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wq.Select(web3challenge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *Web3ChallengeQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *Web3ChallengeQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *Web3ChallengeQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *Web3ChallengeQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *Web3ChallengeQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the Web3ChallengeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *Web3ChallengeQuery) Clone() *Web3ChallengeQuery {
	if wq == nil {
		return nil
	}
	return &Web3ChallengeQuery{
		config:       wq.config,
		limit:        wq.limit,
		offset:       wq.offset,
		order:        append([]OrderFunc{}, wq.order...),
		predicates:   append([]predicate.Web3Challenge{}, wq.predicates...),
		withWeb3User: wq.withWeb3User.Clone(),
		// clone intermediate query.
		sql:    wq.sql.Clone(),
		path:   wq.path,
		unique: wq.unique,
	}
}

// WithWeb3User tells the query-builder to eager-load the nodes that are connected to
// the "web3_user" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *Web3ChallengeQuery) WithWeb3User(opts ...func(*Web3UserQuery)) *Web3ChallengeQuery {
	query := &Web3UserQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withWeb3User = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Web3Challenge.Query().
//		GroupBy(web3challenge.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *Web3ChallengeQuery) GroupBy(field string, fields ...string) *Web3ChallengeGroupBy {
	group := &Web3ChallengeGroupBy{config: wq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//	}
//
//	client.Web3Challenge.Query().
//		Select(web3challenge.FieldUUID).
//		Scan(ctx, &v)
//
func (wq *Web3ChallengeQuery) Select(fields ...string) *Web3ChallengeSelect {
	wq.fields = append(wq.fields, fields...)
	return &Web3ChallengeSelect{Web3ChallengeQuery: wq}
}

func (wq *Web3ChallengeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wq.fields {
		if !web3challenge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *Web3ChallengeQuery) sqlAll(ctx context.Context) ([]*Web3Challenge, error) {
	var (
		nodes       = []*Web3Challenge{}
		_spec       = wq.querySpec()
		loadedTypes = [1]bool{
			wq.withWeb3User != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Web3Challenge{config: wq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wq.withWeb3User; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Web3Challenge)
		for i := range nodes {
			fk := nodes[i].Web3UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(web3user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "web3_user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Web3User = n
			}
		}
	}

	return nodes, nil
}

func (wq *Web3ChallengeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	_spec.Node.Columns = wq.fields
	if len(wq.fields) > 0 {
		_spec.Unique = wq.unique != nil && *wq.unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *Web3ChallengeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wq *Web3ChallengeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   web3challenge.Table,
			Columns: web3challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: web3challenge.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if unique := wq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, web3challenge.FieldID)
		for i := range fields {
			if fields[i] != web3challenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *Web3ChallengeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(web3challenge.Table)
	columns := wq.fields
	if len(columns) == 0 {
		columns = web3challenge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.unique != nil && *wq.unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Web3ChallengeGroupBy is the group-by builder for Web3Challenge entities.
type Web3ChallengeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *Web3ChallengeGroupBy) Aggregate(fns ...AggregateFunc) *Web3ChallengeGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wgb *Web3ChallengeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) StringsX(ctx context.Context) []string {
	v, err := wgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) StringX(ctx context.Context) string {
	v, err := wgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) IntsX(ctx context.Context) []int {
	v, err := wgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) IntX(ctx context.Context) int {
	v, err := wgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *Web3ChallengeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wgb *Web3ChallengeGroupBy) BoolX(ctx context.Context) bool {
	v, err := wgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wgb *Web3ChallengeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wgb.fields {
		if !web3challenge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *Web3ChallengeGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql.Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
		for _, f := range wgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wgb.fields...)...)
}

// Web3ChallengeSelect is the builder for selecting fields of Web3Challenge entities.
type Web3ChallengeSelect struct {
	*Web3ChallengeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ws *Web3ChallengeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	ws.sql = ws.Web3ChallengeQuery.sqlQuery(ctx)
	return ws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ws *Web3ChallengeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ws *Web3ChallengeSelect) StringsX(ctx context.Context) []string {
	v, err := ws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ws *Web3ChallengeSelect) StringX(ctx context.Context) string {
	v, err := ws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ws *Web3ChallengeSelect) IntsX(ctx context.Context) []int {
	v, err := ws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ws *Web3ChallengeSelect) IntX(ctx context.Context) int {
	v, err := ws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ws *Web3ChallengeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ws *Web3ChallengeSelect) Float64X(ctx context.Context) float64 {
	v, err := ws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: Web3ChallengeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ws *Web3ChallengeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ws *Web3ChallengeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{web3challenge.Label}
	default:
		err = fmt.Errorf("ent: Web3ChallengeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ws *Web3ChallengeSelect) BoolX(ctx context.Context) bool {
	v, err := ws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ws *Web3ChallengeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ws.sql.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
