package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/getkin/kin-openapi/openapi3"
	"viseh-api/types"
)

// ############## MUST DO ########################
// local copy of types to fix unintended behaviour
// ###############################################
type patchUser *types.User
type patchMedUser *types.MedUser
type user *types.FullUser
type medUser *types.MedUser

func RegisterModels() {
	api.RegisterModel(rest.ModelOf[respond.Error](), rest.WithDescription("Standard JSON error"), func(s *openapi3.Schema) {
		status := s.Properties["statusCode"]
		status.Value.WithMin(100).WithMax(600)
	})
	api.RegisterModel(rest.ModelOf[patchUser](), rest.WithDescription("User object with values allowed to be patched"), func(s *openapi3.Schema) {
		s.Required = []string{""}
	})
	api.RegisterModel(rest.ModelOf[patchMedUser](), rest.WithDescription("Medical user object with values allowed to be patched"), func(s *openapi3.Schema) {
		s.Required = []string{""}
	})
	api.RegisterModel(rest.ModelOf[user](), rest.WithDescription("User object"), func(s *openapi3.Schema) {
		// id
		id := s.Properties["id"]
		id.Value.Description = "unique user id"
		id.Value.Example = 1
		// name
		name := s.Properties["name"]
		name.Value.Description = "first and last name"
		name.Value.Example = gofakeit.Name()
		// email
		email := s.Properties["email"]
		email.Value.Description = "email address"
		email.Value.Example = gofakeit.Email()
		// phone number
		phoneNumber := s.Properties["phone_number"]
		phoneNumber.Value.Description = "phone number"
		phoneNumber.Value.Example = gofakeit.PhoneFormatted()
		// med id
		medId := s.Properties["med_id"]
		medId.Value.Description = "unique medical id"
		medId.Value.Example = "12345abc"
		// role
		role := s.Properties["role"]
		role.Value.Description = "role"
		role.Value.WithEnum("user", "employee", "admin")
		role.Value.Example = "user"
	})
	api.RegisterModel(rest.ModelOf[medUser](), rest.WithDescription("User object"), func(s *openapi3.Schema) {
		// id
		id := s.Properties["id"]
		id.Value.Description = "unique user id"
		id.Value.Example = 1
		// name
		name := s.Properties["name"]
		name.Value.Description = "first and last name"
		name.Value.Example = gofakeit.Name()
		// email
		email := s.Properties["email"]
		email.Value.Description = "email address"
		email.Value.Example = gofakeit.Email()
		// phone number
		phoneNumber := s.Properties["phone_number"]
		phoneNumber.Value.Description = "phone number"
		phoneNumber.Value.Example = gofakeit.PhoneFormatted()
		// role
		role := s.Properties["role"]
		role.Value.Description = "role"
		role.Value.WithEnum("ambulance", "doctor")
		role.Value.Example = "ambulance"
		// organisation
		organisation := s.Properties["organisation"]
		organisation.Value.Description = "medical organisation"
		organisation.Value.Example = gofakeit.Company()
	})
}
