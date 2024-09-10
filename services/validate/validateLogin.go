package validate

import "viseh-api/types"

func LoginIsValid(l *types.UserLogin) bool {
	if len(l.Email) > 0 &&
		len(l.Password) > 0 {
		return true
	}
	return false
}
