package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MedUser User holds the schema definition for the MedUser entity.
type MedUser struct {
	ent.Schema
}

// Fields of the User.
func (MedUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("john doe").
			MaxLen(50),
		field.String("email").
			Unique().
			MaxLen(100),
		field.String("password"),
		field.String("phone_number"),
		field.Enum("role").
			Values("ambulance", "doctor"),
		field.String("organisation"),
	}
}

// Edges of the User.
func (MedUser) Edges() []ent.Edge {
	return nil
}
