package schema

import "entgo.io/ent"

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return nil
}

func (Role) Edges() []ent.Edge {
	return nil
}
