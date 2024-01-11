package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Note struct {
	ent.Schema
}

func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("notes").Unique(),
		edge.From("orders", Order.Type).Ref("notes").Unique(),
		edge.From("billing_address", BillingAddress.Type).Ref("notes").Unique(),
		edge.From("delivery_address", DeliveryAddress.Type).Ref("notes").Unique(),
		edge.From("tel", Tel.Type).Ref("note").Unique().Unique(),
	}
}
