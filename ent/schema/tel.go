package schema

import "entgo.io/ent"

// Tel holds the schema definition for the Tel entity.
type Tel struct {
	ent.Schema
}

// Fields of the Tel.
func (Tel) Fields() []ent.Field {
	return nil
}

// Edges of the Tel.
func (Tel) Edges() []ent.Edge {
	return nil
}
