package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type BillingAddress struct {
	ent.Schema
}

func (BillingAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("city").NotEmpty().Immutable(),
		field.String("street").NotEmpty().Immutable(),
		field.String("zip").NotEmpty().Match(regexp.MustCompile("[0-9]{5}")).Immutable(),
		field.String("housenumber").NotEmpty().Match(regexp.MustCompile("\\d+.*")).Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (BillingAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", Note.Type),

		edge.From("customer", Customer.Type).Ref("billing_addresses").Unique(),
		edge.From("order", Order.Type).Ref("billing_address").Unique(),
	}
}
