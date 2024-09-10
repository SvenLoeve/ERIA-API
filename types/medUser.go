package types

import "viseh-api/database/ent"

type MedUser struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Role         string `json:"role"`
	Organisation string `json:"organisation"`
}
type MedUsers []MedUser

func ParseMedUser(user *ent.MedUser) MedUser {
	return MedUser{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Role:         user.Role.String(),
		Organisation: user.Organisation,
	}
}

type FullMedUser struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Role         string `json:"role,omitempty"`
	Organisation string `json:"organisation,omitempty"`
	Password     string `json:"password,omitempty"`
}
