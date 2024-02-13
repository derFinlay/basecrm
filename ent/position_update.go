// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/position"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks    []Hook
	mutation *PositionMutation
}

// Where appends a list predicates to the PositionUpdate builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (pu *PositionUpdate) AddNoteIDs(ids ...int) *PositionUpdate {
	pu.mutation.AddNoteIDs(ids...)
	return pu
}

// AddNotes adds the "notes" edges to the Note entity.
func (pu *PositionUpdate) AddNotes(n ...*Note) *PositionUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pu.AddNoteIDs(ids...)
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (pu *PositionUpdate) SetOrderID(id int) *PositionUpdate {
	pu.mutation.SetOrderID(id)
	return pu
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (pu *PositionUpdate) SetNillableOrderID(id *int) *PositionUpdate {
	if id != nil {
		pu = pu.SetOrderID(*id)
	}
	return pu
}

// SetOrder sets the "order" edge to the Order entity.
func (pu *PositionUpdate) SetOrder(o *Order) *PositionUpdate {
	return pu.SetOrderID(o.ID)
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// ClearNotes clears all "notes" edges to the Note entity.
func (pu *PositionUpdate) ClearNotes() *PositionUpdate {
	pu.mutation.ClearNotes()
	return pu
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (pu *PositionUpdate) RemoveNoteIDs(ids ...int) *PositionUpdate {
	pu.mutation.RemoveNoteIDs(ids...)
	return pu
}

// RemoveNotes removes "notes" edges to Note entities.
func (pu *PositionUpdate) RemoveNotes(n ...*Note) *PositionUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pu.RemoveNoteIDs(ids...)
}

// ClearOrder clears the "order" edge to the Order entity.
func (pu *PositionUpdate) ClearOrder() *PositionUpdate {
	pu.mutation.ClearOrder()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(position.FieldDescription, field.TypeString)
	}
	if pu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.NotesTable,
			Columns: []string{position.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedNotesIDs(); len(nodes) > 0 && !pu.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.NotesTable,
			Columns: []string{position.NotesColumn},
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
	if nodes := pu.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.NotesTable,
			Columns: []string{position.NotesColumn},
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
	if pu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.OrderTable,
			Columns: []string{position.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.OrderTable,
			Columns: []string{position.OrderColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PositionMutation
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (puo *PositionUpdateOne) AddNoteIDs(ids ...int) *PositionUpdateOne {
	puo.mutation.AddNoteIDs(ids...)
	return puo
}

// AddNotes adds the "notes" edges to the Note entity.
func (puo *PositionUpdateOne) AddNotes(n ...*Note) *PositionUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return puo.AddNoteIDs(ids...)
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (puo *PositionUpdateOne) SetOrderID(id int) *PositionUpdateOne {
	puo.mutation.SetOrderID(id)
	return puo
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableOrderID(id *int) *PositionUpdateOne {
	if id != nil {
		puo = puo.SetOrderID(*id)
	}
	return puo
}

// SetOrder sets the "order" edge to the Order entity.
func (puo *PositionUpdateOne) SetOrder(o *Order) *PositionUpdateOne {
	return puo.SetOrderID(o.ID)
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// ClearNotes clears all "notes" edges to the Note entity.
func (puo *PositionUpdateOne) ClearNotes() *PositionUpdateOne {
	puo.mutation.ClearNotes()
	return puo
}

// RemoveNoteIDs removes the "notes" edge to Note entities by IDs.
func (puo *PositionUpdateOne) RemoveNoteIDs(ids ...int) *PositionUpdateOne {
	puo.mutation.RemoveNoteIDs(ids...)
	return puo
}

// RemoveNotes removes "notes" edges to Note entities.
func (puo *PositionUpdateOne) RemoveNotes(n ...*Note) *PositionUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return puo.RemoveNoteIDs(ids...)
}

// ClearOrder clears the "order" edge to the Order entity.
func (puo *PositionUpdateOne) ClearOrder() *PositionUpdateOne {
	puo.mutation.ClearOrder()
	return puo
}

// Where appends a list predicates to the PositionUpdate builder.
func (puo *PositionUpdateOne) Where(ps ...predicate.Position) *PositionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PositionUpdateOne) Select(field string, fields ...string) *PositionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Position entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (_node *Position, err error) {
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Position.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for _, f := range fields {
			if !position.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(position.FieldDescription, field.TypeString)
	}
	if puo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.NotesTable,
			Columns: []string{position.NotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedNotesIDs(); len(nodes) > 0 && !puo.mutation.NotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.NotesTable,
			Columns: []string{position.NotesColumn},
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
	if nodes := puo.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.NotesTable,
			Columns: []string{position.NotesColumn},
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
	if puo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.OrderTable,
			Columns: []string{position.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.OrderTable,
			Columns: []string{position.OrderColumn},
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
	_node = &Position{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
