// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/login"
	"github.com/derfinlay/basecrm/ent/loginreset"
	"github.com/derfinlay/basecrm/ent/note"
)

// LoginResetCreate is the builder for creating a LoginReset entity.
type LoginResetCreate struct {
	config
	mutation *LoginResetMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (lrc *LoginResetCreate) SetToken(s string) *LoginResetCreate {
	lrc.mutation.SetToken(s)
	return lrc
}

// SetActive sets the "active" field.
func (lrc *LoginResetCreate) SetActive(b bool) *LoginResetCreate {
	lrc.mutation.SetActive(b)
	return lrc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (lrc *LoginResetCreate) SetNillableActive(b *bool) *LoginResetCreate {
	if b != nil {
		lrc.SetActive(*b)
	}
	return lrc
}

// SetCreatedAt sets the "created_at" field.
func (lrc *LoginResetCreate) SetCreatedAt(t time.Time) *LoginResetCreate {
	lrc.mutation.SetCreatedAt(t)
	return lrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lrc *LoginResetCreate) SetNillableCreatedAt(t *time.Time) *LoginResetCreate {
	if t != nil {
		lrc.SetCreatedAt(*t)
	}
	return lrc
}

// SetUpdatedAt sets the "updated_at" field.
func (lrc *LoginResetCreate) SetUpdatedAt(t time.Time) *LoginResetCreate {
	lrc.mutation.SetUpdatedAt(t)
	return lrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lrc *LoginResetCreate) SetNillableUpdatedAt(t *time.Time) *LoginResetCreate {
	if t != nil {
		lrc.SetUpdatedAt(*t)
	}
	return lrc
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (lrc *LoginResetCreate) AddNoteIDs(ids ...int) *LoginResetCreate {
	lrc.mutation.AddNoteIDs(ids...)
	return lrc
}

// AddNotes adds the "notes" edges to the Note entity.
func (lrc *LoginResetCreate) AddNotes(n ...*Note) *LoginResetCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return lrc.AddNoteIDs(ids...)
}

// SetLoginID sets the "login" edge to the Login entity by ID.
func (lrc *LoginResetCreate) SetLoginID(id int) *LoginResetCreate {
	lrc.mutation.SetLoginID(id)
	return lrc
}

// SetNillableLoginID sets the "login" edge to the Login entity by ID if the given value is not nil.
func (lrc *LoginResetCreate) SetNillableLoginID(id *int) *LoginResetCreate {
	if id != nil {
		lrc = lrc.SetLoginID(*id)
	}
	return lrc
}

// SetLogin sets the "login" edge to the Login entity.
func (lrc *LoginResetCreate) SetLogin(l *Login) *LoginResetCreate {
	return lrc.SetLoginID(l.ID)
}

// Mutation returns the LoginResetMutation object of the builder.
func (lrc *LoginResetCreate) Mutation() *LoginResetMutation {
	return lrc.mutation
}

// Save creates the LoginReset in the database.
func (lrc *LoginResetCreate) Save(ctx context.Context) (*LoginReset, error) {
	lrc.defaults()
	return withHooks(ctx, lrc.sqlSave, lrc.mutation, lrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lrc *LoginResetCreate) SaveX(ctx context.Context) *LoginReset {
	v, err := lrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lrc *LoginResetCreate) Exec(ctx context.Context) error {
	_, err := lrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lrc *LoginResetCreate) ExecX(ctx context.Context) {
	if err := lrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lrc *LoginResetCreate) defaults() {
	if _, ok := lrc.mutation.Active(); !ok {
		v := loginreset.DefaultActive
		lrc.mutation.SetActive(v)
	}
	if _, ok := lrc.mutation.CreatedAt(); !ok {
		v := loginreset.DefaultCreatedAt()
		lrc.mutation.SetCreatedAt(v)
	}
	if _, ok := lrc.mutation.UpdatedAt(); !ok {
		v := loginreset.DefaultUpdatedAt()
		lrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lrc *LoginResetCreate) check() error {
	if _, ok := lrc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "LoginReset.token"`)}
	}
	if v, ok := lrc.mutation.Token(); ok {
		if err := loginreset.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "LoginReset.token": %w`, err)}
		}
	}
	if _, ok := lrc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "LoginReset.active"`)}
	}
	if _, ok := lrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LoginReset.created_at"`)}
	}
	if _, ok := lrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LoginReset.updated_at"`)}
	}
	return nil
}

func (lrc *LoginResetCreate) sqlSave(ctx context.Context) (*LoginReset, error) {
	if err := lrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lrc.mutation.id = &_node.ID
	lrc.mutation.done = true
	return _node, nil
}

func (lrc *LoginResetCreate) createSpec() (*LoginReset, *sqlgraph.CreateSpec) {
	var (
		_node = &LoginReset{config: lrc.config}
		_spec = sqlgraph.NewCreateSpec(loginreset.Table, sqlgraph.NewFieldSpec(loginreset.FieldID, field.TypeInt))
	)
	if value, ok := lrc.mutation.Token(); ok {
		_spec.SetField(loginreset.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := lrc.mutation.Active(); ok {
		_spec.SetField(loginreset.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := lrc.mutation.CreatedAt(); ok {
		_spec.SetField(loginreset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lrc.mutation.UpdatedAt(); ok {
		_spec.SetField(loginreset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := lrc.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   loginreset.NotesTable,
			Columns: []string{loginreset.NotesColumn},
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
	if nodes := lrc.mutation.LoginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   loginreset.LoginTable,
			Columns: []string{loginreset.LoginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.login_login_resets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LoginResetCreateBulk is the builder for creating many LoginReset entities in bulk.
type LoginResetCreateBulk struct {
	config
	err      error
	builders []*LoginResetCreate
}

// Save creates the LoginReset entities in the database.
func (lrcb *LoginResetCreateBulk) Save(ctx context.Context) ([]*LoginReset, error) {
	if lrcb.err != nil {
		return nil, lrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lrcb.builders))
	nodes := make([]*LoginReset, len(lrcb.builders))
	mutators := make([]Mutator, len(lrcb.builders))
	for i := range lrcb.builders {
		func(i int, root context.Context) {
			builder := lrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoginResetMutation)
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
					_, err = mutators[i+1].Mutate(root, lrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lrcb *LoginResetCreateBulk) SaveX(ctx context.Context) []*LoginReset {
	v, err := lrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lrcb *LoginResetCreateBulk) Exec(ctx context.Context) error {
	_, err := lrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lrcb *LoginResetCreateBulk) ExecX(ctx context.Context) {
	if err := lrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
