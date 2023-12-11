package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tel holds the schema definition for the Tel entity.
type Tel struct {
	ent.Schema
}

// Fields of the Tel.
func (Tel) Fields() []ent.Field {
	return []ent.Field{
		field.String("tel").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Tel.
func (Tel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("tel"),
	}
}
