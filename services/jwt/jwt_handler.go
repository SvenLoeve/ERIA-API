package jwt

import (
	"fmt"
	"net/http"
	"os"
)

// Handler authorizes all endpoints with JSON Web Tokens.
// requiredRole: configures role permissions.
// Returns true if successfully authorized the user.
// Returns false if authorization failed.
func Handler(w http.ResponseWriter, r *http.Request, requiredRole int, uid int) bool {
	enabled := os.Getenv("JWT_ENABLED")
	if enabled == "false" { // JWT is disabled. For more see .env
		return true
	}
	jwtToken := "" // Fail safe
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) > 7 {
		jwtToken = bearerToken[len("Bearer "):]
	}

	if jwtToken == "" { // No JWT found.
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Bearer token missing."))
		if err != nil {
			fmt.Println(err)
		}
		return false
	} else if !ValidateJWT(jwtToken) {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Token as either expired or already revoked."))
		if err != nil {
			fmt.Println(err)
		}
		return false
	}

	uidClaim, audience, err := GetClaims(jwtToken)
	if err != nil {
		fmt.Println(err)
		return false
	}

	uidMatch := false
	if uid != 0 {
		uidMatch = uidClaim == uid
	}

	if hasPermissions(audience, requiredRole, uidMatch) {
		return true // Grant access. User role has permissions or JWT token uid matches the endpoint uid.
	}
	w.WriteHeader(http.StatusUnauthorized)
	_, err = w.Write([]byte("Invalid audience. Not authorized."))
	if err != nil {
		fmt.Println(err)
	}
	return false

}
