// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/dependsonskipped"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/predicate"
)

// DependsOnSkippedQuery is the builder for querying DependsOnSkipped entities.
type DependsOnSkippedQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DependsOnSkipped
	// eager-loading edges.
	withSkipped *ImplicitSkippedMessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DependsOnSkippedQuery builder.
func (dosq *DependsOnSkippedQuery) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedQuery {
	dosq.predicates = append(dosq.predicates, ps...)
	return dosq
}

// Limit adds a limit step to the query.
func (dosq *DependsOnSkippedQuery) Limit(limit int) *DependsOnSkippedQuery {
	dosq.limit = &limit
	return dosq
}

// Offset adds an offset step to the query.
func (dosq *DependsOnSkippedQuery) Offset(offset int) *DependsOnSkippedQuery {
	dosq.offset = &offset
	return dosq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dosq *DependsOnSkippedQuery) Unique(unique bool) *DependsOnSkippedQuery {
	dosq.unique = &unique
	return dosq
}

// Order adds an order step to the query.
func (dosq *DependsOnSkippedQuery) Order(o ...OrderFunc) *DependsOnSkippedQuery {
	dosq.order = append(dosq.order, o...)
	return dosq
}

// QuerySkipped chains the current query on the "skipped" edge.
func (dosq *DependsOnSkippedQuery) QuerySkipped() *ImplicitSkippedMessageQuery {
	query := &ImplicitSkippedMessageQuery{config: dosq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dosq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dosq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dependsonskipped.Table, dependsonskipped.FieldID, selector),
			sqlgraph.To(implicitskippedmessage.Table, implicitskippedmessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dependsonskipped.SkippedTable, dependsonskipped.SkippedColumn),
		)
		fromU = sqlgraph.SetNeighbors(dosq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DependsOnSkipped entity from the query.
// Returns a *NotFoundError when no DependsOnSkipped was found.
func (dosq *DependsOnSkippedQuery) First(ctx context.Context) (*DependsOnSkipped, error) {
	nodes, err := dosq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dependsonskipped.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) FirstX(ctx context.Context) *DependsOnSkipped {
	node, err := dosq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DependsOnSkipped ID from the query.
// Returns a *NotFoundError when no DependsOnSkipped ID was found.
func (dosq *DependsOnSkippedQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dosq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dependsonskipped.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) FirstIDX(ctx context.Context) int {
	id, err := dosq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DependsOnSkipped entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DependsOnSkipped entity is found.
// Returns a *NotFoundError when no DependsOnSkipped entities are found.
func (dosq *DependsOnSkippedQuery) Only(ctx context.Context) (*DependsOnSkipped, error) {
	nodes, err := dosq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dependsonskipped.Label}
	default:
		return nil, &NotSingularError{dependsonskipped.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) OnlyX(ctx context.Context) *DependsOnSkipped {
	node, err := dosq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DependsOnSkipped ID in the query.
// Returns a *NotSingularError when more than one DependsOnSkipped ID is found.
// Returns a *NotFoundError when no entities are found.
func (dosq *DependsOnSkippedQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dosq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = &NotSingularError{dependsonskipped.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) OnlyIDX(ctx context.Context) int {
	id, err := dosq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DependsOnSkippeds.
func (dosq *DependsOnSkippedQuery) All(ctx context.Context) ([]*DependsOnSkipped, error) {
	if err := dosq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dosq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) AllX(ctx context.Context) []*DependsOnSkipped {
	nodes, err := dosq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DependsOnSkipped IDs.
func (dosq *DependsOnSkippedQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dosq.Select(dependsonskipped.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) IDsX(ctx context.Context) []int {
	ids, err := dosq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dosq *DependsOnSkippedQuery) Count(ctx context.Context) (int, error) {
	if err := dosq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dosq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) CountX(ctx context.Context) int {
	count, err := dosq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dosq *DependsOnSkippedQuery) Exist(ctx context.Context) (bool, error) {
	if err := dosq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dosq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) ExistX(ctx context.Context) bool {
	exist, err := dosq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DependsOnSkippedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dosq *DependsOnSkippedQuery) Clone() *DependsOnSkippedQuery {
	if dosq == nil {
		return nil
	}
	return &DependsOnSkippedQuery{
		config:      dosq.config,
		limit:       dosq.limit,
		offset:      dosq.offset,
		order:       append([]OrderFunc{}, dosq.order...),
		predicates:  append([]predicate.DependsOnSkipped{}, dosq.predicates...),
		withSkipped: dosq.withSkipped.Clone(),
		// clone intermediate query.
		sql:    dosq.sql.Clone(),
		path:   dosq.path,
		unique: dosq.unique,
	}
}

// WithSkipped tells the query-builder to eager-load the nodes that are connected to
// the "skipped" edge. The optional arguments are used to configure the query builder of the edge.
func (dosq *DependsOnSkippedQuery) WithSkipped(opts ...func(*ImplicitSkippedMessageQuery)) *DependsOnSkippedQuery {
	query := &ImplicitSkippedMessageQuery{config: dosq.config}
	for _, opt := range opts {
		opt(query)
	}
	dosq.withSkipped = query
	return dosq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DependsOnSkipped.Query().
//		GroupBy(dependsonskipped.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dosq *DependsOnSkippedQuery) GroupBy(field string, fields ...string) *DependsOnSkippedGroupBy {
	grbuild := &DependsOnSkippedGroupBy{config: dosq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dosq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dosq.sqlQuery(ctx), nil
	}
	grbuild.label = dependsonskipped.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.DependsOnSkipped.Query().
//		Select(dependsonskipped.FieldName).
//		Scan(ctx, &v)
//
func (dosq *DependsOnSkippedQuery) Select(fields ...string) *DependsOnSkippedSelect {
	dosq.fields = append(dosq.fields, fields...)
	selbuild := &DependsOnSkippedSelect{DependsOnSkippedQuery: dosq}
	selbuild.label = dependsonskipped.Label
	selbuild.flds, selbuild.scan = &dosq.fields, selbuild.Scan
	return selbuild
}

func (dosq *DependsOnSkippedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dosq.fields {
		if !dependsonskipped.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dosq.path != nil {
		prev, err := dosq.path(ctx)
		if err != nil {
			return err
		}
		dosq.sql = prev
	}
	return nil
}

func (dosq *DependsOnSkippedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DependsOnSkipped, error) {
	var (
		nodes       = []*DependsOnSkipped{}
		_spec       = dosq.querySpec()
		loadedTypes = [1]bool{
			dosq.withSkipped != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DependsOnSkipped).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DependsOnSkipped{config: dosq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dosq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dosq.withSkipped; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DependsOnSkipped)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Skipped = []*ImplicitSkippedMessage{}
		}
		query.withFKs = true
		query.Where(predicate.ImplicitSkippedMessage(func(s *sql.Selector) {
			s.Where(sql.InValues(dependsonskipped.SkippedColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.depends_on_skipped_skipped
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "depends_on_skipped_skipped" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "depends_on_skipped_skipped" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Skipped = append(node.Edges.Skipped, n)
		}
	}

	return nodes, nil
}

func (dosq *DependsOnSkippedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dosq.querySpec()
	_spec.Node.Columns = dosq.fields
	if len(dosq.fields) > 0 {
		_spec.Unique = dosq.unique != nil && *dosq.unique
	}
	return sqlgraph.CountNodes(ctx, dosq.driver, _spec)
}

func (dosq *DependsOnSkippedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dosq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dosq *DependsOnSkippedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dependsonskipped.Table,
			Columns: dependsonskipped.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dependsonskipped.FieldID,
			},
		},
		From:   dosq.sql,
		Unique: true,
	}
	if unique := dosq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dosq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dependsonskipped.FieldID)
		for i := range fields {
			if fields[i] != dependsonskipped.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dosq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dosq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dosq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dosq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dosq *DependsOnSkippedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dosq.driver.Dialect())
	t1 := builder.Table(dependsonskipped.Table)
	columns := dosq.fields
	if len(columns) == 0 {
		columns = dependsonskipped.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dosq.sql != nil {
		selector = dosq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dosq.unique != nil && *dosq.unique {
		selector.Distinct()
	}
	for _, p := range dosq.predicates {
		p(selector)
	}
	for _, p := range dosq.order {
		p(selector)
	}
	if offset := dosq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dosq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DependsOnSkippedGroupBy is the group-by builder for DependsOnSkipped entities.
type DependsOnSkippedGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dosgb *DependsOnSkippedGroupBy) Aggregate(fns ...AggregateFunc) *DependsOnSkippedGroupBy {
	dosgb.fns = append(dosgb.fns, fns...)
	return dosgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dosgb *DependsOnSkippedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dosgb.path(ctx)
	if err != nil {
		return err
	}
	dosgb.sql = query
	return dosgb.sqlScan(ctx, v)
}

func (dosgb *DependsOnSkippedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dosgb.fields {
		if !dependsonskipped.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dosgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dosgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dosgb *DependsOnSkippedGroupBy) sqlQuery() *sql.Selector {
	selector := dosgb.sql.Select()
	aggregation := make([]string, 0, len(dosgb.fns))
	for _, fn := range dosgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dosgb.fields)+len(dosgb.fns))
		for _, f := range dosgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dosgb.fields...)...)
}

// DependsOnSkippedSelect is the builder for selecting fields of DependsOnSkipped entities.
type DependsOnSkippedSelect struct {
	*DependsOnSkippedQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (doss *DependsOnSkippedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := doss.prepareQuery(ctx); err != nil {
		return err
	}
	doss.sql = doss.DependsOnSkippedQuery.sqlQuery(ctx)
	return doss.sqlScan(ctx, v)
}

func (doss *DependsOnSkippedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := doss.sql.Query()
	if err := doss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
