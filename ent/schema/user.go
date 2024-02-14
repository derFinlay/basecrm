package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MaxLen(20),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.Time("last_login").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customers", Customer.Type).Ref("created_by"),
		edge.To("notes", Note.Type),
		edge.From("orders", Order.Type).Ref("created_by"),
		edge.From("sessions", UserSession.Type).Ref("user"),
	}
}
