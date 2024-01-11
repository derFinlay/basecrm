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
	"github.com/derfinlay/basecrm/ent/deliveryaddress"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/predicate"
	"github.com/derfinlay/basecrm/ent/tel"
)

// DeliveryAddressUpdate is the builder for updating DeliveryAddress entities.
type DeliveryAddressUpdate struct {
	config
	hooks    []Hook
	mutation *DeliveryAddressMutation
}

// Where appends a list predicates to the DeliveryAddressUpdate builder.
func (dau *DeliveryAddressUpdate) Where(ps ...predicate.DeliveryAddress) *DeliveryAddressUpdate {
	dau.mutation.Where(ps...)
	return dau
}

// SetNumber sets the "number" field.
func (dau *DeliveryAddressUpdate) SetNumber(s string) *DeliveryAddressUpdate {
	dau.mutation.SetNumber(s)
	return dau
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (dau *DeliveryAddressUpdate) SetNillableNumber(s *string) *DeliveryAddressUpdate {
	if s != nil {
		dau.SetNumber(*s)
	}
	return dau
}

// SetUpdatedAt sets the "updated_at" field.
func (dau *DeliveryAddressUpdate) SetUpdatedAt(t time.Time) *DeliveryAddressUpdate {
	dau.mutation.SetUpdatedAt(t)
	return dau
}

// SetTelephoneID sets the "telephone" edge to the Tel entity by ID.
func (dau *DeliveryAddressUpdate) SetTelephoneID(id int) *DeliveryAddressUpdate {
	dau.mutation.SetTelephoneID(id)
	return dau
}

// SetNillableTelephoneID sets the "telephone" edge to the Tel entity by ID if the given value is not nil.
func (dau *DeliveryAddressUpdate) SetNillableTelephoneID(id *int) *DeliveryAddressUpdate {
	if id != nil {
		dau = dau.SetTelephoneID(*id)
	}
	return dau
}

// SetTelephone sets the "telephone" edge to the Tel entity.
func (dau *DeliveryAddressUpdate) SetTelephone(t *Tel) *DeliveryAddressUpdate {
	return dau.SetTelephoneID(t.ID)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (dau *DeliveryAddressUpdate) AddNoteIDs(ids ...int) *DeliveryAddressUpdate {
	dau.mutation.AddNoteIDs(ids...)
	return dau
}

// AddNotes adds the "notes" edges to the Note entity.
func (dau *DeliveryAddressUpdate) AddNotes(n ...*Note) *DeliveryAddressUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return dau.AddNoteIDs(ids...)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (dau *DeliveryAddressUpdate) SetCustomerID(id int) *DeliveryAddressUpdate {
	dau.mutation.SetCustomerID(id)
	return dau
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (dau *DeliveryAddressUpdate) SetNillableCustomerID(id *int) *DeliveryAddressUpdate {
	if id != nil {
		dau = dau.SetCustomerID(*id)
	}
	return dau
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (dau *DeliveryAddressUpdate) SetCustomer(c *Customer) *DeliveryAddressUpdate {
	return dau.SetCustomerID(c.ID)
}

// Mutation returns the DeliveryAddressMutation object of the builder.
func (dau *DeliveryAddressUpdate) Mutation() *DeliveryAddressMutation {
	return dau.mutation
}

// ClearTelephone clears the "telephone" edge to the Tel entity.
func (dau *DeliveryAddressUpdate) ClearTelephone() *DeliveryAddressUpdate {
	dau.mutation.ClearTelephone()
	return dau
}

// ClearNotes clears all "notes" edges to the Note entity.
func (dau *DeliveryAddressUpdate) ClearNotes() *DeliveryAddressUpdate {
	dau.mutation.ClearNotes()
	return dau
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (dau *DeliveryAddressUpdate) RemoveNoteIDs(ids ...int) *DeliveryAddressUpdate {
	dau.mutation.RemoveNoteIDs(ids...)
	return dau
}

// RemoveNotes removes "notes" edges to Note entities.
func (dau *DeliveryAddressUpdate) RemoveNotes(n ...*Note) *DeliveryAddressUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return dau.RemoveNoteIDs(ids...)
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (dau *DeliveryAddressUpdate) ClearCustomer() *DeliveryAddressUpdate {
	dau.mutation.ClearCustomer()
	return dau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dau *DeliveryAddressUpdate) Save(ctx context.Context) (int, error) {
	dau.defaults()
	return withHooks(ctx, dau.sqlSave, dau.mutation, dau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dau *DeliveryAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := dau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dau *DeliveryAddressUpdate) Exec(ctx context.Context) error {
	_, err := dau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dau *DeliveryAddressUpdate) ExecX(ctx context.Context) {
	if err := dau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dau *DeliveryAddressUpdate) defaults() {
	if _, ok := dau.mutation.UpdatedAt(); !ok {
		v := deliveryaddress.UpdateDefaultUpdatedAt()
		dau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dau *DeliveryAddressUpdate) check() error {
	if v, ok := dau.mutation.Number(); ok {
		if err := deliveryaddress.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "DeliveryAddress.number": %w`, err)}
		}
	}
	return nil
}

func (dau *DeliveryAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deliveryaddress.Table, deliveryaddress.Columns, sqlgraph.NewFieldSpec(deliveryaddress.FieldID, field.TypeInt))
	if ps := dau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dau.mutation.Number(); ok {
		_spec.SetField(deliveryaddress.FieldNumber, field.TypeString, value)
	}
	if value, ok := dau.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryaddress.FieldUpdatedAt, field.TypeTime, value)
	}
	if dau.mutation.TelephoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryaddress.TelephoneTable,
			Columns: []string{deliveryaddress.TelephoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dau.mutation.TelephoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryaddress.TelephoneTable,
			Columns: []string{deliveryaddress.TelephoneColumn},
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
	if dau.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deliveryaddress.NotesTable,
			Columns: []string{deliveryaddress.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dau.mutation.RemovedNotesIDs(); len(nodes) > 0 && !dau.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deliveryaddress.NotesTable,
			Columns: []string{deliveryaddress.NotesColumn},
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
	if nodes := dau.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deliveryaddress.NotesTable,
			Columns: []string{deliveryaddress.NotesColumn},
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
	if dau.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deliveryaddress.CustomerTable,
			Columns: []string{deliveryaddress.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dau.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deliveryaddress.CustomerTable,
			Columns: []string{deliveryaddress.CustomerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, dau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deliveryaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dau.mutation.done = true
	return n, nil
}

// DeliveryAddressUpdateOne is the builder for updating a single DeliveryAddress entity.
type DeliveryAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeliveryAddressMutation
}

// SetNumber sets the "number" field.
func (dauo *DeliveryAddressUpdateOne) SetNumber(s string) *DeliveryAddressUpdateOne {
	dauo.mutation.SetNumber(s)
	return dauo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (dauo *DeliveryAddressUpdateOne) SetNillableNumber(s *string) *DeliveryAddressUpdateOne {
	if s != nil {
		dauo.SetNumber(*s)
	}
	return dauo
}

// SetUpdatedAt sets the "updated_at" field.
func (dauo *DeliveryAddressUpdateOne) SetUpdatedAt(t time.Time) *DeliveryAddressUpdateOne {
	dauo.mutation.SetUpdatedAt(t)
	return dauo
}

// SetTelephoneID sets the "telephone" edge to the Tel entity by ID.
func (dauo *DeliveryAddressUpdateOne) SetTelephoneID(id int) *DeliveryAddressUpdateOne {
	dauo.mutation.SetTelephoneID(id)
	return dauo
}

// SetNillableTelephoneID sets the "telephone" edge to the Tel entity by ID if the given value is not nil.
func (dauo *DeliveryAddressUpdateOne) SetNillableTelephoneID(id *int) *DeliveryAddressUpdateOne {
	if id != nil {
		dauo = dauo.SetTelephoneID(*id)
	}
	return dauo
}

// SetTelephone sets the "telephone" edge to the Tel entity.
func (dauo *DeliveryAddressUpdateOne) SetTelephone(t *Tel) *DeliveryAddressUpdateOne {
	return dauo.SetTelephoneID(t.ID)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (dauo *DeliveryAddressUpdateOne) AddNoteIDs(ids ...int) *DeliveryAddressUpdateOne {
	dauo.mutation.AddNoteIDs(ids...)
	return dauo
}

// AddNotes adds the "notes" edges to the Note entity.
func (dauo *DeliveryAddressUpdateOne) AddNotes(n ...*Note) *DeliveryAddressUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return dauo.AddNoteIDs(ids...)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (dauo *DeliveryAddressUpdateOne) SetCustomerID(id int) *DeliveryAddressUpdateOne {
	dauo.mutation.SetCustomerID(id)
	return dauo
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (dauo *DeliveryAddressUpdateOne) SetNillableCustomerID(id *int) *DeliveryAddressUpdateOne {
	if id != nil {
		dauo = dauo.SetCustomerID(*id)
	}
	return dauo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (dauo *DeliveryAddressUpdateOne) SetCustomer(c *Customer) *DeliveryAddressUpdateOne {
	return dauo.SetCustomerID(c.ID)
}

// Mutation returns the DeliveryAddressMutation object of the builder.
func (dauo *DeliveryAddressUpdateOne) Mutation() *DeliveryAddressMutation {
	return dauo.mutation
}

// ClearTelephone clears the "telephone" edge to the Tel entity.
func (dauo *DeliveryAddressUpdateOne) ClearTelephone() *DeliveryAddressUpdateOne {
	dauo.mutation.ClearTelephone()
	return dauo
}

// ClearNotes clears all "notes" edges to the Note entity.
func (dauo *DeliveryAddressUpdateOne) ClearNotes() *DeliveryAddressUpdateOne {
	dauo.mutation.ClearNotes()
	return dauo
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (dauo *DeliveryAddressUpdateOne) RemoveNoteIDs(ids ...int) *DeliveryAddressUpdateOne {
	dauo.mutation.RemoveNoteIDs(ids...)
	return dauo
}

// RemoveNotes removes "notes" edges to Note entities.
func (dauo *DeliveryAddressUpdateOne) RemoveNotes(n ...*Note) *DeliveryAddressUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return dauo.RemoveNoteIDs(ids...)
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (dauo *DeliveryAddressUpdateOne) ClearCustomer() *DeliveryAddressUpdateOne {
	dauo.mutation.ClearCustomer()
	return dauo
}

// Where appends a list predicates to the DeliveryAddressUpdate builder.
func (dauo *DeliveryAddressUpdateOne) Where(ps ...predicate.DeliveryAddress) *DeliveryAddressUpdateOne {
	dauo.mutation.Where(ps...)
	return dauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dauo *DeliveryAddressUpdateOne) Select(field string, fields ...string) *DeliveryAddressUpdateOne {
	dauo.fields = append([]string{field}, fields...)
	return dauo
}

// Save executes the query and returns the updated DeliveryAddress entity.
func (dauo *DeliveryAddressUpdateOne) Save(ctx context.Context) (*DeliveryAddress, error) {
	dauo.defaults()
	return withHooks(ctx, dauo.sqlSave, dauo.mutation, dauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dauo *DeliveryAddressUpdateOne) SaveX(ctx context.Context) *DeliveryAddress {
	node, err := dauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dauo *DeliveryAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := dauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dauo *DeliveryAddressUpdateOne) ExecX(ctx context.Context) {
	if err := dauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dauo *DeliveryAddressUpdateOne) defaults() {
	if _, ok := dauo.mutation.UpdatedAt(); !ok {
		v := deliveryaddress.UpdateDefaultUpdatedAt()
		dauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dauo *DeliveryAddressUpdateOne) check() error {
	if v, ok := dauo.mutation.Number(); ok {
		if err := deliveryaddress.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "DeliveryAddress.number": %w`, err)}
		}
	}
	return nil
}

func (dauo *DeliveryAddressUpdateOne) sqlSave(ctx context.Context) (_node *DeliveryAddress, err error) {
	if err := dauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deliveryaddress.Table, deliveryaddress.Columns, sqlgraph.NewFieldSpec(deliveryaddress.FieldID, field.TypeInt))
	id, ok := dauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeliveryAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deliveryaddress.FieldID)
		for _, f := range fields {
			if !deliveryaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deliveryaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dauo.mutation.Number(); ok {
		_spec.SetField(deliveryaddress.FieldNumber, field.TypeString, value)
	}
	if value, ok := dauo.mutation.UpdatedAt(); ok {
		_spec.SetField(deliveryaddress.FieldUpdatedAt, field.TypeTime, value)
	}
	if dauo.mutation.TelephoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryaddress.TelephoneTable,
			Columns: []string{deliveryaddress.TelephoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dauo.mutation.TelephoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deliveryaddress.TelephoneTable,
			Columns: []string{deliveryaddress.TelephoneColumn},
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
	if dauo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deliveryaddress.NotesTable,
			Columns: []string{deliveryaddress.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dauo.mutation.RemovedNotesIDs(); len(nodes) > 0 && !dauo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deliveryaddress.NotesTable,
			Columns: []string{deliveryaddress.NotesColumn},
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
	if nodes := dauo.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deliveryaddress.NotesTable,
			Columns: []string{deliveryaddress.NotesColumn},
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
	if dauo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deliveryaddress.CustomerTable,
			Columns: []string{deliveryaddress.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dauo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deliveryaddress.CustomerTable,
			Columns: []string{deliveryaddress.CustomerColumn},
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
	_node = &DeliveryAddress{config: dauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deliveryaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dauo.mutation.done = true
	return _node, nil
}
