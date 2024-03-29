// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/derfinlay/basecrm/ent/billingaddress"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/deliveryaddress"
	"github.com/derfinlay/basecrm/ent/order"
	"github.com/derfinlay/basecrm/ent/user"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Tax holds the value of the "tax" field.
	Tax float64 `json:"tax,omitempty"`
	// Due holds the value of the "due" field.
	Due time.Time `json:"due,omitempty"`
	// PrintedAt holds the value of the "printed_at" field.
	PrintedAt time.Time `json:"printed_at,omitempty"`
	// PayedAt holds the value of the "payed_at" field.
	PayedAt time.Time `json:"payed_at,omitempty"`
	// DoneAt holds the value of the "done_at" field.
	DoneAt time.Time `json:"done_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges            OrderEdges `json:"edges"`
	customer_orders  *int
	order_customer   *int
	order_created_by *int
	selectValues     sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// BillingAddress holds the value of the billing_address edge.
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
	// DeliveryAddress holds the value of the delivery_address edge.
	DeliveryAddress *DeliveryAddress `json:"delivery_address,omitempty"`
	// Notes holds the value of the notes edge.
	Notes []*Note `json:"notes,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy *User `json:"created_by,omitempty"`
	// Positions holds the value of the positions edge.
	Positions []*Position `json:"positions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// BillingAddressOrErr returns the BillingAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) BillingAddressOrErr() (*BillingAddress, error) {
	if e.loadedTypes[1] {
		if e.BillingAddress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: billingaddress.Label}
		}
		return e.BillingAddress, nil
	}
	return nil, &NotLoadedError{edge: "billing_address"}
}

// DeliveryAddressOrErr returns the DeliveryAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) DeliveryAddressOrErr() (*DeliveryAddress, error) {
	if e.loadedTypes[2] {
		if e.DeliveryAddress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: deliveryaddress.Label}
		}
		return e.DeliveryAddress, nil
	}
	return nil, &NotLoadedError{edge: "delivery_address"}
}

// NotesOrErr returns the Notes value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) NotesOrErr() ([]*Note, error) {
	if e.loadedTypes[3] {
		return e.Notes, nil
	}
	return nil, &NotLoadedError{edge: "notes"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CreatedByOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.CreatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.CreatedBy, nil
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) PositionsOrErr() ([]*Position, error) {
	if e.loadedTypes[5] {
		return e.Positions, nil
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldTax:
			values[i] = new(sql.NullFloat64)
		case order.FieldID:
			values[i] = new(sql.NullInt64)
		case order.FieldDue, order.FieldPrintedAt, order.FieldPayedAt, order.FieldDoneAt, order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // customer_orders
			values[i] = new(sql.NullInt64)
		case order.ForeignKeys[1]: // order_customer
			values[i] = new(sql.NullInt64)
		case order.ForeignKeys[2]: // order_created_by
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldTax:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tax", values[i])
			} else if value.Valid {
				o.Tax = value.Float64
			}
		case order.FieldDue:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due", values[i])
			} else if value.Valid {
				o.Due = value.Time
			}
		case order.FieldPrintedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field printed_at", values[i])
			} else if value.Valid {
				o.PrintedAt = value.Time
			}
		case order.FieldPayedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field payed_at", values[i])
			} else if value.Valid {
				o.PayedAt = value.Time
			}
		case order.FieldDoneAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field done_at", values[i])
			} else if value.Valid {
				o.DoneAt = value.Time
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_orders", value)
			} else if value.Valid {
				o.customer_orders = new(int)
				*o.customer_orders = int(value.Int64)
			}
		case order.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_customer", value)
			} else if value.Valid {
				o.order_customer = new(int)
				*o.order_customer = int(value.Int64)
			}
		case order.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_created_by", value)
			} else if value.Valid {
				o.order_created_by = new(int)
				*o.order_created_by = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the Order entity.
func (o *Order) QueryCustomer() *CustomerQuery {
	return NewOrderClient(o.config).QueryCustomer(o)
}

// QueryBillingAddress queries the "billing_address" edge of the Order entity.
func (o *Order) QueryBillingAddress() *BillingAddressQuery {
	return NewOrderClient(o.config).QueryBillingAddress(o)
}

// QueryDeliveryAddress queries the "delivery_address" edge of the Order entity.
func (o *Order) QueryDeliveryAddress() *DeliveryAddressQuery {
	return NewOrderClient(o.config).QueryDeliveryAddress(o)
}

// QueryNotes queries the "notes" edge of the Order entity.
func (o *Order) QueryNotes() *NoteQuery {
	return NewOrderClient(o.config).QueryNotes(o)
}

// QueryCreatedBy queries the "created_by" edge of the Order entity.
func (o *Order) QueryCreatedBy() *UserQuery {
	return NewOrderClient(o.config).QueryCreatedBy(o)
}

// QueryPositions queries the "positions" edge of the Order entity.
func (o *Order) QueryPositions() *PositionQuery {
	return NewOrderClient(o.config).QueryPositions(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("tax=")
	builder.WriteString(fmt.Sprintf("%v", o.Tax))
	builder.WriteString(", ")
	builder.WriteString("due=")
	builder.WriteString(o.Due.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("printed_at=")
	builder.WriteString(o.PrintedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payed_at=")
	builder.WriteString(o.PayedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("done_at=")
	builder.WriteString(o.DoneAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
