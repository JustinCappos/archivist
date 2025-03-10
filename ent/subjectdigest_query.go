// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/predicate"
	"github.com/testifysec/archivist/ent/subject"
	"github.com/testifysec/archivist/ent/subjectdigest"
)

// SubjectDigestQuery is the builder for querying SubjectDigest entities.
type SubjectDigestQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.SubjectDigest
	withSubject *SubjectQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*SubjectDigest) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubjectDigestQuery builder.
func (sdq *SubjectDigestQuery) Where(ps ...predicate.SubjectDigest) *SubjectDigestQuery {
	sdq.predicates = append(sdq.predicates, ps...)
	return sdq
}

// Limit adds a limit step to the query.
func (sdq *SubjectDigestQuery) Limit(limit int) *SubjectDigestQuery {
	sdq.limit = &limit
	return sdq
}

// Offset adds an offset step to the query.
func (sdq *SubjectDigestQuery) Offset(offset int) *SubjectDigestQuery {
	sdq.offset = &offset
	return sdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdq *SubjectDigestQuery) Unique(unique bool) *SubjectDigestQuery {
	sdq.unique = &unique
	return sdq
}

// Order adds an order step to the query.
func (sdq *SubjectDigestQuery) Order(o ...OrderFunc) *SubjectDigestQuery {
	sdq.order = append(sdq.order, o...)
	return sdq
}

// QuerySubject chains the current query on the "subject" edge.
func (sdq *SubjectDigestQuery) QuerySubject() *SubjectQuery {
	query := &SubjectQuery{config: sdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subjectdigest.Table, subjectdigest.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subjectdigest.SubjectTable, subjectdigest.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(sdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SubjectDigest entity from the query.
// Returns a *NotFoundError when no SubjectDigest was found.
func (sdq *SubjectDigestQuery) First(ctx context.Context) (*SubjectDigest, error) {
	nodes, err := sdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subjectdigest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdq *SubjectDigestQuery) FirstX(ctx context.Context) *SubjectDigest {
	node, err := sdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SubjectDigest ID from the query.
// Returns a *NotFoundError when no SubjectDigest ID was found.
func (sdq *SubjectDigestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subjectdigest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdq *SubjectDigestQuery) FirstIDX(ctx context.Context) int {
	id, err := sdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SubjectDigest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SubjectDigest entity is found.
// Returns a *NotFoundError when no SubjectDigest entities are found.
func (sdq *SubjectDigestQuery) Only(ctx context.Context) (*SubjectDigest, error) {
	nodes, err := sdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subjectdigest.Label}
	default:
		return nil, &NotSingularError{subjectdigest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdq *SubjectDigestQuery) OnlyX(ctx context.Context) *SubjectDigest {
	node, err := sdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SubjectDigest ID in the query.
// Returns a *NotSingularError when more than one SubjectDigest ID is found.
// Returns a *NotFoundError when no entities are found.
func (sdq *SubjectDigestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subjectdigest.Label}
	default:
		err = &NotSingularError{subjectdigest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdq *SubjectDigestQuery) OnlyIDX(ctx context.Context) int {
	id, err := sdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SubjectDigests.
func (sdq *SubjectDigestQuery) All(ctx context.Context) ([]*SubjectDigest, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sdq *SubjectDigestQuery) AllX(ctx context.Context) []*SubjectDigest {
	nodes, err := sdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SubjectDigest IDs.
func (sdq *SubjectDigestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sdq.Select(subjectdigest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdq *SubjectDigestQuery) IDsX(ctx context.Context) []int {
	ids, err := sdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdq *SubjectDigestQuery) Count(ctx context.Context) (int, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sdq *SubjectDigestQuery) CountX(ctx context.Context) int {
	count, err := sdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdq *SubjectDigestQuery) Exist(ctx context.Context) (bool, error) {
	if err := sdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sdq *SubjectDigestQuery) ExistX(ctx context.Context) bool {
	exist, err := sdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubjectDigestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdq *SubjectDigestQuery) Clone() *SubjectDigestQuery {
	if sdq == nil {
		return nil
	}
	return &SubjectDigestQuery{
		config:      sdq.config,
		limit:       sdq.limit,
		offset:      sdq.offset,
		order:       append([]OrderFunc{}, sdq.order...),
		predicates:  append([]predicate.SubjectDigest{}, sdq.predicates...),
		withSubject: sdq.withSubject.Clone(),
		// clone intermediate query.
		sql:    sdq.sql.Clone(),
		path:   sdq.path,
		unique: sdq.unique,
	}
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (sdq *SubjectDigestQuery) WithSubject(opts ...func(*SubjectQuery)) *SubjectDigestQuery {
	query := &SubjectQuery{config: sdq.config}
	for _, opt := range opts {
		opt(query)
	}
	sdq.withSubject = query
	return sdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Algorithm string `json:"algorithm,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SubjectDigest.Query().
//		GroupBy(subjectdigest.FieldAlgorithm).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sdq *SubjectDigestQuery) GroupBy(field string, fields ...string) *SubjectDigestGroupBy {
	grbuild := &SubjectDigestGroupBy{config: sdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sdq.sqlQuery(ctx), nil
	}
	grbuild.label = subjectdigest.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Algorithm string `json:"algorithm,omitempty"`
//	}
//
//	client.SubjectDigest.Query().
//		Select(subjectdigest.FieldAlgorithm).
//		Scan(ctx, &v)
func (sdq *SubjectDigestQuery) Select(fields ...string) *SubjectDigestSelect {
	sdq.fields = append(sdq.fields, fields...)
	selbuild := &SubjectDigestSelect{SubjectDigestQuery: sdq}
	selbuild.label = subjectdigest.Label
	selbuild.flds, selbuild.scan = &sdq.fields, selbuild.Scan
	return selbuild
}

func (sdq *SubjectDigestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sdq.fields {
		if !subjectdigest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sdq.path != nil {
		prev, err := sdq.path(ctx)
		if err != nil {
			return err
		}
		sdq.sql = prev
	}
	return nil
}

func (sdq *SubjectDigestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SubjectDigest, error) {
	var (
		nodes       = []*SubjectDigest{}
		withFKs     = sdq.withFKs
		_spec       = sdq.querySpec()
		loadedTypes = [1]bool{
			sdq.withSubject != nil,
		}
	)
	if sdq.withSubject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subjectdigest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SubjectDigest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SubjectDigest{config: sdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sdq.withSubject; query != nil {
		if err := sdq.loadSubject(ctx, query, nodes, nil,
			func(n *SubjectDigest, e *Subject) { n.Edges.Subject = e }); err != nil {
			return nil, err
		}
	}
	for i := range sdq.loadTotal {
		if err := sdq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sdq *SubjectDigestQuery) loadSubject(ctx context.Context, query *SubjectQuery, nodes []*SubjectDigest, init func(*SubjectDigest), assign func(*SubjectDigest, *Subject)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SubjectDigest)
	for i := range nodes {
		if nodes[i].subject_subject_digests == nil {
			continue
		}
		fk := *nodes[i].subject_subject_digests
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(subject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_subject_digests" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sdq *SubjectDigestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdq.querySpec()
	if len(sdq.modifiers) > 0 {
		_spec.Modifiers = sdq.modifiers
	}
	_spec.Node.Columns = sdq.fields
	if len(sdq.fields) > 0 {
		_spec.Unique = sdq.unique != nil && *sdq.unique
	}
	return sqlgraph.CountNodes(ctx, sdq.driver, _spec)
}

func (sdq *SubjectDigestQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sdq *SubjectDigestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subjectdigest.Table,
			Columns: subjectdigest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subjectdigest.FieldID,
			},
		},
		From:   sdq.sql,
		Unique: true,
	}
	if unique := sdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subjectdigest.FieldID)
		for i := range fields {
			if fields[i] != subjectdigest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sdq *SubjectDigestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdq.driver.Dialect())
	t1 := builder.Table(subjectdigest.Table)
	columns := sdq.fields
	if len(columns) == 0 {
		columns = subjectdigest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sdq.sql != nil {
		selector = sdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sdq.unique != nil && *sdq.unique {
		selector.Distinct()
	}
	for _, p := range sdq.predicates {
		p(selector)
	}
	for _, p := range sdq.order {
		p(selector)
	}
	if offset := sdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubjectDigestGroupBy is the group-by builder for SubjectDigest entities.
type SubjectDigestGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SubjectDigestGroupBy) Aggregate(fns ...AggregateFunc) *SubjectDigestGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sdgb *SubjectDigestGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sdgb.path(ctx)
	if err != nil {
		return err
	}
	sdgb.sql = query
	return sdgb.sqlScan(ctx, v)
}

func (sdgb *SubjectDigestGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sdgb.fields {
		if !subjectdigest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sdgb *SubjectDigestGroupBy) sqlQuery() *sql.Selector {
	selector := sdgb.sql.Select()
	aggregation := make([]string, 0, len(sdgb.fns))
	for _, fn := range sdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sdgb.fields)+len(sdgb.fns))
		for _, f := range sdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sdgb.fields...)...)
}

// SubjectDigestSelect is the builder for selecting fields of SubjectDigest entities.
type SubjectDigestSelect struct {
	*SubjectDigestQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SubjectDigestSelect) Scan(ctx context.Context, v any) error {
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	sds.sql = sds.SubjectDigestQuery.sqlQuery(ctx)
	return sds.sqlScan(ctx, v)
}

func (sds *SubjectDigestSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sds.sql.Query()
	if err := sds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
