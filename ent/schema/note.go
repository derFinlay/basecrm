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
		field.String("title").MaxLen(200),
		field.String("content").MaxLen(1000),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("billing_address", BillingAddress.Type).Ref("notes").Unique(),
		edge.From("customer", Customer.Type).Ref("notes").Unique(),
		edge.From("delivery_address", DeliveryAddress.Type).Ref("notes").Unique(),
		edge.From("login_reset", LoginReset.Type).Ref("notes").Unique(),
		edge.From("login", Login.Type).Ref("notes").Unique(),
		edge.From("order", Order.Type).Ref("notes").Unique(),
		edge.From("position", Position.Type).Ref("notes").Unique(),
		edge.From("product", Product.Type).Ref("notes").Unique(),
		edge.From("role", Role.Type).Ref("notes").Unique(),
		edge.From("tel", Tel.Type).Ref("notes").Unique(),

		edge.To("created_by", User.Type).Unique(),
	}
}
