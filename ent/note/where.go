// Code generated by ent, DO NOT EDIT.

package note

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldContent, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldCreatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Note {
	return predicate.Note(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Note {
	return predicate.Note(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Note {
	return predicate.Note(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Note {
	return predicate.Note(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Note {
	return predicate.Note(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Note {
	return predicate.Note(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Note {
	return predicate.Note(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Note {
	return predicate.Note(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Note {
	return predicate.Note(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Note {
	return predicate.Note(sql.FieldContainsFold(FieldContent, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldCreatedAt, v))
}

// HasBillingAddress applies the HasEdge predicate on the "billing_address" edge.
func HasBillingAddress() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BillingAddressTable, BillingAddressColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingAddressWith applies the HasEdge predicate on the "billing_address" edge with a given conditions (other predicates).
func HasBillingAddressWith(preds ...predicate.BillingAddress) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newBillingAddressStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeliveryAddress applies the HasEdge predicate on the "delivery_address" edge.
func HasDeliveryAddress() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeliveryAddressTable, DeliveryAddressColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeliveryAddressWith applies the HasEdge predicate on the "delivery_address" edge with a given conditions (other predicates).
func HasDeliveryAddressWith(preds ...predicate.DeliveryAddress) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newDeliveryAddressStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoginReset applies the HasEdge predicate on the "login_reset" edge.
func HasLoginReset() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LoginResetTable, LoginResetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoginResetWith applies the HasEdge predicate on the "login_reset" edge with a given conditions (other predicates).
func HasLoginResetWith(preds ...predicate.LoginReset) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newLoginResetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogin applies the HasEdge predicate on the "login" edge.
func HasLogin() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LoginTable, LoginColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoginWith applies the HasEdge predicate on the "login" edge with a given conditions (other predicates).
func HasLoginWith(preds ...predicate.Login) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newLoginStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosition applies the HasEdge predicate on the "position" edge.
func HasPosition() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionWith applies the HasEdge predicate on the "position" edge with a given conditions (other predicates).
func HasPositionWith(preds ...predicate.Position) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newPositionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTel applies the HasEdge predicate on the "tel" edge.
func HasTel() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TelTable, TelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTelWith applies the HasEdge predicate on the "tel" edge with a given conditions (other predicates).
func HasTelWith(preds ...predicate.Tel) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newTelStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedBy applies the HasEdge predicate on the "created_by" edge.
func HasCreatedBy() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "created_by" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.User) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newCreatedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Note) predicate.Note {
	return predicate.Note(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Note) predicate.Note {
	return predicate.Note(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Note) predicate.Note {
	return predicate.Note(sql.NotPredicates(p))
}
