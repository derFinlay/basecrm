package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LoginReset struct {
	ent.Schema
}

func (LoginReset) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").NotEmpty().Unique().Immutable(),
		field.Bool("active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (LoginReset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("login", Login.Type).Ref("login_resets").Unique(),
	}
}
