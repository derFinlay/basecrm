package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable().NotEmpty().MaxLen(50),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", Note.Type),
	}
}
