// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"message-push/app/receiver/common/model/po/ent/demo"
	"message-push/app/receiver/common/model/po/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DemoQuery is the builder for querying Demo entities.
type DemoQuery struct {
	config
	ctx        *QueryContext
	order      []demo.OrderOption
	inters     []Interceptor
	predicates []predicate.Demo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DemoQuery builder.
func (dq *DemoQuery) Where(ps ...predicate.Demo) *DemoQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DemoQuery) Limit(limit int) *DemoQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DemoQuery) Offset(offset int) *DemoQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DemoQuery) Unique(unique bool) *DemoQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DemoQuery) Order(o ...demo.OrderOption) *DemoQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// First returns the first Demo entity from the query.
// Returns a *NotFoundError when no Demo was found.
func (dq *DemoQuery) First(ctx context.Context) (*Demo, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{demo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DemoQuery) FirstX(ctx context.Context) *Demo {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Demo ID from the query.
// Returns a *NotFoundError when no Demo ID was found.
func (dq *DemoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{demo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DemoQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Demo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Demo entity is found.
// Returns a *NotFoundError when no Demo entities are found.
func (dq *DemoQuery) Only(ctx context.Context) (*Demo, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{demo.Label}
	default:
		return nil, &NotSingularError{demo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DemoQuery) OnlyX(ctx context.Context) *Demo {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Demo ID in the query.
// Returns a *NotSingularError when more than one Demo ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DemoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{demo.Label}
	default:
		err = &NotSingularError{demo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DemoQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Demos.
func (dq *DemoQuery) All(ctx context.Context) ([]*Demo, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Demo, *DemoQuery]()
	return withInterceptors[[]*Demo](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DemoQuery) AllX(ctx context.Context) []*Demo {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Demo IDs.
func (dq *DemoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(demo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DemoQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DemoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DemoQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DemoQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DemoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DemoQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DemoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DemoQuery) Clone() *DemoQuery {
	if dq == nil {
		return nil
	}
	return &DemoQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]demo.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Demo{}, dq.predicates...),
		// clone intermediate query.
		sql:       dq.sql.Clone(),
		path:      dq.path,
		modifiers: append([]func(*sql.Selector){}, dq.modifiers...),
	}
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
//	client.Demo.Query().
//		GroupBy(demo.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DemoQuery) GroupBy(field string, fields ...string) *DemoGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DemoGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = demo.Label
	grbuild.scan = grbuild.Scan
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
//	client.Demo.Query().
//		Select(demo.FieldName).
//		Scan(ctx, &v)
func (dq *DemoQuery) Select(fields ...string) *DemoSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DemoSelect{DemoQuery: dq}
	sbuild.label = demo.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DemoSelect configured with the given aggregations.
func (dq *DemoQuery) Aggregate(fns ...AggregateFunc) *DemoSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DemoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !demo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DemoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Demo, error) {
	var (
		nodes = []*Demo{}
		_spec = dq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Demo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Demo{config: dq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dq *DemoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DemoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(demo.Table, demo.Columns, sqlgraph.NewFieldSpec(demo.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, demo.FieldID)
		for i := range fields {
			if fields[i] != demo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DemoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(demo.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = demo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dq.modifiers {
		m(selector)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dq *DemoQuery) ForUpdate(opts ...sql.LockOption) *DemoQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dq *DemoQuery) ForShare(opts ...sql.LockOption) *DemoQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dq *DemoQuery) Modify(modifiers ...func(s *sql.Selector)) *DemoSelect {
	dq.modifiers = append(dq.modifiers, modifiers...)
	return dq.Select()
}

// DemoGroupBy is the group-by builder for Demo entities.
type DemoGroupBy struct {
	selector
	build *DemoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DemoGroupBy) Aggregate(fns ...AggregateFunc) *DemoGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DemoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DemoQuery, *DemoGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DemoGroupBy) sqlScan(ctx context.Context, root *DemoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DemoSelect is the builder for selecting fields of Demo entities.
type DemoSelect struct {
	*DemoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DemoSelect) Aggregate(fns ...AggregateFunc) *DemoSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DemoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DemoQuery, *DemoSelect](ctx, ds.DemoQuery, ds, ds.inters, v)
}

func (ds *DemoSelect) sqlScan(ctx context.Context, root *DemoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ds *DemoSelect) Modify(modifiers ...func(s *sql.Selector)) *DemoSelect {
	ds.modifiers = append(ds.modifiers, modifiers...)
	return ds
}
