package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MaxLen(20),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("phone").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}

}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
		edge.To("billing_address", Address.Type).Unique(),
		edge.To("addresses", Address.Type),
		edge.To("phone", Tel.Type),
		edge.To("created_by", User.Type).Unique(),
		edge.To("notes", Note.Type),
	}
}
