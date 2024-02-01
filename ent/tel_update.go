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
	"github.com/derfinlay/basecrm/ent/predicate"
	"github.com/derfinlay/basecrm/ent/tel"
)

// TelUpdate is the builder for updating Tel entities.
type TelUpdate struct {
	config
	hooks    []Hook
	mutation *TelMutation
}

// Where appends a list predicates to the TelUpdate builder.
func (tu *TelUpdate) Where(ps ...predicate.Tel) *TelUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TelUpdate) SetUpdatedAt(t time.Time) *TelUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNotesID sets the "notes" edge to the Note entity by ID.
func (tu *TelUpdate) SetNotesID(id int) *TelUpdate {
	tu.mutation.SetNotesID(id)
	return tu
}

// SetNillableNotesID sets the "notes" edge to the Note entity by ID if the given value is not nil.
func (tu *TelUpdate) SetNillableNotesID(id *int) *TelUpdate {
	if id != nil {
		tu = tu.SetNotesID(*id)
	}
	return tu
}

// SetNotes sets the "notes" edge to the Note entity.
func (tu *TelUpdate) SetNotes(n *Note) *TelUpdate {
	return tu.SetNotesID(n.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (tu *TelUpdate) SetCustomerID(id int) *TelUpdate {
	tu.mutation.SetCustomerID(id)
	return tu
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (tu *TelUpdate) SetNillableCustomerID(id *int) *TelUpdate {
	if id != nil {
		tu = tu.SetCustomerID(*id)
	}
	return tu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (tu *TelUpdate) SetCustomer(c *Customer) *TelUpdate {
	return tu.SetCustomerID(c.ID)
}

// Mutation returns the TelMutation object of the builder.
func (tu *TelUpdate) Mutation() *TelMutation {
	return tu.mutation
}

// ClearNotes clears the "notes" edge to the Note entity.
func (tu *TelUpdate) ClearNotes() *TelUpdate {
	tu.mutation.ClearNotes()
	return tu
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (tu *TelUpdate) ClearCustomer() *TelUpdate {
	tu.mutation.ClearCustomer()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TelUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TelUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TelUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TelUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TelUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := tel.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tel.Table, tel.Columns, sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tel.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tel.NotesTable,
			Columns: []string{tel.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tel.NotesTable,
			Columns: []string{tel.NotesColumn},
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
	if tu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tel.CustomerTable,
			Columns: []string{tel.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tel.CustomerTable,
			Columns: []string{tel.CustomerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TelUpdateOne is the builder for updating a single Tel entity.
type TelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TelMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TelUpdateOne) SetUpdatedAt(t time.Time) *TelUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNotesID sets the "notes" edge to the Note entity by ID.
func (tuo *TelUpdateOne) SetNotesID(id int) *TelUpdateOne {
	tuo.mutation.SetNotesID(id)
	return tuo
}

// SetNillableNotesID sets the "notes" edge to the Note entity by ID if the given value is not nil.
func (tuo *TelUpdateOne) SetNillableNotesID(id *int) *TelUpdateOne {
	if id != nil {
		tuo = tuo.SetNotesID(*id)
	}
	return tuo
}

// SetNotes sets the "notes" edge to the Note entity.
func (tuo *TelUpdateOne) SetNotes(n *Note) *TelUpdateOne {
	return tuo.SetNotesID(n.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (tuo *TelUpdateOne) SetCustomerID(id int) *TelUpdateOne {
	tuo.mutation.SetCustomerID(id)
	return tuo
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (tuo *TelUpdateOne) SetNillableCustomerID(id *int) *TelUpdateOne {
	if id != nil {
		tuo = tuo.SetCustomerID(*id)
	}
	return tuo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (tuo *TelUpdateOne) SetCustomer(c *Customer) *TelUpdateOne {
	return tuo.SetCustomerID(c.ID)
}

// Mutation returns the TelMutation object of the builder.
func (tuo *TelUpdateOne) Mutation() *TelMutation {
	return tuo.mutation
}

// ClearNotes clears the "notes" edge to the Note entity.
func (tuo *TelUpdateOne) ClearNotes() *TelUpdateOne {
	tuo.mutation.ClearNotes()
	return tuo
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (tuo *TelUpdateOne) ClearCustomer() *TelUpdateOne {
	tuo.mutation.ClearCustomer()
	return tuo
}

// Where appends a list predicates to the TelUpdate builder.
func (tuo *TelUpdateOne) Where(ps ...predicate.Tel) *TelUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TelUpdateOne) Select(field string, fields ...string) *TelUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tel entity.
func (tuo *TelUpdateOne) Save(ctx context.Context) (*Tel, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TelUpdateOne) SaveX(ctx context.Context) *Tel {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TelUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TelUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TelUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := tel.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TelUpdateOne) sqlSave(ctx context.Context) (_node *Tel, err error) {
	_spec := sqlgraph.NewUpdateSpec(tel.Table, tel.Columns, sqlgraph.NewFieldSpec(tel.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tel.FieldID)
		for _, f := range fields {
			if !tel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tel.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tel.NotesTable,
			Columns: []string{tel.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tel.NotesTable,
			Columns: []string{tel.NotesColumn},
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
	if tuo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tel.CustomerTable,
			Columns: []string{tel.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tel.CustomerTable,
			Columns: []string{tel.CustomerColumn},
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
	_node = &Tel{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
