// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/derfinlay/basecrm/ent/position"
)

// Position is the model entity for the Position schema.
type Position struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Position) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case position.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Position fields.
func (po *Position) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case position.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Position.
// This includes values selected through modifiers, order, etc.
func (po *Position) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// Update returns a builder for updating this Position.
// Note that you need to call Position.Unwrap() before calling this method if this Position
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Position) Update() *PositionUpdateOne {
	return NewPositionClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Position entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Position) Unwrap() *Position {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Position is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Position) String() string {
	var builder strings.Builder
	builder.WriteString("Position(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Positions is a parsable slice of Position.
type Positions []*Position