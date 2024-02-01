// Code generated by ent, DO NOT EDIT.

package deliveryaddress

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/derfinlay/basecrm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldID, id))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldCity, v))
}

// Street applies equality check predicate on the "street" field. It's identical to StreetEQ.
func Street(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldStreet, v))
}

// Zip applies equality check predicate on the "zip" field. It's identical to ZipEQ.
func Zip(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldZip, v))
}

// Housenumber applies equality check predicate on the "housenumber" field. It's identical to HousenumberEQ.
func Housenumber(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldHousenumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContainsFold(FieldCity, v))
}

// StreetEQ applies the EQ predicate on the "street" field.
func StreetEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldStreet, v))
}

// StreetNEQ applies the NEQ predicate on the "street" field.
func StreetNEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldStreet, v))
}

// StreetIn applies the In predicate on the "street" field.
func StreetIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldStreet, vs...))
}

// StreetNotIn applies the NotIn predicate on the "street" field.
func StreetNotIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldStreet, vs...))
}

// StreetGT applies the GT predicate on the "street" field.
func StreetGT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldStreet, v))
}

// StreetGTE applies the GTE predicate on the "street" field.
func StreetGTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldStreet, v))
}

// StreetLT applies the LT predicate on the "street" field.
func StreetLT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldStreet, v))
}

// StreetLTE applies the LTE predicate on the "street" field.
func StreetLTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldStreet, v))
}

// StreetContains applies the Contains predicate on the "street" field.
func StreetContains(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContains(FieldStreet, v))
}

// StreetHasPrefix applies the HasPrefix predicate on the "street" field.
func StreetHasPrefix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasPrefix(FieldStreet, v))
}

// StreetHasSuffix applies the HasSuffix predicate on the "street" field.
func StreetHasSuffix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasSuffix(FieldStreet, v))
}

// StreetEqualFold applies the EqualFold predicate on the "street" field.
func StreetEqualFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEqualFold(FieldStreet, v))
}

// StreetContainsFold applies the ContainsFold predicate on the "street" field.
func StreetContainsFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContainsFold(FieldStreet, v))
}

// ZipEQ applies the EQ predicate on the "zip" field.
func ZipEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldZip, v))
}

// ZipNEQ applies the NEQ predicate on the "zip" field.
func ZipNEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldZip, v))
}

// ZipIn applies the In predicate on the "zip" field.
func ZipIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldZip, vs...))
}

// ZipNotIn applies the NotIn predicate on the "zip" field.
func ZipNotIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldZip, vs...))
}

// ZipGT applies the GT predicate on the "zip" field.
func ZipGT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldZip, v))
}

// ZipGTE applies the GTE predicate on the "zip" field.
func ZipGTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldZip, v))
}

// ZipLT applies the LT predicate on the "zip" field.
func ZipLT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldZip, v))
}

// ZipLTE applies the LTE predicate on the "zip" field.
func ZipLTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldZip, v))
}

// ZipContains applies the Contains predicate on the "zip" field.
func ZipContains(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContains(FieldZip, v))
}

// ZipHasPrefix applies the HasPrefix predicate on the "zip" field.
func ZipHasPrefix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasPrefix(FieldZip, v))
}

// ZipHasSuffix applies the HasSuffix predicate on the "zip" field.
func ZipHasSuffix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasSuffix(FieldZip, v))
}

// ZipEqualFold applies the EqualFold predicate on the "zip" field.
func ZipEqualFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEqualFold(FieldZip, v))
}

// ZipContainsFold applies the ContainsFold predicate on the "zip" field.
func ZipContainsFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContainsFold(FieldZip, v))
}

// HousenumberEQ applies the EQ predicate on the "housenumber" field.
func HousenumberEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldHousenumber, v))
}

// HousenumberNEQ applies the NEQ predicate on the "housenumber" field.
func HousenumberNEQ(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldHousenumber, v))
}

// HousenumberIn applies the In predicate on the "housenumber" field.
func HousenumberIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldHousenumber, vs...))
}

// HousenumberNotIn applies the NotIn predicate on the "housenumber" field.
func HousenumberNotIn(vs ...string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldHousenumber, vs...))
}

// HousenumberGT applies the GT predicate on the "housenumber" field.
func HousenumberGT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldHousenumber, v))
}

// HousenumberGTE applies the GTE predicate on the "housenumber" field.
func HousenumberGTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldHousenumber, v))
}

// HousenumberLT applies the LT predicate on the "housenumber" field.
func HousenumberLT(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldHousenumber, v))
}

// HousenumberLTE applies the LTE predicate on the "housenumber" field.
func HousenumberLTE(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldHousenumber, v))
}

// HousenumberContains applies the Contains predicate on the "housenumber" field.
func HousenumberContains(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContains(FieldHousenumber, v))
}

// HousenumberHasPrefix applies the HasPrefix predicate on the "housenumber" field.
func HousenumberHasPrefix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasPrefix(FieldHousenumber, v))
}

// HousenumberHasSuffix applies the HasSuffix predicate on the "housenumber" field.
func HousenumberHasSuffix(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldHasSuffix(FieldHousenumber, v))
}

// HousenumberEqualFold applies the EqualFold predicate on the "housenumber" field.
func HousenumberEqualFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEqualFold(FieldHousenumber, v))
}

// HousenumberContainsFold applies the ContainsFold predicate on the "housenumber" field.
func HousenumberContainsFold(v string) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldContainsFold(FieldHousenumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasTelephone applies the HasEdge predicate on the "telephone" edge.
func HasTelephone() predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TelephoneTable, TelephoneColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTelephoneWith applies the HasEdge predicate on the "telephone" edge with a given conditions (other predicates).
func HasTelephoneWith(preds ...predicate.Tel) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := newTelephoneStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotes applies the HasEdge predicate on the "notes" edge.
func HasNotes() predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotesTable, NotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotesWith applies the HasEdge predicate on the "notes" edge with a given conditions (other predicates).
func HasNotesWith(preds ...predicate.Note) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := newNotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(func(s *sql.Selector) {
		step := newOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeliveryAddress) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeliveryAddress) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeliveryAddress) predicate.DeliveryAddress {
	return predicate.DeliveryAddress(sql.NotPredicates(p))
}
