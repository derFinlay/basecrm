package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
		field.Float("price").Validate(ValidateUnitPrice),
	}
}

func (Product) Edges() []ent.Edge {
	return nil
}
