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
		field.String("city").NotEmpty(),
		field.String("street").NotEmpty(),
		field.String("zip").Match(regexp.MustCompile("[0-9]{5}")),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (BillingAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("billing_addresses").Unique(),

		edge.To("notes", Note.Type),
	}
}
