package validate

import (
	"viseh-api/database/ent"
)

func UserIsValid(u *ent.User) bool {
	if len(u.Name) > 0 &&
		len(u.Email) > 0 &&
		len(u.Password) > 0 &&
		len(u.PhoneNumber) > 0 &&
		len(u.Role) > 0 &&
		len(u.MedID) > 0 {
		return true
	}
	return false
}
