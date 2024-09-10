package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Key holds the schema definition for the Key entity.
type Key struct {
	ent.Schema
}

// Fields of the Chip.
func (Key) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Default(time.Now().Format("2006/02/01")),
		field.String("key"),
	}
}

// Edges of the Chip.
func (Key) Edges() []ent.Edge {
	return nil
}
