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
	"github.com/derfinlay/basecrm/ent/address"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/predicate"
	"github.com/derfinlay/basecrm/ent/tel"
	"github.com/derfinlay/basecrm/ent/user"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetEmail sets the "email" field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableEmail(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetEmail(*s)
	}
	return cu
}

// SetPassword sets the "password" field.
func (cu *CustomerUpdate) SetPassword(s string) *CustomerUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillablePassword(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetPassword(*s)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CustomerUpdate) SetUpdatedAt(t time.Time) *CustomerUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (cu *CustomerUpdate) AddOrderIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddOrderIDs(ids...)
	return cu
}

// AddOrders adds the "orders" edges to the Order entity.
func (cu *CustomerUpdate) AddOrders(o ...*Order) *CustomerUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.AddOrderIDs(ids...)
}

// SetBillingAddressID sets the "billing_address" edge to the Address entity by ID.
func (cu *CustomerUpdate) SetBillingAddressID(id int) *CustomerUpdate {
	cu.mutation.SetBillingAddressID(id)
	return cu
}

// SetNillableBillingAddressID sets the "billing_address" edge to the Address entity by ID if the given value is not nil.
func (cu *CustomerUpdate) SetNillableBillingAddressID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetBillingAddressID(*id)
	}
	return cu
}

// SetBillingAddress sets the "billing_address" edge to the Address entity.
func (cu *CustomerUpdate) SetBillingAddress(a *Address) *CustomerUpdate {
	return cu.SetBillingAddressID(a.ID)
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (cu *CustomerUpdate) AddAddressIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddAddressIDs(ids...)
	return cu
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (cu *CustomerUpdate) AddAddresses(a ...*Address) *CustomerUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddAddressIDs(ids...)
}

// AddPhoneIDs adds the "phone" edge to the Tel entity by IDs.
func (cu *CustomerUpdate) AddPhoneIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddPhoneIDs(ids...)
	return cu
}

// AddPhone adds the "phone" edges to the Tel entity.
func (cu *CustomerUpdate) AddPhone(t ...*Tel) *CustomerUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddPhoneIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (cu *CustomerUpdate) SetCreatedByID(id int) *CustomerUpdate {
	cu.mutation.SetCreatedByID(id)
	return cu
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (cu *CustomerUpdate) SetNillableCreatedByID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetCreatedByID(*id)
	}
	return cu
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (cu *CustomerUpdate) SetCreatedBy(u *User) *CustomerUpdate {
	return cu.SetCreatedByID(u.ID)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (cu *CustomerUpdate) AddNoteIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddNoteIDs(ids...)
	return cu
}

// AddNotes adds the "notes" edges to the Note entity.
func (cu *CustomerUpdate) AddNotes(n ...*Note) *CustomerUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cu.AddNoteIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearOrders clears all "orders" edges to the Order entity.
func (cu *CustomerUpdate) ClearOrders() *CustomerUpdate {
	cu.mutation.ClearOrders()
	return cu
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (cu *CustomerUpdate) RemoveOrderIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveOrderIDs(ids...)
	return cu
}

// RemoveOrders removes "orders" edges to Order entities.
func (cu *CustomerUpdate) RemoveOrders(o ...*Order) *CustomerUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.RemoveOrderIDs(ids...)
}

// ClearBillingAddress clears the "billing_address" edge to the Address entity.
func (cu *CustomerUpdate) ClearBillingAddress() *CustomerUpdate {
	cu.mutation.ClearBillingAddress()
	return cu
}

// ClearAddresses clears all "addresses" edges to the Address entity.
func (cu *CustomerUpdate) ClearAddresses() *CustomerUpdate {
	cu.mutation.ClearAddresses()
	return cu
}

// RemoveAddressIDs removes the "addresses" edge to Address entities by IDs.
func (cu *CustomerUpdate) RemoveAddressIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveAddressIDs(ids...)
	return cu
}

// RemoveAddresses removes "addresses" edges to Address entities.
func (cu *CustomerUpdate) RemoveAddresses(a ...*Address) *CustomerUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveAddressIDs(ids...)
}

// ClearPhone clears all "phone" edges to the Tel entity.
func (cu *CustomerUpdate) ClearPhone() *CustomerUpdate {
	cu.mutation.ClearPhone()
	return cu
}

// RemovePhoneIDs removes the "phone" edge to Tel entities by IDs.
func (cu *CustomerUpdate) RemovePhoneIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemovePhoneIDs(ids...)
	return cu
}

// RemovePhone removes "phone" edges to Tel entities.
func (cu *CustomerUpdate) RemovePhone(t ...*Tel) *CustomerUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemovePhoneIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (cu *CustomerUpdate) ClearCreatedBy() *CustomerUpdate {
	cu.mutation.ClearCreatedBy()
	return cu
}

// ClearNotes clears all "notes" edges to the Note entity.
func (cu *CustomerUpdate) ClearNotes() *CustomerUpdate {
	cu.mutation.ClearNotes()
	return cu
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (cu *CustomerUpdate) RemoveNoteIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveNoteIDs(ids...)
	return cu
}

// RemoveNotes removes "notes" edges to Note entities.
func (cu *CustomerUpdate) RemoveNotes(n ...*Note) *CustomerUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cu.RemoveNoteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CustomerUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := customer.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CustomerUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Email(); ok {
		if err := customer.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Customer.email": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Customer.password": %w`, err)}
		}
	}
	return nil
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !cu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
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
	if nodes := cu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
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
	if cu.mutation.BillingAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.BillingAddressTable,
			Columns: []string{customer.BillingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BillingAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.BillingAddressTable,
			Columns: []string{customer.BillingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedAddressesIDs(); len(nodes) > 0 && !cu.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PhoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PhoneTable,
			Columns: []string{customer.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPhoneIDs(); len(nodes) > 0 && !cu.mutation.PhoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PhoneTable,
			Columns: []string{customer.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PhoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PhoneTable,
			Columns: []string{customer.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.CreatedByTable,
			Columns: []string{customer.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.CreatedByTable,
			Columns: []string{customer.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.NotesTable,
			Columns: []string{customer.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedNotesIDs(); len(nodes) > 0 && !cu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.NotesTable,
			Columns: []string{customer.NotesColumn},
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
	if nodes := cu.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.NotesTable,
			Columns: []string{customer.NotesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomerMutation
}

// SetName sets the "name" field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetEmail sets the "email" field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableEmail(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetEmail(*s)
	}
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *CustomerUpdateOne) SetPassword(s string) *CustomerUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillablePassword(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetPassword(*s)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CustomerUpdateOne) SetUpdatedAt(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (cuo *CustomerUpdateOne) AddOrderIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddOrderIDs(ids...)
	return cuo
}

// AddOrders adds the "orders" edges to the Order entity.
func (cuo *CustomerUpdateOne) AddOrders(o ...*Order) *CustomerUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.AddOrderIDs(ids...)
}

// SetBillingAddressID sets the "billing_address" edge to the Address entity by ID.
func (cuo *CustomerUpdateOne) SetBillingAddressID(id int) *CustomerUpdateOne {
	cuo.mutation.SetBillingAddressID(id)
	return cuo
}

// SetNillableBillingAddressID sets the "billing_address" edge to the Address entity by ID if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableBillingAddressID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetBillingAddressID(*id)
	}
	return cuo
}

// SetBillingAddress sets the "billing_address" edge to the Address entity.
func (cuo *CustomerUpdateOne) SetBillingAddress(a *Address) *CustomerUpdateOne {
	return cuo.SetBillingAddressID(a.ID)
}

// AddAddressIDs adds the "addresses" edge to the Address entity by IDs.
func (cuo *CustomerUpdateOne) AddAddressIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddAddressIDs(ids...)
	return cuo
}

// AddAddresses adds the "addresses" edges to the Address entity.
func (cuo *CustomerUpdateOne) AddAddresses(a ...*Address) *CustomerUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddAddressIDs(ids...)
}

// AddPhoneIDs adds the "phone" edge to the Tel entity by IDs.
func (cuo *CustomerUpdateOne) AddPhoneIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddPhoneIDs(ids...)
	return cuo
}

// AddPhone adds the "phone" edges to the Tel entity.
func (cuo *CustomerUpdateOne) AddPhone(t ...*Tel) *CustomerUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddPhoneIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (cuo *CustomerUpdateOne) SetCreatedByID(id int) *CustomerUpdateOne {
	cuo.mutation.SetCreatedByID(id)
	return cuo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableCreatedByID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetCreatedByID(*id)
	}
	return cuo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (cuo *CustomerUpdateOne) SetCreatedBy(u *User) *CustomerUpdateOne {
	return cuo.SetCreatedByID(u.ID)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (cuo *CustomerUpdateOne) AddNoteIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddNoteIDs(ids...)
	return cuo
}

// AddNotes adds the "notes" edges to the Note entity.
func (cuo *CustomerUpdateOne) AddNotes(n ...*Note) *CustomerUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuo.AddNoteIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearOrders clears all "orders" edges to the Order entity.
func (cuo *CustomerUpdateOne) ClearOrders() *CustomerUpdateOne {
	cuo.mutation.ClearOrders()
	return cuo
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (cuo *CustomerUpdateOne) RemoveOrderIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveOrderIDs(ids...)
	return cuo
}

// RemoveOrders removes "orders" edges to Order entities.
func (cuo *CustomerUpdateOne) RemoveOrders(o ...*Order) *CustomerUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.RemoveOrderIDs(ids...)
}

// ClearBillingAddress clears the "billing_address" edge to the Address entity.
func (cuo *CustomerUpdateOne) ClearBillingAddress() *CustomerUpdateOne {
	cuo.mutation.ClearBillingAddress()
	return cuo
}

// ClearAddresses clears all "addresses" edges to the Address entity.
func (cuo *CustomerUpdateOne) ClearAddresses() *CustomerUpdateOne {
	cuo.mutation.ClearAddresses()
	return cuo
}

// RemoveAddressIDs removes the "addresses" edge to Address entities by IDs.
func (cuo *CustomerUpdateOne) RemoveAddressIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveAddressIDs(ids...)
	return cuo
}

// RemoveAddresses removes "addresses" edges to Address entities.
func (cuo *CustomerUpdateOne) RemoveAddresses(a ...*Address) *CustomerUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveAddressIDs(ids...)
}

// ClearPhone clears all "phone" edges to the Tel entity.
func (cuo *CustomerUpdateOne) ClearPhone() *CustomerUpdateOne {
	cuo.mutation.ClearPhone()
	return cuo
}

// RemovePhoneIDs removes the "phone" edge to Tel entities by IDs.
func (cuo *CustomerUpdateOne) RemovePhoneIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemovePhoneIDs(ids...)
	return cuo
}

// RemovePhone removes "phone" edges to Tel entities.
func (cuo *CustomerUpdateOne) RemovePhone(t ...*Tel) *CustomerUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemovePhoneIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (cuo *CustomerUpdateOne) ClearCreatedBy() *CustomerUpdateOne {
	cuo.mutation.ClearCreatedBy()
	return cuo
}

// ClearNotes clears all "notes" edges to the Note entity.
func (cuo *CustomerUpdateOne) ClearNotes() *CustomerUpdateOne {
	cuo.mutation.ClearNotes()
	return cuo
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (cuo *CustomerUpdateOne) RemoveNoteIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveNoteIDs(ids...)
	return cuo
}

// RemoveNotes removes "notes" edges to Note entities.
func (cuo *CustomerUpdateOne) RemoveNotes(n ...*Note) *CustomerUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuo.RemoveNoteIDs(ids...)
}

// Where appends a list predicates to the CustomerUpdate builder.
func (cuo *CustomerUpdateOne) Where(ps ...predicate.Customer) *CustomerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CustomerUpdateOne) Select(field string, fields ...string) *CustomerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CustomerUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := customer.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CustomerUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := customer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Customer.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Email(); ok {
		if err := customer.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Customer.email": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Customer.password": %w`, err)}
		}
	}
	return nil
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Customer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for _, f := range fields {
			if !customer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(customer.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.SetField(customer.FieldPassword, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !cuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
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
	if nodes := cuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrdersTable,
			Columns: []string{customer.OrdersColumn},
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
	if cuo.mutation.BillingAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.BillingAddressTable,
			Columns: []string{customer.BillingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BillingAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.BillingAddressTable,
			Columns: []string{customer.BillingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedAddressesIDs(); len(nodes) > 0 && !cuo.mutation.AddressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressesTable,
			Columns: []string{customer.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PhoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PhoneTable,
			Columns: []string{customer.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPhoneIDs(); len(nodes) > 0 && !cuo.mutation.PhoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PhoneTable,
			Columns: []string{customer.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PhoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.PhoneTable,
			Columns: []string{customer.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.CreatedByTable,
			Columns: []string{customer.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.CreatedByTable,
			Columns: []string{customer.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.NotesTable,
			Columns: []string{customer.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedNotesIDs(); len(nodes) > 0 && !cuo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.NotesTable,
			Columns: []string{customer.NotesColumn},
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
	if nodes := cuo.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.NotesTable,
			Columns: []string{customer.NotesColumn},
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
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}