package validate

import (
	"viseh-api/types"
)

func NfcIsValid(c types.Chip) bool {
	if len(c.Data.Name) > 0 &&
		len(c.Keys) > 0 &&
		c.Version > 0 {
		return true
	}
	return false
}
func NfcEncryptedIsValid(c types.ChipEncryptedV2) bool {
	if len(c.Data) > 0 &&
		len(c.Keys) > 0 &&
		c.Version > 0 {
		return true
	}
	return false
}
