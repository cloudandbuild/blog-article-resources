package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", User.Type),
	}
}
