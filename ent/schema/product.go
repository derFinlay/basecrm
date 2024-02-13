package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable().NotEmpty().Unique(),
		field.String("description").Optional().Immutable(),
		field.Float("price").Validate(ValidateUnitPrice),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", Note.Type),
	}
}
