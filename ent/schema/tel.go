package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tel struct {
	ent.Schema
}

func (Tel) Fields() []ent.Field {
	return []ent.Field{
		field.String("tel").NotEmpty().Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Tel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", Note.Type),

		edge.From("customer", Customer.Type).Ref("tels").Unique(),
	}
}
