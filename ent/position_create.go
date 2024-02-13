// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/note"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/position"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PositionCreate) SetName(s string) *PositionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PositionCreate) SetDescription(s string) *PositionCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PositionCreate) SetNillableDescription(s *string) *PositionCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetUnitPrice sets the "unit_price" field.
func (pc *PositionCreate) SetUnitPrice(f float64) *PositionCreate {
	pc.mutation.SetUnitPrice(f)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PositionCreate) SetAmount(i int) *PositionCreate {
	pc.mutation.SetAmount(i)
	return pc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pc *PositionCreate) SetNillableAmount(i *int) *PositionCreate {
	if i != nil {
		pc.SetAmount(*i)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PositionCreate) SetCreatedAt(t time.Time) *PositionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreatedAt(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (pc *PositionCreate) AddNoteIDs(ids ...int) *PositionCreate {
	pc.mutation.AddNoteIDs(ids...)
	return pc
}

// AddNotes adds the "notes" edges to the Note entity.
func (pc *PositionCreate) AddNotes(n ...*Note) *PositionCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pc.AddNoteIDs(ids...)
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (pc *PositionCreate) SetOrderID(id int) *PositionCreate {
	pc.mutation.SetOrderID(id)
	return pc
}

// SetNillableOrderID sets the "order" edge to the Order entity by ID if the given value is not nil.
func (pc *PositionCreate) SetNillableOrderID(id *int) *PositionCreate {
	if id != nil {
		pc = pc.SetOrderID(*id)
	}
	return pc
}

// SetOrder sets the "order" edge to the Order entity.
func (pc *PositionCreate) SetOrder(o *Order) *PositionCreate {
	return pc.SetOrderID(o.ID)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PositionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PositionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PositionCreate) defaults() {
	if _, ok := pc.mutation.Amount(); !ok {
		v := position.DefaultAmount
		pc.mutation.SetAmount(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := position.DefaultCreatedAt
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PositionCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Position.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := position.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Position.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.UnitPrice(); !ok {
		return &ValidationError{Name: "unit_price", err: errors.New(`ent: missing required field "Position.unit_price"`)}
	}
	if v, ok := pc.mutation.UnitPrice(); ok {
		if err := position.UnitPriceValidator(v); err != nil {
			return &ValidationError{Name: "unit_price", err: fmt.Errorf(`ent: validator failed for field "Position.unit_price": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Position.amount"`)}
	}
	if v, ok := pc.mutation.Amount(); ok {
		if err := position.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Position.amount": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Position.created_at"`)}
	}
	return nil
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		_node = &Position{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(position.Table, sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(position.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.UnitPrice(); ok {
		_spec.SetField(position.FieldUnitPrice, field.TypeFloat64, value)
		_node.UnitPrice = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(position.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(position.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := pc.mutation.NotesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OrderIDs(); len(nodes) > 0 {
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
		_node.order_positions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PositionCreateBulk is the builder for creating many Position entities in bulk.
type PositionCreateBulk struct {
	config
	err      error
	builders []*PositionCreate
}

// Save creates the Position entities in the database.
func (pcb *PositionCreateBulk) Save(ctx context.Context) ([]*Position, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Position, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PositionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PositionCreateBulk) SaveX(ctx context.Context) []*Position {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PositionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PositionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
