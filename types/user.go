package types

import "viseh-api/database/ent"

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	MedID       string `json:"med_id"`
	Role        string `json:"role"`
}
type Users []User

func ParseUser(user *ent.User) User {
	return User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		MedID:       user.MedID,
		Role:        user.Role.String(),
	}
}

type FullUser struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	MedID       string `json:"med_id,omitempty"`
	Role        string `json:"role,omitempty"`
	Password    string `json:"password,omitempty"`
}
