// Code generated by ent, DO NOT EDIT.

package deliveryaddress

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the deliveryaddress type in the database.
	Label = "delivery_address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "street"
	// FieldZip holds the string denoting the zip field in the database.
	FieldZip = "zip"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTelephone holds the string denoting the telephone edge name in mutations.
	EdgeTelephone = "telephone"
	// EdgeNotes holds the string denoting the notes edge name in mutations.
	EdgeNotes = "notes"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// Table holds the table name of the deliveryaddress in the database.
	Table = "delivery_addresses"
	// TelephoneTable is the table that holds the telephone relation/edge.
	TelephoneTable = "delivery_addresses"
	// TelephoneInverseTable is the table name for the Tel entity.
	// It exists in this package in order to avoid circular dependency with the "tel" package.
	TelephoneInverseTable = "tels"
	// TelephoneColumn is the table column denoting the telephone relation/edge.
	TelephoneColumn = "delivery_address_telephone"
	// NotesTable is the table that holds the notes relation/edge.
	NotesTable = "notes"
	// NotesInverseTable is the table name for the Note entity.
	// It exists in this package in order to avoid circular dependency with the "note" package.
	NotesInverseTable = "notes"
	// NotesColumn is the table column denoting the notes relation/edge.
	NotesColumn = "delivery_address_notes"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "delivery_addresses"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_delivery_addresses"
)

// Columns holds all SQL columns for deliveryaddress fields.
var Columns = []string{
	FieldID,
	FieldCity,
	FieldStreet,
	FieldZip,
	FieldNumber,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "delivery_addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"customer_delivery_addresses",
	"delivery_address_telephone",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// StreetValidator is a validator for the "street" field. It is called by the builders before save.
	StreetValidator func(string) error
	// ZipValidator is a validator for the "zip" field. It is called by the builders before save.
	ZipValidator func(string) error
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the DeliveryAddress queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByStreet orders the results by the street field.
func ByStreet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStreet, opts...).ToFunc()
}

// ByZip orders the results by the zip field.
func ByZip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldZip, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTelephoneField orders the results by telephone field.
func ByTelephoneField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTelephoneStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotesCount orders the results by notes count.
func ByNotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotesStep(), opts...)
	}
}

// ByNotes orders the results by notes terms.
func ByNotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newTelephoneStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TelephoneInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TelephoneTable, TelephoneColumn),
	)
}
func newNotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotesTable, NotesColumn),
	)
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
