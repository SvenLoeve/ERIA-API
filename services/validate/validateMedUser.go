package validate

import (
	"viseh-api/database/ent"
)

func MedUserIsValid(mu *ent.MedUser) bool {
	if len(mu.Name) > 0 &&
		len(mu.Email) > 0 &&
		len(mu.Password) > 0 &&
		len(mu.PhoneNumber) > 0 &&
		len(mu.Role) > 0 &&
		len(mu.Organisation) > 0 {
		return true
	}
	return false
}
