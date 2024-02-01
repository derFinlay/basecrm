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
		field.Float("tax").Immutable(),
		field.Time("due"), //time when payment should be done
		field.Time("printed_at"),
		field.Time("payed_at"),
		field.Time("done_at"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", Customer.Type).Unique(),
		edge.To("billing_address", BillingAddress.Type).Unique(),
		edge.To("delivery_address", DeliveryAddress.Type).Unique(),
		edge.To("notes", Note.Type),
		edge.To("created_by", User.Type).Unique(),
		edge.To("positions", Position.Type),
	}
}
