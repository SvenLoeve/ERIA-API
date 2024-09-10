package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chip holds the schema definition for the Chip entity.
type Chip struct {
	ent.Schema
}

// Fields of the Chip.
func (Chip) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Lil chip").
			MaxLen(50),
		field.String("chip_type").
			MaxLen(100),
		field.Time("last_updated"),
	}
}

// Edges of the Chip.
func (Chip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("chips").
			Unique(),
	}
}
