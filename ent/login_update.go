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
	"github.com/derfinlay/basecrm/ent/login"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// LoginUpdate is the builder for updating Login entities.
type LoginUpdate struct {
	config
	hooks    []Hook
	mutation *LoginMutation
}

// Where appends a list predicates to the LoginUpdate builder.
func (lu *LoginUpdate) Where(ps ...predicate.Login) *LoginUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUsername sets the "username" field.
func (lu *LoginUpdate) SetUsername(s string) *LoginUpdate {
	lu.mutation.SetUsername(s)
	return lu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (lu *LoginUpdate) SetNillableUsername(s *string) *LoginUpdate {
	if s != nil {
		lu.SetUsername(*s)
	}
	return lu
}

// SetPassword sets the "password" field.
func (lu *LoginUpdate) SetPassword(s string) *LoginUpdate {
	lu.mutation.SetPassword(s)
	return lu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (lu *LoginUpdate) SetNillablePassword(s *string) *LoginUpdate {
	if s != nil {
		lu.SetPassword(*s)
	}
	return lu
}

// SetLastLogin sets the "last_login" field.
func (lu *LoginUpdate) SetLastLogin(t time.Time) *LoginUpdate {
	lu.mutation.SetLastLogin(t)
	return lu
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (lu *LoginUpdate) SetNillableLastLogin(t *time.Time) *LoginUpdate {
	if t != nil {
		lu.SetLastLogin(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LoginUpdate) SetUpdatedAt(t time.Time) *LoginUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// AddCustomerIDs adds the "customer" edge to the Customer entity by IDs.
func (lu *LoginUpdate) AddCustomerIDs(ids ...int) *LoginUpdate {
	lu.mutation.AddCustomerIDs(ids...)
	return lu
}

// AddCustomer adds the "customer" edges to the Customer entity.
func (lu *LoginUpdate) AddCustomer(c ...*Customer) *LoginUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.AddCustomerIDs(ids...)
}

// Mutation returns the LoginMutation object of the builder.
func (lu *LoginUpdate) Mutation() *LoginMutation {
	return lu.mutation
}

// ClearCustomer clears all "customer" edges to the Customer entity.
func (lu *LoginUpdate) ClearCustomer() *LoginUpdate {
	lu.mutation.ClearCustomer()
	return lu
}

// RemoveCustomerIDs removes the "customer" edge to Customer entities by IDs.
func (lu *LoginUpdate) RemoveCustomerIDs(ids ...int) *LoginUpdate {
	lu.mutation.RemoveCustomerIDs(ids...)
	return lu
}

// RemoveCustomer removes "customer" edges to Customer entities.
func (lu *LoginUpdate) RemoveCustomer(c ...*Customer) *LoginUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.RemoveCustomerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LoginUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LoginUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LoginUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LoginUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LoginUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := login.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LoginUpdate) check() error {
	if v, ok := lu.mutation.Username(); ok {
		if err := login.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Login.username": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Password(); ok {
		if err := login.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Login.password": %w`, err)}
		}
	}
	return nil
}

func (lu *LoginUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(login.Table, login.Columns, sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Username(); ok {
		_spec.SetField(login.FieldUsername, field.TypeString, value)
	}
	if value, ok := lu.mutation.Password(); ok {
		_spec.SetField(login.FieldPassword, field.TypeString, value)
	}
	if value, ok := lu.mutation.LastLogin(); ok {
		_spec.SetField(login.FieldLastLogin, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(login.FieldUpdatedAt, field.TypeTime, value)
	}
	if lu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   login.CustomerTable,
			Columns: []string{login.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedCustomerIDs(); len(nodes) > 0 && !lu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{login.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LoginUpdateOne is the builder for updating a single Login entity.
type LoginUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginMutation
}

// SetUsername sets the "username" field.
func (luo *LoginUpdateOne) SetUsername(s string) *LoginUpdateOne {
	luo.mutation.SetUsername(s)
	return luo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (luo *LoginUpdateOne) SetNillableUsername(s *string) *LoginUpdateOne {
	if s != nil {
		luo.SetUsername(*s)
	}
	return luo
}

// SetPassword sets the "password" field.
func (luo *LoginUpdateOne) SetPassword(s string) *LoginUpdateOne {
	luo.mutation.SetPassword(s)
	return luo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (luo *LoginUpdateOne) SetNillablePassword(s *string) *LoginUpdateOne {
	if s != nil {
		luo.SetPassword(*s)
	}
	return luo
}

// SetLastLogin sets the "last_login" field.
func (luo *LoginUpdateOne) SetLastLogin(t time.Time) *LoginUpdateOne {
	luo.mutation.SetLastLogin(t)
	return luo
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (luo *LoginUpdateOne) SetNillableLastLogin(t *time.Time) *LoginUpdateOne {
	if t != nil {
		luo.SetLastLogin(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LoginUpdateOne) SetUpdatedAt(t time.Time) *LoginUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// AddCustomerIDs adds the "customer" edge to the Customer entity by IDs.
func (luo *LoginUpdateOne) AddCustomerIDs(ids ...int) *LoginUpdateOne {
	luo.mutation.AddCustomerIDs(ids...)
	return luo
}

// AddCustomer adds the "customer" edges to the Customer entity.
func (luo *LoginUpdateOne) AddCustomer(c ...*Customer) *LoginUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.AddCustomerIDs(ids...)
}

// Mutation returns the LoginMutation object of the builder.
func (luo *LoginUpdateOne) Mutation() *LoginMutation {
	return luo.mutation
}

// ClearCustomer clears all "customer" edges to the Customer entity.
func (luo *LoginUpdateOne) ClearCustomer() *LoginUpdateOne {
	luo.mutation.ClearCustomer()
	return luo
}

// RemoveCustomerIDs removes the "customer" edge to Customer entities by IDs.
func (luo *LoginUpdateOne) RemoveCustomerIDs(ids ...int) *LoginUpdateOne {
	luo.mutation.RemoveCustomerIDs(ids...)
	return luo
}

// RemoveCustomer removes "customer" edges to Customer entities.
func (luo *LoginUpdateOne) RemoveCustomer(c ...*Customer) *LoginUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.RemoveCustomerIDs(ids...)
}

// Where appends a list predicates to the LoginUpdate builder.
func (luo *LoginUpdateOne) Where(ps ...predicate.Login) *LoginUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LoginUpdateOne) Select(field string, fields ...string) *LoginUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Login entity.
func (luo *LoginUpdateOne) Save(ctx context.Context) (*Login, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LoginUpdateOne) SaveX(ctx context.Context) *Login {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LoginUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LoginUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LoginUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := login.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LoginUpdateOne) check() error {
	if v, ok := luo.mutation.Username(); ok {
		if err := login.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Login.username": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Password(); ok {
		if err := login.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Login.password": %w`, err)}
		}
	}
	return nil
}

func (luo *LoginUpdateOne) sqlSave(ctx context.Context) (_node *Login, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(login.Table, login.Columns, sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Login.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, login.FieldID)
		for _, f := range fields {
			if !login.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != login.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Username(); ok {
		_spec.SetField(login.FieldUsername, field.TypeString, value)
	}
	if value, ok := luo.mutation.Password(); ok {
		_spec.SetField(login.FieldPassword, field.TypeString, value)
	}
	if value, ok := luo.mutation.LastLogin(); ok {
		_spec.SetField(login.FieldLastLogin, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(login.FieldUpdatedAt, field.TypeTime, value)
	}
	if luo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   login.CustomerTable,
			Columns: []string{login.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedCustomerIDs(); len(nodes) > 0 && !luo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Login{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{login.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
