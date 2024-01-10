package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("test"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", Customer.Type).Unique(),
		edge.To("address", BillingAddress.Type).Unique(),
		edge.To("notes", Note.Type),
	}
}
