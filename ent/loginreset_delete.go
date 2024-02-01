// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/loginreset"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// LoginResetDelete is the builder for deleting a LoginReset entity.
type LoginResetDelete struct {
	config
	hooks    []Hook
	mutation *LoginResetMutation
}

// Where appends a list predicates to the LoginResetDelete builder.
func (lrd *LoginResetDelete) Where(ps ...predicate.LoginReset) *LoginResetDelete {
	lrd.mutation.Where(ps...)
	return lrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lrd *LoginResetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lrd.sqlExec, lrd.mutation, lrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lrd *LoginResetDelete) ExecX(ctx context.Context) int {
	n, err := lrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lrd *LoginResetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(loginreset.Table, sqlgraph.NewFieldSpec(loginreset.FieldID, field.TypeInt))
	if ps := lrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lrd.mutation.done = true
	return affected, err
}

// LoginResetDeleteOne is the builder for deleting a single LoginReset entity.
type LoginResetDeleteOne struct {
	lrd *LoginResetDelete
}

// Where appends a list predicates to the LoginResetDelete builder.
func (lrdo *LoginResetDeleteOne) Where(ps ...predicate.LoginReset) *LoginResetDeleteOne {
	lrdo.lrd.mutation.Where(ps...)
	return lrdo
}

// Exec executes the deletion query.
func (lrdo *LoginResetDeleteOne) Exec(ctx context.Context) error {
	n, err := lrdo.lrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loginreset.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lrdo *LoginResetDeleteOne) ExecX(ctx context.Context) {
	if err := lrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
