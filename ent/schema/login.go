package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Login struct {
	ent.Schema
}

func (Login) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").NotEmpty(),
		field.String("email").NotEmpty(),
		field.Time("last_login"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Login) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("login").Unique(),

		edge.To("login_resets", LoginReset.Type),
	}
}
