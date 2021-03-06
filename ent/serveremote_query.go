// Code generated by entc, DO NOT EDIT.

package ent

import (
	"betterdsc/ent/predicate"
	"betterdsc/ent/serveremote"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServerEmoteQuery is the builder for querying ServerEmote entities.
type ServerEmoteQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ServerEmote
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServerEmoteQuery builder.
func (seq *ServerEmoteQuery) Where(ps ...predicate.ServerEmote) *ServerEmoteQuery {
	seq.predicates = append(seq.predicates, ps...)
	return seq
}

// Limit adds a limit step to the query.
func (seq *ServerEmoteQuery) Limit(limit int) *ServerEmoteQuery {
	seq.limit = &limit
	return seq
}

// Offset adds an offset step to the query.
func (seq *ServerEmoteQuery) Offset(offset int) *ServerEmoteQuery {
	seq.offset = &offset
	return seq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (seq *ServerEmoteQuery) Unique(unique bool) *ServerEmoteQuery {
	seq.unique = &unique
	return seq
}

// Order adds an order step to the query.
func (seq *ServerEmoteQuery) Order(o ...OrderFunc) *ServerEmoteQuery {
	seq.order = append(seq.order, o...)
	return seq
}

// First returns the first ServerEmote entity from the query.
// Returns a *NotFoundError when no ServerEmote was found.
func (seq *ServerEmoteQuery) First(ctx context.Context) (*ServerEmote, error) {
	nodes, err := seq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{serveremote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (seq *ServerEmoteQuery) FirstX(ctx context.Context) *ServerEmote {
	node, err := seq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServerEmote ID from the query.
// Returns a *NotFoundError when no ServerEmote ID was found.
func (seq *ServerEmoteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{serveremote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (seq *ServerEmoteQuery) FirstIDX(ctx context.Context) int {
	id, err := seq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServerEmote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServerEmote entity is found.
// Returns a *NotFoundError when no ServerEmote entities are found.
func (seq *ServerEmoteQuery) Only(ctx context.Context) (*ServerEmote, error) {
	nodes, err := seq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{serveremote.Label}
	default:
		return nil, &NotSingularError{serveremote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (seq *ServerEmoteQuery) OnlyX(ctx context.Context) *ServerEmote {
	node, err := seq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServerEmote ID in the query.
// Returns a *NotSingularError when more than one ServerEmote ID is found.
// Returns a *NotFoundError when no entities are found.
func (seq *ServerEmoteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = &NotSingularError{serveremote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (seq *ServerEmoteQuery) OnlyIDX(ctx context.Context) int {
	id, err := seq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServerEmotes.
func (seq *ServerEmoteQuery) All(ctx context.Context) ([]*ServerEmote, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return seq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (seq *ServerEmoteQuery) AllX(ctx context.Context) []*ServerEmote {
	nodes, err := seq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServerEmote IDs.
func (seq *ServerEmoteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := seq.Select(serveremote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (seq *ServerEmoteQuery) IDsX(ctx context.Context) []int {
	ids, err := seq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (seq *ServerEmoteQuery) Count(ctx context.Context) (int, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return seq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (seq *ServerEmoteQuery) CountX(ctx context.Context) int {
	count, err := seq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (seq *ServerEmoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := seq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return seq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (seq *ServerEmoteQuery) ExistX(ctx context.Context) bool {
	exist, err := seq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServerEmoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (seq *ServerEmoteQuery) Clone() *ServerEmoteQuery {
	if seq == nil {
		return nil
	}
	return &ServerEmoteQuery{
		config:     seq.config,
		limit:      seq.limit,
		offset:     seq.offset,
		order:      append([]OrderFunc{}, seq.order...),
		predicates: append([]predicate.ServerEmote{}, seq.predicates...),
		// clone intermediate query.
		sql:    seq.sql.Clone(),
		path:   seq.path,
		unique: seq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ServerID string `json:"server_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServerEmote.Query().
//		GroupBy(serveremote.FieldServerID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (seq *ServerEmoteQuery) GroupBy(field string, fields ...string) *ServerEmoteGroupBy {
	group := &ServerEmoteGroupBy{config: seq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return seq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ServerID string `json:"server_id,omitempty"`
//	}
//
//	client.ServerEmote.Query().
//		Select(serveremote.FieldServerID).
//		Scan(ctx, &v)
//
func (seq *ServerEmoteQuery) Select(fields ...string) *ServerEmoteSelect {
	seq.fields = append(seq.fields, fields...)
	return &ServerEmoteSelect{ServerEmoteQuery: seq}
}

func (seq *ServerEmoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range seq.fields {
		if !serveremote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if seq.path != nil {
		prev, err := seq.path(ctx)
		if err != nil {
			return err
		}
		seq.sql = prev
	}
	return nil
}

func (seq *ServerEmoteQuery) sqlAll(ctx context.Context) ([]*ServerEmote, error) {
	var (
		nodes = []*ServerEmote{}
		_spec = seq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ServerEmote{config: seq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, seq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (seq *ServerEmoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := seq.querySpec()
	_spec.Node.Columns = seq.fields
	if len(seq.fields) > 0 {
		_spec.Unique = seq.unique != nil && *seq.unique
	}
	return sqlgraph.CountNodes(ctx, seq.driver, _spec)
}

func (seq *ServerEmoteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := seq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (seq *ServerEmoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serveremote.Table,
			Columns: serveremote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serveremote.FieldID,
			},
		},
		From:   seq.sql,
		Unique: true,
	}
	if unique := seq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := seq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serveremote.FieldID)
		for i := range fields {
			if fields[i] != serveremote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := seq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := seq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := seq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := seq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (seq *ServerEmoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(seq.driver.Dialect())
	t1 := builder.Table(serveremote.Table)
	columns := seq.fields
	if len(columns) == 0 {
		columns = serveremote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if seq.sql != nil {
		selector = seq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if seq.unique != nil && *seq.unique {
		selector.Distinct()
	}
	for _, p := range seq.predicates {
		p(selector)
	}
	for _, p := range seq.order {
		p(selector)
	}
	if offset := seq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := seq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServerEmoteGroupBy is the group-by builder for ServerEmote entities.
type ServerEmoteGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (segb *ServerEmoteGroupBy) Aggregate(fns ...AggregateFunc) *ServerEmoteGroupBy {
	segb.fns = append(segb.fns, fns...)
	return segb
}

// Scan applies the group-by query and scans the result into the given value.
func (segb *ServerEmoteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := segb.path(ctx)
	if err != nil {
		return err
	}
	segb.sql = query
	return segb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := segb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(segb.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := segb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) StringsX(ctx context.Context) []string {
	v, err := segb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = segb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) StringX(ctx context.Context) string {
	v, err := segb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(segb.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := segb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) IntsX(ctx context.Context) []int {
	v, err := segb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = segb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) IntX(ctx context.Context) int {
	v, err := segb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(segb.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := segb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := segb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = segb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) Float64X(ctx context.Context) float64 {
	v, err := segb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(segb.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := segb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := segb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (segb *ServerEmoteGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = segb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (segb *ServerEmoteGroupBy) BoolX(ctx context.Context) bool {
	v, err := segb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (segb *ServerEmoteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range segb.fields {
		if !serveremote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := segb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := segb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (segb *ServerEmoteGroupBy) sqlQuery() *sql.Selector {
	selector := segb.sql.Select()
	aggregation := make([]string, 0, len(segb.fns))
	for _, fn := range segb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(segb.fields)+len(segb.fns))
		for _, f := range segb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(segb.fields...)...)
}

// ServerEmoteSelect is the builder for selecting fields of ServerEmote entities.
type ServerEmoteSelect struct {
	*ServerEmoteQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ses *ServerEmoteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ses.prepareQuery(ctx); err != nil {
		return err
	}
	ses.sql = ses.ServerEmoteQuery.sqlQuery(ctx)
	return ses.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ses *ServerEmoteSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ses.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ses.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ses.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ses *ServerEmoteSelect) StringsX(ctx context.Context) []string {
	v, err := ses.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ses.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ses *ServerEmoteSelect) StringX(ctx context.Context) string {
	v, err := ses.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ses.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ses.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ses *ServerEmoteSelect) IntsX(ctx context.Context) []int {
	v, err := ses.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ses.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ses *ServerEmoteSelect) IntX(ctx context.Context) int {
	v, err := ses.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ses.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ses.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ses *ServerEmoteSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ses.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ses.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ses *ServerEmoteSelect) Float64X(ctx context.Context) float64 {
	v, err := ses.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ses.fields) > 1 {
		return nil, errors.New("ent: ServerEmoteSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ses.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ses *ServerEmoteSelect) BoolsX(ctx context.Context) []bool {
	v, err := ses.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ses *ServerEmoteSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ses.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serveremote.Label}
	default:
		err = fmt.Errorf("ent: ServerEmoteSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ses *ServerEmoteSelect) BoolX(ctx context.Context) bool {
	v, err := ses.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ses *ServerEmoteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ses.sql.Query()
	if err := ses.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
