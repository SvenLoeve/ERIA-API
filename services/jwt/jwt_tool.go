package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
	"time"
)

var secret = os.Getenv("JWT_SECRET")

// CreateJWT creates a signed JSON Web Token for the user.
func CreateJWT(name string, uid int, role string) string {
	// Create claims for the JSON Web Token.
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "Viseh",
		Subject:   name,
		ID:        strconv.Itoa(uid),
		Audience:  []string{role},
	}
	// Sign a JSON Web Token with the given claims.
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := t.SignedString([]byte(secret))
	return token
}

// ValidateJWT validates whether the token is valid or invalid.
func ValidateJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

// GetClaims parses the token and returns user object.
func GetClaims(tokenString string) (int, string, error) {
	// Parse JWT
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, "", err
	}

	uid, err := strconv.Atoi(token.Claims.(*jwt.RegisteredClaims).ID)
	audience := token.Claims.(*jwt.RegisteredClaims).Audience[0]
	return uid, audience, err
}

func IsAdministrator(r *http.Request) bool {
	enabled := os.Getenv("JWT_ENABLED")
	if enabled == "false" {
		return true
	}
	bearerToken := r.Header.Get("Authorization")
	jwtToken := bearerToken[len("Bearer "):]
	_, audience, _ := GetClaims(jwtToken)

	return audience == "admin"
}
