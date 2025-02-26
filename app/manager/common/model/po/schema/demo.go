package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Test holds the schema definition for the Test entity.
type Demo struct {
	ent.Schema
}

// Fields of the Test.
func (Demo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Positive(),
		field.String("name").
			Optional().
			MaxLen(10),
		field.Time("create_time").
			Optional().
			Nillable(),
	}
}

// Edges of the Test.
func (Demo) Edges() []ent.Edge {
	return nil
}
