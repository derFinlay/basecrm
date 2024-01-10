package schema

import "entgo.io/ent"

type Position struct {
	ent.Schema
}

func (Position) Fields() []ent.Field {
	return nil
}

func (Position) Edges() []ent.Edge {
	return nil
}
