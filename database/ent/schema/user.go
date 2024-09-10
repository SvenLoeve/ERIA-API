package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("john doe").
			MaxLen(50),
		field.String("email").
			Unique().
			MaxLen(100),
		field.String("password"),
		field.String("phone_number"),
		field.String("encryption_key").
			Sensitive(),
		field.String("med_id").
			Unique(),
		field.Enum("role").
			Values("user", "employee", "admin"),
		field.Bytes("data").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chips", Chip.Type),
	}
}
