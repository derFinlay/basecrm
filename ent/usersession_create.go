// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/user"
	"github.com/derfinlay/basecrm/ent/usersession"
)

// UserSessionCreate is the builder for creating a UserSession entity.
type UserSessionCreate struct {
	config
	mutation *UserSessionMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (usc *UserSessionCreate) SetToken(s string) *UserSessionCreate {
	usc.mutation.SetToken(s)
	return usc
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserSessionCreate) SetCreatedAt(t time.Time) *UserSessionCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserSessionCreate) SetNillableCreatedAt(t *time.Time) *UserSessionCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UserSessionCreate) SetUpdatedAt(t time.Time) *UserSessionCreate {
	usc.mutation.SetUpdatedAt(t)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UserSessionCreate) SetNillableUpdatedAt(t *time.Time) *UserSessionCreate {
	if t != nil {
		usc.SetUpdatedAt(*t)
	}
	return usc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usc *UserSessionCreate) SetUserID(id int) *UserSessionCreate {
	usc.mutation.SetUserID(id)
	return usc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usc *UserSessionCreate) SetNillableUserID(id *int) *UserSessionCreate {
	if id != nil {
		usc = usc.SetUserID(*id)
	}
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSessionCreate) SetUser(u *User) *UserSessionCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserSessionMutation object of the builder.
func (usc *UserSessionCreate) Mutation() *UserSessionMutation {
	return usc.mutation
}

// Save creates the UserSession in the database.
func (usc *UserSessionCreate) Save(ctx context.Context) (*UserSession, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSessionCreate) SaveX(ctx context.Context) *UserSession {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSessionCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSessionCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSessionCreate) defaults() {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		v := usersession.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		v := usersession.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSessionCreate) check() error {
	if _, ok := usc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "UserSession.token"`)}
	}
	if v, ok := usc.mutation.Token(); ok {
		if err := usersession.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "UserSession.token": %w`, err)}
		}
	}
	if _, ok := usc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserSession.created_at"`)}
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserSession.updated_at"`)}
	}
	return nil
}

func (usc *UserSessionCreate) sqlSave(ctx context.Context) (*UserSession, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSessionCreate) createSpec() (*UserSession, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSession{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersession.Table, sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt))
	)
	if value, ok := usc.mutation.Token(); ok {
		_spec.SetField(usersession.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.SetField(usersession.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.SetField(usersession.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   usersession.UserTable,
			Columns: []string{usersession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSessionCreateBulk is the builder for creating many UserSession entities in bulk.
type UserSessionCreateBulk struct {
	config
	err      error
	builders []*UserSessionCreate
}

// Save creates the UserSession entities in the database.
func (uscb *UserSessionCreateBulk) Save(ctx context.Context) ([]*UserSession, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSession, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSessionMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSessionCreateBulk) SaveX(ctx context.Context) []*UserSession {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSessionCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSessionCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
