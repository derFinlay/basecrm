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
	"github.com/derfinlay/basecrm/ent/address"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/predicate"
	"github.com/derfinlay/basecrm/ent/tel"
	"github.com/derfinlay/basecrm/ent/user"
)

// CustomerQuery is the builder for querying Customer entities.
type CustomerQuery struct {
	config
	ctx                *QueryContext
	order              []customer.OrderOption
	inters             []Interceptor
	predicates         []predicate.Customer
	withOrders         *OrderQuery
	withBillingAddress *AddressQuery
	withAddresses      *AddressQuery
	withPhone          *TelQuery
	withCreatedBy      *UserQuery
	withNotes          *NoteQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerQuery builder.
func (cq *CustomerQuery) Where(ps ...predicate.Customer) *CustomerQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CustomerQuery) Limit(limit int) *CustomerQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CustomerQuery) Offset(offset int) *CustomerQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CustomerQuery) Unique(unique bool) *CustomerQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CustomerQuery) Order(o ...customer.OrderOption) *CustomerQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryOrders chains the current query on the "orders" edge.
func (cq *CustomerQuery) QueryOrders() *OrderQuery {
	query := (&OrderClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.OrdersTable, customer.OrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBillingAddress chains the current query on the "billing_address" edge.
func (cq *CustomerQuery) QueryBillingAddress() *AddressQuery {
	query := (&AddressClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, customer.BillingAddressTable, customer.BillingAddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAddresses chains the current query on the "addresses" edge.
func (cq *CustomerQuery) QueryAddresses() *AddressQuery {
	query := (&AddressClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.AddressesTable, customer.AddressesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPhone chains the current query on the "phone" edge.
func (cq *CustomerQuery) QueryPhone() *TelQuery {
	query := (&TelClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(tel.Table, tel.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.PhoneTable, customer.PhoneColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedBy chains the current query on the "created_by" edge.
func (cq *CustomerQuery) QueryCreatedBy() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, customer.CreatedByTable, customer.CreatedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotes chains the current query on the "notes" edge.
func (cq *CustomerQuery) QueryNotes() *NoteQuery {
	query := (&NoteClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(note.Table, note.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.NotesTable, customer.NotesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Customer entity from the query.
// Returns a *NotFoundError when no Customer was found.
func (cq *CustomerQuery) First(ctx context.Context) (*Customer, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CustomerQuery) FirstX(ctx context.Context) *Customer {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Customer ID from the query.
// Returns a *NotFoundError when no Customer ID was found.
func (cq *CustomerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CustomerQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Customer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Customer entity is found.
// Returns a *NotFoundError when no Customer entities are found.
func (cq *CustomerQuery) Only(ctx context.Context) (*Customer, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customer.Label}
	default:
		return nil, &NotSingularError{customer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CustomerQuery) OnlyX(ctx context.Context) *Customer {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Customer ID in the query.
// Returns a *NotSingularError when more than one Customer ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CustomerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customer.Label}
	default:
		err = &NotSingularError{customer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CustomerQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Customers.
func (cq *CustomerQuery) All(ctx context.Context) ([]*Customer, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Customer, *CustomerQuery]()
	return withInterceptors[[]*Customer](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CustomerQuery) AllX(ctx context.Context) []*Customer {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Customer IDs.
func (cq *CustomerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(customer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CustomerQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CustomerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CustomerQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CustomerQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CustomerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CustomerQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CustomerQuery) Clone() *CustomerQuery {
	if cq == nil {
		return nil
	}
	return &CustomerQuery{
		config:             cq.config,
		ctx:                cq.ctx.Clone(),
		order:              append([]customer.OrderOption{}, cq.order...),
		inters:             append([]Interceptor{}, cq.inters...),
		predicates:         append([]predicate.Customer{}, cq.predicates...),
		withOrders:         cq.withOrders.Clone(),
		withBillingAddress: cq.withBillingAddress.Clone(),
		withAddresses:      cq.withAddresses.Clone(),
		withPhone:          cq.withPhone.Clone(),
		withCreatedBy:      cq.withCreatedBy.Clone(),
		withNotes:          cq.withNotes.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithOrders tells the query-builder to eager-load the nodes that are connected to
// the "orders" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithOrders(opts ...func(*OrderQuery)) *CustomerQuery {
	query := (&OrderClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withOrders = query
	return cq
}

// WithBillingAddress tells the query-builder to eager-load the nodes that are connected to
// the "billing_address" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithBillingAddress(opts ...func(*AddressQuery)) *CustomerQuery {
	query := (&AddressClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withBillingAddress = query
	return cq
}

// WithAddresses tells the query-builder to eager-load the nodes that are connected to
// the "addresses" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithAddresses(opts ...func(*AddressQuery)) *CustomerQuery {
	query := (&AddressClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAddresses = query
	return cq
}

// WithPhone tells the query-builder to eager-load the nodes that are connected to
// the "phone" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithPhone(opts ...func(*TelQuery)) *CustomerQuery {
	query := (&TelClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withPhone = query
	return cq
}

// WithCreatedBy tells the query-builder to eager-load the nodes that are connected to
// the "created_by" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithCreatedBy(opts ...func(*UserQuery)) *CustomerQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCreatedBy = query
	return cq
}

// WithNotes tells the query-builder to eager-load the nodes that are connected to
// the "notes" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithNotes(opts ...func(*NoteQuery)) *CustomerQuery {
	query := (&NoteClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withNotes = query
	return cq
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
//	client.Customer.Query().
//		GroupBy(customer.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CustomerQuery) GroupBy(field string, fields ...string) *CustomerGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = customer.Label
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
//	client.Customer.Query().
//		Select(customer.FieldName).
//		Scan(ctx, &v)
func (cq *CustomerQuery) Select(fields ...string) *CustomerSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CustomerSelect{CustomerQuery: cq}
	sbuild.label = customer.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerSelect configured with the given aggregations.
func (cq *CustomerQuery) Aggregate(fns ...AggregateFunc) *CustomerSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CustomerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !customer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CustomerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Customer, error) {
	var (
		nodes       = []*Customer{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [6]bool{
			cq.withOrders != nil,
			cq.withBillingAddress != nil,
			cq.withAddresses != nil,
			cq.withPhone != nil,
			cq.withCreatedBy != nil,
			cq.withNotes != nil,
		}
	)
	if cq.withBillingAddress != nil || cq.withCreatedBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, customer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Customer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Customer{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withOrders; query != nil {
		if err := cq.loadOrders(ctx, query, nodes,
			func(n *Customer) { n.Edges.Orders = []*Order{} },
			func(n *Customer, e *Order) { n.Edges.Orders = append(n.Edges.Orders, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withBillingAddress; query != nil {
		if err := cq.loadBillingAddress(ctx, query, nodes, nil,
			func(n *Customer, e *Address) { n.Edges.BillingAddress = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAddresses; query != nil {
		if err := cq.loadAddresses(ctx, query, nodes,
			func(n *Customer) { n.Edges.Addresses = []*Address{} },
			func(n *Customer, e *Address) { n.Edges.Addresses = append(n.Edges.Addresses, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withPhone; query != nil {
		if err := cq.loadPhone(ctx, query, nodes,
			func(n *Customer) { n.Edges.Phone = []*Tel{} },
			func(n *Customer, e *Tel) { n.Edges.Phone = append(n.Edges.Phone, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCreatedBy; query != nil {
		if err := cq.loadCreatedBy(ctx, query, nodes, nil,
			func(n *Customer, e *User) { n.Edges.CreatedBy = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withNotes; query != nil {
		if err := cq.loadNotes(ctx, query, nodes,
			func(n *Customer) { n.Edges.Notes = []*Note{} },
			func(n *Customer, e *Note) { n.Edges.Notes = append(n.Edges.Notes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CustomerQuery) loadOrders(ctx context.Context, query *OrderQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Order)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Order(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.OrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.customer_orders
		if fk == nil {
			return fmt.Errorf(`foreign-key "customer_orders" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_orders" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CustomerQuery) loadBillingAddress(ctx context.Context, query *AddressQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Address)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Customer)
	for i := range nodes {
		if nodes[i].customer_billing_address == nil {
			continue
		}
		fk := *nodes[i].customer_billing_address
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(address.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_billing_address" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CustomerQuery) loadAddresses(ctx context.Context, query *AddressQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Address)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Address(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.AddressesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.customer_addresses
		if fk == nil {
			return fmt.Errorf(`foreign-key "customer_addresses" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_addresses" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CustomerQuery) loadPhone(ctx context.Context, query *TelQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Tel)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Tel(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.PhoneColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.customer_phone
		if fk == nil {
			return fmt.Errorf(`foreign-key "customer_phone" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_phone" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CustomerQuery) loadCreatedBy(ctx context.Context, query *UserQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Customer)
	for i := range nodes {
		if nodes[i].customer_created_by == nil {
			continue
		}
		fk := *nodes[i].customer_created_by
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CustomerQuery) loadNotes(ctx context.Context, query *NoteQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Note)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Note(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.NotesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.customer_notes
		if fk == nil {
			return fmt.Errorf(`foreign-key "customer_notes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_notes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CustomerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CustomerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for i := range fields {
			if fields[i] != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CustomerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(customer.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = customer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CustomerGroupBy is the group-by builder for Customer entities.
type CustomerGroupBy struct {
	selector
	build *CustomerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CustomerGroupBy) Aggregate(fns ...AggregateFunc) *CustomerGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CustomerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerQuery, *CustomerGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CustomerGroupBy) sqlScan(ctx context.Context, root *CustomerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerSelect is the builder for selecting fields of Customer entities.
type CustomerSelect struct {
	*CustomerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CustomerSelect) Aggregate(fns ...AggregateFunc) *CustomerSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CustomerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerQuery, *CustomerSelect](ctx, cs.CustomerQuery, cs, cs.inters, v)
}

func (cs *CustomerSelect) sqlScan(ctx context.Context, root *CustomerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
