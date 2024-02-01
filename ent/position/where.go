// Code generated by ent, DO NOT EDIT.

package position

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDescription, v))
}

// UnitPrice applies equality check predicate on the "unit_price" field. It's identical to UnitPriceEQ.
func UnitPrice(v float64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUnitPrice, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldAmount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldDescription, v))
}

// UnitPriceEQ applies the EQ predicate on the "unit_price" field.
func UnitPriceEQ(v float64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUnitPrice, v))
}

// UnitPriceNEQ applies the NEQ predicate on the "unit_price" field.
func UnitPriceNEQ(v float64) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldUnitPrice, v))
}

// UnitPriceIn applies the In predicate on the "unit_price" field.
func UnitPriceIn(vs ...float64) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldUnitPrice, vs...))
}

// UnitPriceNotIn applies the NotIn predicate on the "unit_price" field.
func UnitPriceNotIn(vs ...float64) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldUnitPrice, vs...))
}

// UnitPriceGT applies the GT predicate on the "unit_price" field.
func UnitPriceGT(v float64) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldUnitPrice, v))
}

// UnitPriceGTE applies the GTE predicate on the "unit_price" field.
func UnitPriceGTE(v float64) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldUnitPrice, v))
}

// UnitPriceLT applies the LT predicate on the "unit_price" field.
func UnitPriceLT(v float64) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldUnitPrice, v))
}

// UnitPriceLTE applies the LTE predicate on the "unit_price" field.
func UnitPriceLTE(v float64) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldUnitPrice, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldAmount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Position) predicate.Position {
	return predicate.Position(sql.NotPredicates(p))
}
