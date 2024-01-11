// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/billingaddress"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// BillingAddressQuery is the builder for querying BillingAddress entities.
type BillingAddressQuery struct {
	config
	ctx          *QueryContext
	order        []billingaddress.OrderOption
	inters       []Interceptor
	predicates   []predicate.BillingAddress
	withCustomer *CustomerQuery
	withNotes    *NoteQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingAddressQuery builder.
func (baq *BillingAddressQuery) Where(ps ...predicate.BillingAddress) *BillingAddressQuery {
	baq.predicates = append(baq.predicates, ps...)
	return baq
}

// Limit the number of records to be returned by this query.
func (baq *BillingAddressQuery) Limit(limit int) *BillingAddressQuery {
	baq.ctx.Limit = &limit
	return baq
}

// Offset to start from.
func (baq *BillingAddressQuery) Offset(offset int) *BillingAddressQuery {
	baq.ctx.Offset = &offset
	return baq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (baq *BillingAddressQuery) Unique(unique bool) *BillingAddressQuery {
	baq.ctx.Unique = &unique
	return baq
}

// Order specifies how the records should be ordered.
func (baq *BillingAddressQuery) Order(o ...billingaddress.OrderOption) *BillingAddressQuery {
	baq.order = append(baq.order, o...)
	return baq
}

// QueryCustomer chains the current query on the "customer" edge.
func (baq *BillingAddressQuery) QueryCustomer() *CustomerQuery {
	query := (&CustomerClient{config: baq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := baq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := baq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billingaddress.Table, billingaddress.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, billingaddress.CustomerTable, billingaddress.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(baq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotes chains the current query on the "notes" edge.
func (baq *BillingAddressQuery) QueryNotes() *NoteQuery {
	query := (&NoteClient{config: baq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := baq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := baq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billingaddress.Table, billingaddress.FieldID, selector),
			sqlgraph.To(note.Table, note.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, billingaddress.NotesTable, billingaddress.NotesColumn),
		)
		fromU = sqlgraph.SetNeighbors(baq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BillingAddress entity from the query.
// Returns a *NotFoundError when no BillingAddress was found.
func (baq *BillingAddressQuery) First(ctx context.Context) (*BillingAddress, error) {
	nodes, err := baq.Limit(1).All(setContextOp(ctx, baq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billingaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (baq *BillingAddressQuery) FirstX(ctx context.Context) *BillingAddress {
	node, err := baq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingAddress ID from the query.
// Returns a *NotFoundError when no BillingAddress ID was found.
func (baq *BillingAddressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = baq.Limit(1).IDs(setContextOp(ctx, baq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billingaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (baq *BillingAddressQuery) FirstIDX(ctx context.Context) int {
	id, err := baq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingAddress entity is found.
// Returns a *NotFoundError when no BillingAddress entities are found.
func (baq *BillingAddressQuery) Only(ctx context.Context) (*BillingAddress, error) {
	nodes, err := baq.Limit(2).All(setContextOp(ctx, baq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billingaddress.Label}
	default:
		return nil, &NotSingularError{billingaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (baq *BillingAddressQuery) OnlyX(ctx context.Context) *BillingAddress {
	node, err := baq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingAddress ID in the query.
// Returns a *NotSingularError when more than one BillingAddress ID is found.
// Returns a *NotFoundError when no entities are found.
func (baq *BillingAddressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = baq.Limit(2).IDs(setContextOp(ctx, baq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billingaddress.Label}
	default:
		err = &NotSingularError{billingaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (baq *BillingAddressQuery) OnlyIDX(ctx context.Context) int {
	id, err := baq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingAddresses.
func (baq *BillingAddressQuery) All(ctx context.Context) ([]*BillingAddress, error) {
	ctx = setContextOp(ctx, baq.ctx, "All")
	if err := baq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingAddress, *BillingAddressQuery]()
	return withInterceptors[[]*BillingAddress](ctx, baq, qr, baq.inters)
}

// AllX is like All, but panics if an error occurs.
func (baq *BillingAddressQuery) AllX(ctx context.Context) []*BillingAddress {
	nodes, err := baq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingAddress IDs.
func (baq *BillingAddressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if baq.ctx.Unique == nil && baq.path != nil {
		baq.Unique(true)
	}
	ctx = setContextOp(ctx, baq.ctx, "IDs")
	if err = baq.Select(billingaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (baq *BillingAddressQuery) IDsX(ctx context.Context) []int {
	ids, err := baq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (baq *BillingAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, baq.ctx, "Count")
	if err := baq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, baq, querierCount[*BillingAddressQuery](), baq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (baq *BillingAddressQuery) CountX(ctx context.Context) int {
	count, err := baq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (baq *BillingAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, baq.ctx, "Exist")
	switch _, err := baq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (baq *BillingAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := baq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (baq *BillingAddressQuery) Clone() *BillingAddressQuery {
	if baq == nil {
		return nil
	}
	return &BillingAddressQuery{
		config:       baq.config,
		ctx:          baq.ctx.Clone(),
		order:        append([]billingaddress.OrderOption{}, baq.order...),
		inters:       append([]Interceptor{}, baq.inters...),
		predicates:   append([]predicate.BillingAddress{}, baq.predicates...),
		withCustomer: baq.withCustomer.Clone(),
		withNotes:    baq.withNotes.Clone(),
		// clone intermediate query.
		sql:  baq.sql.Clone(),
		path: baq.path,
	}
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (baq *BillingAddressQuery) WithCustomer(opts ...func(*CustomerQuery)) *BillingAddressQuery {
	query := (&CustomerClient{config: baq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	baq.withCustomer = query
	return baq
}

// WithNotes tells the query-builder to eager-load the nodes that are connected to
// the "notes" edge. The optional arguments are used to configure the query builder of the edge.
func (baq *BillingAddressQuery) WithNotes(opts ...func(*NoteQuery)) *BillingAddressQuery {
	query := (&NoteClient{config: baq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	baq.withNotes = query
	return baq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		City string `json:"city,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillingAddress.Query().
//		GroupBy(billingaddress.FieldCity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (baq *BillingAddressQuery) GroupBy(field string, fields ...string) *BillingAddressGroupBy {
	baq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingAddressGroupBy{build: baq}
	grbuild.flds = &baq.ctx.Fields
	grbuild.label = billingaddress.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		City string `json:"city,omitempty"`
//	}
//
//	client.BillingAddress.Query().
//		Select(billingaddress.FieldCity).
//		Scan(ctx, &v)
func (baq *BillingAddressQuery) Select(fields ...string) *BillingAddressSelect {
	baq.ctx.Fields = append(baq.ctx.Fields, fields...)
	sbuild := &BillingAddressSelect{BillingAddressQuery: baq}
	sbuild.label = billingaddress.Label
	sbuild.flds, sbuild.scan = &baq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingAddressSelect configured with the given aggregations.
func (baq *BillingAddressQuery) Aggregate(fns ...AggregateFunc) *BillingAddressSelect {
	return baq.Select().Aggregate(fns...)
}

func (baq *BillingAddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range baq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, baq); err != nil {
				return err
			}
		}
	}
	for _, f := range baq.ctx.Fields {
		if !billingaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if baq.path != nil {
		prev, err := baq.path(ctx)
		if err != nil {
			return err
		}
		baq.sql = prev
	}
	return nil
}

func (baq *BillingAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingAddress, error) {
	var (
		nodes       = []*BillingAddress{}
		withFKs     = baq.withFKs
		_spec       = baq.querySpec()
		loadedTypes = [2]bool{
			baq.withCustomer != nil,
			baq.withNotes != nil,
		}
	)
	if baq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, billingaddress.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingAddress{config: baq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, baq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := baq.withCustomer; query != nil {
		if err := baq.loadCustomer(ctx, query, nodes, nil,
			func(n *BillingAddress, e *Customer) { n.Edges.Customer = e }); err != nil {
			return nil, err
		}
	}
	if query := baq.withNotes; query != nil {
		if err := baq.loadNotes(ctx, query, nodes,
			func(n *BillingAddress) { n.Edges.Notes = []*Note{} },
			func(n *BillingAddress, e *Note) { n.Edges.Notes = append(n.Edges.Notes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (baq *BillingAddressQuery) loadCustomer(ctx context.Context, query *CustomerQuery, nodes []*BillingAddress, init func(*BillingAddress), assign func(*BillingAddress, *Customer)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BillingAddress)
	for i := range nodes {
		if nodes[i].customer_billing_addresses == nil {
			continue
		}
		fk := *nodes[i].customer_billing_addresses
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(customer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_billing_addresses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (baq *BillingAddressQuery) loadNotes(ctx context.Context, query *NoteQuery, nodes []*BillingAddress, init func(*BillingAddress), assign func(*BillingAddress, *Note)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*BillingAddress)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Note(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(billingaddress.NotesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.billing_address_notes
		if fk == nil {
			return fmt.Errorf(`foreign-key "billing_address_notes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "billing_address_notes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (baq *BillingAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := baq.querySpec()
	_spec.Node.Columns = baq.ctx.Fields
	if len(baq.ctx.Fields) > 0 {
		_spec.Unique = baq.ctx.Unique != nil && *baq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, baq.driver, _spec)
}

func (baq *BillingAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billingaddress.Table, billingaddress.Columns, sqlgraph.NewFieldSpec(billingaddress.FieldID, field.TypeInt))
	_spec.From = baq.sql
	if unique := baq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if baq.path != nil {
		_spec.Unique = true
	}
	if fields := baq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billingaddress.FieldID)
		for i := range fields {
			if fields[i] != billingaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := baq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := baq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := baq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := baq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (baq *BillingAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(baq.driver.Dialect())
	t1 := builder.Table(billingaddress.Table)
	columns := baq.ctx.Fields
	if len(columns) == 0 {
		columns = billingaddress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if baq.sql != nil {
		selector = baq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if baq.ctx.Unique != nil && *baq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range baq.predicates {
		p(selector)
	}
	for _, p := range baq.order {
		p(selector)
	}
	if offset := baq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := baq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BillingAddressGroupBy is the group-by builder for BillingAddress entities.
type BillingAddressGroupBy struct {
	selector
	build *BillingAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bagb *BillingAddressGroupBy) Aggregate(fns ...AggregateFunc) *BillingAddressGroupBy {
	bagb.fns = append(bagb.fns, fns...)
	return bagb
}

// Scan applies the selector query and scans the result into the given value.
func (bagb *BillingAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bagb.build.ctx, "GroupBy")
	if err := bagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingAddressQuery, *BillingAddressGroupBy](ctx, bagb.build, bagb, bagb.build.inters, v)
}

func (bagb *BillingAddressGroupBy) sqlScan(ctx context.Context, root *BillingAddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bagb.fns))
	for _, fn := range bagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bagb.flds)+len(bagb.fns))
		for _, f := range *bagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingAddressSelect is the builder for selecting fields of BillingAddress entities.
type BillingAddressSelect struct {
	*BillingAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bas *BillingAddressSelect) Aggregate(fns ...AggregateFunc) *BillingAddressSelect {
	bas.fns = append(bas.fns, fns...)
	return bas
}

// Scan applies the selector query and scans the result into the given value.
func (bas *BillingAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bas.ctx, "Select")
	if err := bas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingAddressQuery, *BillingAddressSelect](ctx, bas.BillingAddressQuery, bas, bas.inters, v)
}

func (bas *BillingAddressSelect) sqlScan(ctx context.Context, root *BillingAddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bas.fns))
	for _, fn := range bas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}