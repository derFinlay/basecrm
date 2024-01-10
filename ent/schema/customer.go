package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MaxLen(20),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}

}

func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
		edge.To("billing_address", BillingAddress.Type),
		edge.To("delivery_addresses", BillingAddress.Type),
		edge.To("tels", Tel.Type),
		edge.To("created_by", User.Type).Unique(),
		edge.To("notes", Note.Type),
		edge.To("login", Login.Type).Unique(),
	}
}
