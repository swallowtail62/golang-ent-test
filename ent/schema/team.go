package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return nil
}
