// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/login"
	"github.com/derfinlay/basecrm/ent/loginreset"
	"github.com/derfinlay/basecrm/ent/note"
)

// LoginCreate is the builder for creating a Login entity.
type LoginCreate struct {
	config
	mutation *LoginMutation
	hooks    []Hook
}

// SetPassword sets the "password" field.
func (lc *LoginCreate) SetPassword(s string) *LoginCreate {
	lc.mutation.SetPassword(s)
	return lc
}

// SetEmail sets the "email" field.
func (lc *LoginCreate) SetEmail(s string) *LoginCreate {
	lc.mutation.SetEmail(s)
	return lc
}

// SetLastLogin sets the "last_login" field.
func (lc *LoginCreate) SetLastLogin(t time.Time) *LoginCreate {
	lc.mutation.SetLastLogin(t)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LoginCreate) SetCreatedAt(t time.Time) *LoginCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LoginCreate) SetNillableCreatedAt(t *time.Time) *LoginCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LoginCreate) SetUpdatedAt(t time.Time) *LoginCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LoginCreate) SetNillableUpdatedAt(t *time.Time) *LoginCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (lc *LoginCreate) SetCustomerID(id int) *LoginCreate {
	lc.mutation.SetCustomerID(id)
	return lc
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (lc *LoginCreate) SetNillableCustomerID(id *int) *LoginCreate {
	if id != nil {
		lc = lc.SetCustomerID(*id)
	}
	return lc
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (lc *LoginCreate) SetCustomer(c *Customer) *LoginCreate {
	return lc.SetCustomerID(c.ID)
}

// AddLoginResetIDs adds the "login_resets" edge to the LoginReset entity by IDs.
func (lc *LoginCreate) AddLoginResetIDs(ids ...int) *LoginCreate {
	lc.mutation.AddLoginResetIDs(ids...)
	return lc
}

// AddLoginResets adds the "login_resets" edges to the LoginReset entity.
func (lc *LoginCreate) AddLoginResets(l ...*LoginReset) *LoginCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lc.AddLoginResetIDs(ids...)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (lc *LoginCreate) AddNoteIDs(ids ...int) *LoginCreate {
	lc.mutation.AddNoteIDs(ids...)
	return lc
}

// AddNotes adds the "notes" edges to the Note entity.
func (lc *LoginCreate) AddNotes(n ...*Note) *LoginCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return lc.AddNoteIDs(ids...)
}

// Mutation returns the LoginMutation object of the builder.
func (lc *LoginCreate) Mutation() *LoginMutation {
	return lc.mutation
}

// Save creates the Login in the database.
func (lc *LoginCreate) Save(ctx context.Context) (*Login, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LoginCreate) SaveX(ctx context.Context) *Login {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LoginCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LoginCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LoginCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := login.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := login.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LoginCreate) check() error {
	if _, ok := lc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Login.password"`)}
	}
	if v, ok := lc.mutation.Password(); ok {
		if err := login.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Login.password": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Login.email"`)}
	}
	if v, ok := lc.mutation.Email(); ok {
		if err := login.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Login.email": %w`, err)}
		}
	}
	if _, ok := lc.mutation.LastLogin(); !ok {
		return &ValidationError{Name: "last_login", err: errors.New(`ent: missing required field "Login.last_login"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Login.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Login.updated_at"`)}
	}
	return nil
}

func (lc *LoginCreate) sqlSave(ctx context.Context) (*Login, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LoginCreate) createSpec() (*Login, *sqlgraph.CreateSpec) {
	var (
		_node = &Login{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(login.Table, sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.Password(); ok {
		_spec.SetField(login.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := lc.mutation.Email(); ok {
		_spec.SetField(login.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := lc.mutation.LastLogin(); ok {
		_spec.SetField(login.FieldLastLogin, field.TypeTime, value)
		_node.LastLogin = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(login.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(login.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := lc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   login.CustomerTable,
			Columns: []string{login.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_login = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.LoginResetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   login.LoginResetsTable,
			Columns: []string{login.LoginResetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginreset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   login.NotesTable,
			Columns: []string{login.NotesColumn},
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
	return _node, _spec
}

// LoginCreateBulk is the builder for creating many Login entities in bulk.
type LoginCreateBulk struct {
	config
	err      error
	builders []*LoginCreate
}

// Save creates the Login entities in the database.
func (lcb *LoginCreateBulk) Save(ctx context.Context) ([]*Login, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Login, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoginMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LoginCreateBulk) SaveX(ctx context.Context) []*Login {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LoginCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LoginCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
