package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("city").NotEmpty(),
		field.String("street").NotEmpty(),
		field.String("zip").Match(regexp.MustCompile("[0-9]{5}")),
		field.String("number").NotEmpty(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return nil
}
