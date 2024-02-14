// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/predicate"
	"github.com/derfinlay/basecrm/ent/user"
	"github.com/derfinlay/basecrm/ent/usersession"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetLastLogin sets the "last_login" field.
func (uu *UserUpdate) SetLastLogin(t time.Time) *UserUpdate {
	uu.mutation.SetLastLogin(t)
	return uu
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastLogin(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastLogin(*t)
	}
	return uu
}

// ClearLastLogin clears the value of the "last_login" field.
func (uu *UserUpdate) ClearLastLogin() *UserUpdate {
	uu.mutation.ClearLastLogin()
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddCustomerIDs adds the "customers" edge to the Customer entity by IDs.
func (uu *UserUpdate) AddCustomerIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCustomerIDs(ids...)
	return uu
}

// AddCustomers adds the "customers" edges to the Customer entity.
func (uu *UserUpdate) AddCustomers(c ...*Customer) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCustomerIDs(ids...)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (uu *UserUpdate) AddNoteIDs(ids ...int) *UserUpdate {
	uu.mutation.AddNoteIDs(ids...)
	return uu
}

// AddNotes adds the "notes" edges to the Note entity.
func (uu *UserUpdate) AddNotes(n ...*Note) *UserUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uu.AddNoteIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (uu *UserUpdate) AddOrderIDs(ids ...int) *UserUpdate {
	uu.mutation.AddOrderIDs(ids...)
	return uu
}

// AddOrders adds the "orders" edges to the Order entity.
func (uu *UserUpdate) AddOrders(o ...*Order) *UserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.AddOrderIDs(ids...)
}

// AddSessionIDs adds the "sessions" edge to the UserSession entity by IDs.
func (uu *UserUpdate) AddSessionIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSessionIDs(ids...)
	return uu
}

// AddSessions adds the "sessions" edges to the UserSession entity.
func (uu *UserUpdate) AddSessions(u ...*UserSession) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddSessionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCustomers clears all "customers" edges to the Customer entity.
func (uu *UserUpdate) ClearCustomers() *UserUpdate {
	uu.mutation.ClearCustomers()
	return uu
}

// RemoveCustomerIDs removes the "customers" edge to Customer entities by IDs.
func (uu *UserUpdate) RemoveCustomerIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCustomerIDs(ids...)
	return uu
}

// RemoveCustomers removes "customers" edges to Customer entities.
func (uu *UserUpdate) RemoveCustomers(c ...*Customer) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCustomerIDs(ids...)
}

// ClearNotes clears all "notes" edges to the Note entity.
func (uu *UserUpdate) ClearNotes() *UserUpdate {
	uu.mutation.ClearNotes()
	return uu
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (uu *UserUpdate) RemoveNoteIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveNoteIDs(ids...)
	return uu
}

// RemoveNotes removes "notes" edges to Note entities.
func (uu *UserUpdate) RemoveNotes(n ...*Note) *UserUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uu.RemoveNoteIDs(ids...)
}

// ClearOrders clears all "orders" edges to the Order entity.
func (uu *UserUpdate) ClearOrders() *UserUpdate {
	uu.mutation.ClearOrders()
	return uu
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (uu *UserUpdate) RemoveOrderIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveOrderIDs(ids...)
	return uu
}

// RemoveOrders removes "orders" edges to Order entities.
func (uu *UserUpdate) RemoveOrders(o ...*Order) *UserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.RemoveOrderIDs(ids...)
}

// ClearSessions clears all "sessions" edges to the UserSession entity.
func (uu *UserUpdate) ClearSessions() *UserUpdate {
	uu.mutation.ClearSessions()
	return uu
}

// RemoveSessionIDs removes the "sessions" edge to UserSession entities by IDs.
func (uu *UserUpdate) RemoveSessionIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSessionIDs(ids...)
	return uu
}

// RemoveSessions removes "sessions" edges to UserSession entities.
func (uu *UserUpdate) RemoveSessions(u ...*UserSession) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if uu.mutation.LastLoginCleared() {
		_spec.ClearField(user.FieldLastLogin, field.TypeTime)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CustomersTable,
			Columns: []string{user.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCustomersIDs(); len(nodes) > 0 && !uu.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CustomersTable,
			Columns: []string{user.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CustomersTable,
			Columns: []string{user.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedNotesIDs(); len(nodes) > 0 && !uu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !uu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !uu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetLastLogin sets the "last_login" field.
func (uuo *UserUpdateOne) SetLastLogin(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLogin(t)
	return uuo
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastLogin(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastLogin(*t)
	}
	return uuo
}

// ClearLastLogin clears the value of the "last_login" field.
func (uuo *UserUpdateOne) ClearLastLogin() *UserUpdateOne {
	uuo.mutation.ClearLastLogin()
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddCustomerIDs adds the "customers" edge to the Customer entity by IDs.
func (uuo *UserUpdateOne) AddCustomerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCustomerIDs(ids...)
	return uuo
}

// AddCustomers adds the "customers" edges to the Customer entity.
func (uuo *UserUpdateOne) AddCustomers(c ...*Customer) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCustomerIDs(ids...)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (uuo *UserUpdateOne) AddNoteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddNoteIDs(ids...)
	return uuo
}

// AddNotes adds the "notes" edges to the Note entity.
func (uuo *UserUpdateOne) AddNotes(n ...*Note) *UserUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uuo.AddNoteIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (uuo *UserUpdateOne) AddOrderIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddOrderIDs(ids...)
	return uuo
}

// AddOrders adds the "orders" edges to the Order entity.
func (uuo *UserUpdateOne) AddOrders(o ...*Order) *UserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.AddOrderIDs(ids...)
}

// AddSessionIDs adds the "sessions" edge to the UserSession entity by IDs.
func (uuo *UserUpdateOne) AddSessionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSessionIDs(ids...)
	return uuo
}

// AddSessions adds the "sessions" edges to the UserSession entity.
func (uuo *UserUpdateOne) AddSessions(u ...*UserSession) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddSessionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCustomers clears all "customers" edges to the Customer entity.
func (uuo *UserUpdateOne) ClearCustomers() *UserUpdateOne {
	uuo.mutation.ClearCustomers()
	return uuo
}

// RemoveCustomerIDs removes the "customers" edge to Customer entities by IDs.
func (uuo *UserUpdateOne) RemoveCustomerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCustomerIDs(ids...)
	return uuo
}

// RemoveCustomers removes "customers" edges to Customer entities.
func (uuo *UserUpdateOne) RemoveCustomers(c ...*Customer) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCustomerIDs(ids...)
}

// ClearNotes clears all "notes" edges to the Note entity.
func (uuo *UserUpdateOne) ClearNotes() *UserUpdateOne {
	uuo.mutation.ClearNotes()
	return uuo
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (uuo *UserUpdateOne) RemoveNoteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveNoteIDs(ids...)
	return uuo
}

// RemoveNotes removes "notes" edges to Note entities.
func (uuo *UserUpdateOne) RemoveNotes(n ...*Note) *UserUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uuo.RemoveNoteIDs(ids...)
}

// ClearOrders clears all "orders" edges to the Order entity.
func (uuo *UserUpdateOne) ClearOrders() *UserUpdateOne {
	uuo.mutation.ClearOrders()
	return uuo
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (uuo *UserUpdateOne) RemoveOrderIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveOrderIDs(ids...)
	return uuo
}

// RemoveOrders removes "orders" edges to Order entities.
func (uuo *UserUpdateOne) RemoveOrders(o ...*Order) *UserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.RemoveOrderIDs(ids...)
}

// ClearSessions clears all "sessions" edges to the UserSession entity.
func (uuo *UserUpdateOne) ClearSessions() *UserUpdateOne {
	uuo.mutation.ClearSessions()
	return uuo
}

// RemoveSessionIDs removes the "sessions" edge to UserSession entities by IDs.
func (uuo *UserUpdateOne) RemoveSessionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSessionIDs(ids...)
	return uuo
}

// RemoveSessions removes "sessions" edges to UserSession entities.
func (uuo *UserUpdateOne) RemoveSessions(u ...*UserSession) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveSessionIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if uuo.mutation.LastLoginCleared() {
		_spec.ClearField(user.FieldLastLogin, field.TypeTime)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CustomersTable,
			Columns: []string{user.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCustomersIDs(); len(nodes) > 0 && !uuo.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CustomersTable,
			Columns: []string{user.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.CustomersTable,
			Columns: []string{user.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedNotesIDs(); len(nodes) > 0 && !uuo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.NotesTable,
			Columns: []string{user.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !uuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !uuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
