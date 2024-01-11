// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/derfinlay/basecrm/ent/deliveryaddress"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// DeliveryAddressDelete is the builder for deleting a DeliveryAddress entity.
type DeliveryAddressDelete struct {
	config
	hooks    []Hook
	mutation *DeliveryAddressMutation
}

// Where appends a list predicates to the DeliveryAddressDelete builder.
func (dad *DeliveryAddressDelete) Where(ps ...predicate.DeliveryAddress) *DeliveryAddressDelete {
	dad.mutation.Where(ps...)
	return dad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dad *DeliveryAddressDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dad.sqlExec, dad.mutation, dad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dad *DeliveryAddressDelete) ExecX(ctx context.Context) int {
	n, err := dad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dad *DeliveryAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(deliveryaddress.Table, sqlgraph.NewFieldSpec(deliveryaddress.FieldID, field.TypeInt))
	if ps := dad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dad.mutation.done = true
	return affected, err
}

// DeliveryAddressDeleteOne is the builder for deleting a single DeliveryAddress entity.
type DeliveryAddressDeleteOne struct {
	dad *DeliveryAddressDelete
}

// Where appends a list predicates to the DeliveryAddressDelete builder.
func (dado *DeliveryAddressDeleteOne) Where(ps ...predicate.DeliveryAddress) *DeliveryAddressDeleteOne {
	dado.dad.mutation.Where(ps...)
	return dado
}

// Exec executes the deletion query.
func (dado *DeliveryAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := dado.dad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deliveryaddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dado *DeliveryAddressDeleteOne) ExecX(ctx context.Context) {
	if err := dado.Exec(ctx); err != nil {
		panic(err)
	}
}