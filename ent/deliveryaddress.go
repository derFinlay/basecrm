// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/deliveryaddress"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/tel"
)

// DeliveryAddress is the model entity for the DeliveryAddress schema.
type DeliveryAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Street holds the value of the "street" field.
	Street string `json:"street,omitempty"`
	// Zip holds the value of the "zip" field.
	Zip string `json:"zip,omitempty"`
	// Housenumber holds the value of the "housenumber" field.
	Housenumber string `json:"housenumber,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeliveryAddressQuery when eager-loading is set.
	Edges                       DeliveryAddressEdges `json:"edges"`
	customer_delivery_addresses *int
	delivery_address_telephone  *int
	order_delivery_address      *int
	selectValues                sql.SelectValues
}

// DeliveryAddressEdges holds the relations/edges for other nodes in the graph.
type DeliveryAddressEdges struct {
	// Telephone holds the value of the telephone edge.
	Telephone *Tel `json:"telephone,omitempty"`
	// Notes holds the value of the notes edge.
	Notes []*Note `json:"notes,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// Orders holds the value of the orders edge.
	Orders *Order `json:"orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TelephoneOrErr returns the Telephone value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeliveryAddressEdges) TelephoneOrErr() (*Tel, error) {
	if e.loadedTypes[0] {
		if e.Telephone == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tel.Label}
		}
		return e.Telephone, nil
	}
	return nil, &NotLoadedError{edge: "telephone"}
}

// NotesOrErr returns the Notes value or an error if the edge
// was not loaded in eager-loading.
func (e DeliveryAddressEdges) NotesOrErr() ([]*Note, error) {
	if e.loadedTypes[1] {
		return e.Notes, nil
	}
	return nil, &NotLoadedError{edge: "notes"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeliveryAddressEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[2] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeliveryAddressEdges) OrdersOrErr() (*Order, error) {
	if e.loadedTypes[3] {
		if e.Orders == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeliveryAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deliveryaddress.FieldID:
			values[i] = new(sql.NullInt64)
		case deliveryaddress.FieldCity, deliveryaddress.FieldStreet, deliveryaddress.FieldZip, deliveryaddress.FieldHousenumber:
			values[i] = new(sql.NullString)
		case deliveryaddress.FieldCreatedAt, deliveryaddress.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deliveryaddress.ForeignKeys[0]: // customer_delivery_addresses
			values[i] = new(sql.NullInt64)
		case deliveryaddress.ForeignKeys[1]: // delivery_address_telephone
			values[i] = new(sql.NullInt64)
		case deliveryaddress.ForeignKeys[2]: // order_delivery_address
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeliveryAddress fields.
func (da *DeliveryAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deliveryaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			da.ID = int(value.Int64)
		case deliveryaddress.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				da.City = value.String
			}
		case deliveryaddress.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field street", values[i])
			} else if value.Valid {
				da.Street = value.String
			}
		case deliveryaddress.FieldZip:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zip", values[i])
			} else if value.Valid {
				da.Zip = value.String
			}
		case deliveryaddress.FieldHousenumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field housenumber", values[i])
			} else if value.Valid {
				da.Housenumber = value.String
			}
		case deliveryaddress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				da.CreatedAt = value.Time
			}
		case deliveryaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				da.UpdatedAt = value.Time
			}
		case deliveryaddress.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_delivery_addresses", value)
			} else if value.Valid {
				da.customer_delivery_addresses = new(int)
				*da.customer_delivery_addresses = int(value.Int64)
			}
		case deliveryaddress.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field delivery_address_telephone", value)
			} else if value.Valid {
				da.delivery_address_telephone = new(int)
				*da.delivery_address_telephone = int(value.Int64)
			}
		case deliveryaddress.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_delivery_address", value)
			} else if value.Valid {
				da.order_delivery_address = new(int)
				*da.order_delivery_address = int(value.Int64)
			}
		default:
			da.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DeliveryAddress.
// This includes values selected through modifiers, order, etc.
func (da *DeliveryAddress) Value(name string) (ent.Value, error) {
	return da.selectValues.Get(name)
}

// QueryTelephone queries the "telephone" edge of the DeliveryAddress entity.
func (da *DeliveryAddress) QueryTelephone() *TelQuery {
	return NewDeliveryAddressClient(da.config).QueryTelephone(da)
}

// QueryNotes queries the "notes" edge of the DeliveryAddress entity.
func (da *DeliveryAddress) QueryNotes() *NoteQuery {
	return NewDeliveryAddressClient(da.config).QueryNotes(da)
}

// QueryCustomer queries the "customer" edge of the DeliveryAddress entity.
func (da *DeliveryAddress) QueryCustomer() *CustomerQuery {
	return NewDeliveryAddressClient(da.config).QueryCustomer(da)
}

// QueryOrders queries the "orders" edge of the DeliveryAddress entity.
func (da *DeliveryAddress) QueryOrders() *OrderQuery {
	return NewDeliveryAddressClient(da.config).QueryOrders(da)
}

// Update returns a builder for updating this DeliveryAddress.
// Note that you need to call DeliveryAddress.Unwrap() before calling this method if this DeliveryAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (da *DeliveryAddress) Update() *DeliveryAddressUpdateOne {
	return NewDeliveryAddressClient(da.config).UpdateOne(da)
}

// Unwrap unwraps the DeliveryAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (da *DeliveryAddress) Unwrap() *DeliveryAddress {
	_tx, ok := da.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeliveryAddress is not a transactional entity")
	}
	da.config.driver = _tx.drv
	return da
}

// String implements the fmt.Stringer.
func (da *DeliveryAddress) String() string {
	var builder strings.Builder
	builder.WriteString("DeliveryAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", da.ID))
	builder.WriteString("city=")
	builder.WriteString(da.City)
	builder.WriteString(", ")
	builder.WriteString("street=")
	builder.WriteString(da.Street)
	builder.WriteString(", ")
	builder.WriteString("zip=")
	builder.WriteString(da.Zip)
	builder.WriteString(", ")
	builder.WriteString("housenumber=")
	builder.WriteString(da.Housenumber)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(da.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(da.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DeliveryAddresses is a parsable slice of DeliveryAddress.
type DeliveryAddresses []*DeliveryAddress
