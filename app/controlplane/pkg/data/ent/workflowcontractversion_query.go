// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontractversion"
	"github.com/google/uuid"
)

// WorkflowContractVersionQuery is the builder for querying WorkflowContractVersion entities.
type WorkflowContractVersionQuery struct {
	config
	ctx          *QueryContext
	order        []workflowcontractversion.OrderOption
	inters       []Interceptor
	predicates   []predicate.WorkflowContractVersion
	withContract *WorkflowContractQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkflowContractVersionQuery builder.
func (wcvq *WorkflowContractVersionQuery) Where(ps ...predicate.WorkflowContractVersion) *WorkflowContractVersionQuery {
	wcvq.predicates = append(wcvq.predicates, ps...)
	return wcvq
}

// Limit the number of records to be returned by this query.
func (wcvq *WorkflowContractVersionQuery) Limit(limit int) *WorkflowContractVersionQuery {
	wcvq.ctx.Limit = &limit
	return wcvq
}

// Offset to start from.
func (wcvq *WorkflowContractVersionQuery) Offset(offset int) *WorkflowContractVersionQuery {
	wcvq.ctx.Offset = &offset
	return wcvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wcvq *WorkflowContractVersionQuery) Unique(unique bool) *WorkflowContractVersionQuery {
	wcvq.ctx.Unique = &unique
	return wcvq
}

// Order specifies how the records should be ordered.
func (wcvq *WorkflowContractVersionQuery) Order(o ...workflowcontractversion.OrderOption) *WorkflowContractVersionQuery {
	wcvq.order = append(wcvq.order, o...)
	return wcvq
}

// QueryContract chains the current query on the "contract" edge.
func (wcvq *WorkflowContractVersionQuery) QueryContract() *WorkflowContractQuery {
	query := (&WorkflowContractClient{config: wcvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wcvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wcvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowcontractversion.Table, workflowcontractversion.FieldID, selector),
			sqlgraph.To(workflowcontract.Table, workflowcontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workflowcontractversion.ContractTable, workflowcontractversion.ContractColumn),
		)
		fromU = sqlgraph.SetNeighbors(wcvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkflowContractVersion entity from the query.
// Returns a *NotFoundError when no WorkflowContractVersion was found.
func (wcvq *WorkflowContractVersionQuery) First(ctx context.Context) (*WorkflowContractVersion, error) {
	nodes, err := wcvq.Limit(1).All(setContextOp(ctx, wcvq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workflowcontractversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) FirstX(ctx context.Context) *WorkflowContractVersion {
	node, err := wcvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkflowContractVersion ID from the query.
// Returns a *NotFoundError when no WorkflowContractVersion ID was found.
func (wcvq *WorkflowContractVersionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wcvq.Limit(1).IDs(setContextOp(ctx, wcvq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workflowcontractversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wcvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkflowContractVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkflowContractVersion entity is found.
// Returns a *NotFoundError when no WorkflowContractVersion entities are found.
func (wcvq *WorkflowContractVersionQuery) Only(ctx context.Context) (*WorkflowContractVersion, error) {
	nodes, err := wcvq.Limit(2).All(setContextOp(ctx, wcvq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workflowcontractversion.Label}
	default:
		return nil, &NotSingularError{workflowcontractversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) OnlyX(ctx context.Context) *WorkflowContractVersion {
	node, err := wcvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkflowContractVersion ID in the query.
// Returns a *NotSingularError when more than one WorkflowContractVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (wcvq *WorkflowContractVersionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wcvq.Limit(2).IDs(setContextOp(ctx, wcvq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workflowcontractversion.Label}
	default:
		err = &NotSingularError{workflowcontractversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wcvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkflowContractVersions.
func (wcvq *WorkflowContractVersionQuery) All(ctx context.Context) ([]*WorkflowContractVersion, error) {
	ctx = setContextOp(ctx, wcvq.ctx, ent.OpQueryAll)
	if err := wcvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkflowContractVersion, *WorkflowContractVersionQuery]()
	return withInterceptors[[]*WorkflowContractVersion](ctx, wcvq, qr, wcvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) AllX(ctx context.Context) []*WorkflowContractVersion {
	nodes, err := wcvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkflowContractVersion IDs.
func (wcvq *WorkflowContractVersionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wcvq.ctx.Unique == nil && wcvq.path != nil {
		wcvq.Unique(true)
	}
	ctx = setContextOp(ctx, wcvq.ctx, ent.OpQueryIDs)
	if err = wcvq.Select(workflowcontractversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wcvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wcvq *WorkflowContractVersionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wcvq.ctx, ent.OpQueryCount)
	if err := wcvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wcvq, querierCount[*WorkflowContractVersionQuery](), wcvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) CountX(ctx context.Context) int {
	count, err := wcvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wcvq *WorkflowContractVersionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wcvq.ctx, ent.OpQueryExist)
	switch _, err := wcvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wcvq *WorkflowContractVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := wcvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkflowContractVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wcvq *WorkflowContractVersionQuery) Clone() *WorkflowContractVersionQuery {
	if wcvq == nil {
		return nil
	}
	return &WorkflowContractVersionQuery{
		config:       wcvq.config,
		ctx:          wcvq.ctx.Clone(),
		order:        append([]workflowcontractversion.OrderOption{}, wcvq.order...),
		inters:       append([]Interceptor{}, wcvq.inters...),
		predicates:   append([]predicate.WorkflowContractVersion{}, wcvq.predicates...),
		withContract: wcvq.withContract.Clone(),
		// clone intermediate query.
		sql:  wcvq.sql.Clone(),
		path: wcvq.path,
	}
}

// WithContract tells the query-builder to eager-load the nodes that are connected to
// the "contract" edge. The optional arguments are used to configure the query builder of the edge.
func (wcvq *WorkflowContractVersionQuery) WithContract(opts ...func(*WorkflowContractQuery)) *WorkflowContractVersionQuery {
	query := (&WorkflowContractClient{config: wcvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wcvq.withContract = query
	return wcvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Body []byte `json:"body,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkflowContractVersion.Query().
//		GroupBy(workflowcontractversion.FieldBody).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wcvq *WorkflowContractVersionQuery) GroupBy(field string, fields ...string) *WorkflowContractVersionGroupBy {
	wcvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkflowContractVersionGroupBy{build: wcvq}
	grbuild.flds = &wcvq.ctx.Fields
	grbuild.label = workflowcontractversion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Body []byte `json:"body,omitempty"`
//	}
//
//	client.WorkflowContractVersion.Query().
//		Select(workflowcontractversion.FieldBody).
//		Scan(ctx, &v)
func (wcvq *WorkflowContractVersionQuery) Select(fields ...string) *WorkflowContractVersionSelect {
	wcvq.ctx.Fields = append(wcvq.ctx.Fields, fields...)
	sbuild := &WorkflowContractVersionSelect{WorkflowContractVersionQuery: wcvq}
	sbuild.label = workflowcontractversion.Label
	sbuild.flds, sbuild.scan = &wcvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkflowContractVersionSelect configured with the given aggregations.
func (wcvq *WorkflowContractVersionQuery) Aggregate(fns ...AggregateFunc) *WorkflowContractVersionSelect {
	return wcvq.Select().Aggregate(fns...)
}

func (wcvq *WorkflowContractVersionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wcvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wcvq); err != nil {
				return err
			}
		}
	}
	for _, f := range wcvq.ctx.Fields {
		if !workflowcontractversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wcvq.path != nil {
		prev, err := wcvq.path(ctx)
		if err != nil {
			return err
		}
		wcvq.sql = prev
	}
	return nil
}

func (wcvq *WorkflowContractVersionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkflowContractVersion, error) {
	var (
		nodes       = []*WorkflowContractVersion{}
		withFKs     = wcvq.withFKs
		_spec       = wcvq.querySpec()
		loadedTypes = [1]bool{
			wcvq.withContract != nil,
		}
	)
	if wcvq.withContract != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workflowcontractversion.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkflowContractVersion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkflowContractVersion{config: wcvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wcvq.modifiers) > 0 {
		_spec.Modifiers = wcvq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wcvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wcvq.withContract; query != nil {
		if err := wcvq.loadContract(ctx, query, nodes, nil,
			func(n *WorkflowContractVersion, e *WorkflowContract) { n.Edges.Contract = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wcvq *WorkflowContractVersionQuery) loadContract(ctx context.Context, query *WorkflowContractQuery, nodes []*WorkflowContractVersion, init func(*WorkflowContractVersion), assign func(*WorkflowContractVersion, *WorkflowContract)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkflowContractVersion)
	for i := range nodes {
		if nodes[i].workflow_contract_versions == nil {
			continue
		}
		fk := *nodes[i].workflow_contract_versions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workflowcontract.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workflow_contract_versions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wcvq *WorkflowContractVersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wcvq.querySpec()
	if len(wcvq.modifiers) > 0 {
		_spec.Modifiers = wcvq.modifiers
	}
	_spec.Node.Columns = wcvq.ctx.Fields
	if len(wcvq.ctx.Fields) > 0 {
		_spec.Unique = wcvq.ctx.Unique != nil && *wcvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wcvq.driver, _spec)
}

func (wcvq *WorkflowContractVersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workflowcontractversion.Table, workflowcontractversion.Columns, sqlgraph.NewFieldSpec(workflowcontractversion.FieldID, field.TypeUUID))
	_spec.From = wcvq.sql
	if unique := wcvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wcvq.path != nil {
		_spec.Unique = true
	}
	if fields := wcvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowcontractversion.FieldID)
		for i := range fields {
			if fields[i] != workflowcontractversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wcvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wcvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wcvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wcvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wcvq *WorkflowContractVersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wcvq.driver.Dialect())
	t1 := builder.Table(workflowcontractversion.Table)
	columns := wcvq.ctx.Fields
	if len(columns) == 0 {
		columns = workflowcontractversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wcvq.sql != nil {
		selector = wcvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wcvq.ctx.Unique != nil && *wcvq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wcvq.modifiers {
		m(selector)
	}
	for _, p := range wcvq.predicates {
		p(selector)
	}
	for _, p := range wcvq.order {
		p(selector)
	}
	if offset := wcvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wcvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (wcvq *WorkflowContractVersionQuery) ForUpdate(opts ...sql.LockOption) *WorkflowContractVersionQuery {
	if wcvq.driver.Dialect() == dialect.Postgres {
		wcvq.Unique(false)
	}
	wcvq.modifiers = append(wcvq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return wcvq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (wcvq *WorkflowContractVersionQuery) ForShare(opts ...sql.LockOption) *WorkflowContractVersionQuery {
	if wcvq.driver.Dialect() == dialect.Postgres {
		wcvq.Unique(false)
	}
	wcvq.modifiers = append(wcvq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return wcvq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wcvq *WorkflowContractVersionQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkflowContractVersionSelect {
	wcvq.modifiers = append(wcvq.modifiers, modifiers...)
	return wcvq.Select()
}

// WorkflowContractVersionGroupBy is the group-by builder for WorkflowContractVersion entities.
type WorkflowContractVersionGroupBy struct {
	selector
	build *WorkflowContractVersionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wcvgb *WorkflowContractVersionGroupBy) Aggregate(fns ...AggregateFunc) *WorkflowContractVersionGroupBy {
	wcvgb.fns = append(wcvgb.fns, fns...)
	return wcvgb
}

// Scan applies the selector query and scans the result into the given value.
func (wcvgb *WorkflowContractVersionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcvgb.build.ctx, ent.OpQueryGroupBy)
	if err := wcvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowContractVersionQuery, *WorkflowContractVersionGroupBy](ctx, wcvgb.build, wcvgb, wcvgb.build.inters, v)
}

func (wcvgb *WorkflowContractVersionGroupBy) sqlScan(ctx context.Context, root *WorkflowContractVersionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wcvgb.fns))
	for _, fn := range wcvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wcvgb.flds)+len(wcvgb.fns))
		for _, f := range *wcvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wcvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkflowContractVersionSelect is the builder for selecting fields of WorkflowContractVersion entities.
type WorkflowContractVersionSelect struct {
	*WorkflowContractVersionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wcvs *WorkflowContractVersionSelect) Aggregate(fns ...AggregateFunc) *WorkflowContractVersionSelect {
	wcvs.fns = append(wcvs.fns, fns...)
	return wcvs
}

// Scan applies the selector query and scans the result into the given value.
func (wcvs *WorkflowContractVersionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcvs.ctx, ent.OpQuerySelect)
	if err := wcvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowContractVersionQuery, *WorkflowContractVersionSelect](ctx, wcvs.WorkflowContractVersionQuery, wcvs, wcvs.inters, v)
}

func (wcvs *WorkflowContractVersionSelect) sqlScan(ctx context.Context, root *WorkflowContractVersionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wcvs.fns))
	for _, fn := range wcvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wcvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wcvs *WorkflowContractVersionSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkflowContractVersionSelect {
	wcvs.modifiers = append(wcvs.modifiers, modifiers...)
	return wcvs
}
