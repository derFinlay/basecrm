package schema

import (
	"errors"
	"math"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

var ErrInvalidUnitPrice = errors.New("invalid unit price")

func ValidateUnitPrice(unitPrice float64) error {
	//check if price has more than 2 decimal places
	if math.Round((unitPrice*100)/100) != unitPrice {
		return ErrInvalidUnitPrice
	}
	return nil
}

type Position struct {
	ent.Schema
}

func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(40).Immutable(),
		field.String("description").Optional().MaxLen(20).Immutable(),
		field.Float("unit_price").Validate(ValidateUnitPrice).Immutable(),
		field.Int("amount").Min(1).Max(1000).Default(1).Immutable(),
		field.Time("created_at").Default(time.Now()).Immutable(),
	}
}

func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", Note.Type),

		edge.From("order", Order.Type).Ref("positions").Unique(),
	}
}
