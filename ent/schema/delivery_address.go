package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeliveryAddress struct {
	ent.Schema
}

func (DeliveryAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("city").NotEmpty().Immutable(),
		field.String("street").NotEmpty().Immutable(),
		field.String("zip").Match(regexp.MustCompile("[0-9]{5}")).Immutable(),
		field.String("number").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (DeliveryAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("telephone", Tel.Type).Unique(),

		edge.From("customer", Customer.Type).Ref("delivery_addresses"),
	}
}
