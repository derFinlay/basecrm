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
	"github.com/derfinlay/basecrm/ent/predicate"
	"github.com/derfinlay/basecrm/ent/user"
	"github.com/derfinlay/basecrm/ent/usersession"
)

// UserSessionUpdate is the builder for updating UserSession entities.
type UserSessionUpdate struct {
	config
	hooks    []Hook
	mutation *UserSessionMutation
}

// Where appends a list predicates to the UserSessionUpdate builder.
func (usu *UserSessionUpdate) Where(ps ...predicate.UserSession) *UserSessionUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetUpdatedAt sets the "updated_at" field.
func (usu *UserSessionUpdate) SetUpdatedAt(t time.Time) *UserSessionUpdate {
	usu.mutation.SetUpdatedAt(t)
	return usu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usu *UserSessionUpdate) SetUserID(id int) *UserSessionUpdate {
	usu.mutation.SetUserID(id)
	return usu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usu *UserSessionUpdate) SetNillableUserID(id *int) *UserSessionUpdate {
	if id != nil {
		usu = usu.SetUserID(*id)
	}
	return usu
}

// SetUser sets the "user" edge to the User entity.
func (usu *UserSessionUpdate) SetUser(u *User) *UserSessionUpdate {
	return usu.SetUserID(u.ID)
}

// Mutation returns the UserSessionMutation object of the builder.
func (usu *UserSessionUpdate) Mutation() *UserSessionMutation {
	return usu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (usu *UserSessionUpdate) ClearUser() *UserSessionUpdate {
	usu.mutation.ClearUser()
	return usu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserSessionUpdate) Save(ctx context.Context) (int, error) {
	usu.defaults()
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserSessionUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserSessionUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserSessionUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usu *UserSessionUpdate) defaults() {
	if _, ok := usu.mutation.UpdatedAt(); !ok {
		v := usersession.UpdateDefaultUpdatedAt()
		usu.mutation.SetUpdatedAt(v)
	}
}

func (usu *UserSessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(usersession.Table, usersession.Columns, sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.UpdatedAt(); ok {
		_spec.SetField(usersession.FieldUpdatedAt, field.TypeTime, value)
	}
	if usu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserSessionUpdateOne is the builder for updating a single UserSession entity.
type UserSessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSessionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (usuo *UserSessionUpdateOne) SetUpdatedAt(t time.Time) *UserSessionUpdateOne {
	usuo.mutation.SetUpdatedAt(t)
	return usuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usuo *UserSessionUpdateOne) SetUserID(id int) *UserSessionUpdateOne {
	usuo.mutation.SetUserID(id)
	return usuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usuo *UserSessionUpdateOne) SetNillableUserID(id *int) *UserSessionUpdateOne {
	if id != nil {
		usuo = usuo.SetUserID(*id)
	}
	return usuo
}

// SetUser sets the "user" edge to the User entity.
func (usuo *UserSessionUpdateOne) SetUser(u *User) *UserSessionUpdateOne {
	return usuo.SetUserID(u.ID)
}

// Mutation returns the UserSessionMutation object of the builder.
func (usuo *UserSessionUpdateOne) Mutation() *UserSessionMutation {
	return usuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (usuo *UserSessionUpdateOne) ClearUser() *UserSessionUpdateOne {
	usuo.mutation.ClearUser()
	return usuo
}

// Where appends a list predicates to the UserSessionUpdate builder.
func (usuo *UserSessionUpdateOne) Where(ps ...predicate.UserSession) *UserSessionUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserSessionUpdateOne) Select(field string, fields ...string) *UserSessionUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserSession entity.
func (usuo *UserSessionUpdateOne) Save(ctx context.Context) (*UserSession, error) {
	usuo.defaults()
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserSessionUpdateOne) SaveX(ctx context.Context) *UserSession {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserSessionUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserSessionUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usuo *UserSessionUpdateOne) defaults() {
	if _, ok := usuo.mutation.UpdatedAt(); !ok {
		v := usersession.UpdateDefaultUpdatedAt()
		usuo.mutation.SetUpdatedAt(v)
	}
}

func (usuo *UserSessionUpdateOne) sqlSave(ctx context.Context) (_node *UserSession, err error) {
	_spec := sqlgraph.NewUpdateSpec(usersession.Table, usersession.Columns, sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserSession.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersession.FieldID)
		for _, f := range fields {
			if !usersession.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usersession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usersession.FieldUpdatedAt, field.TypeTime, value)
	}
	if usuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserSession{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}
