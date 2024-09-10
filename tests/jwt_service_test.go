package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"viseh-api/services/jwt"
)

// TestJWTInvalidToken validates a false JWT token.
func TestJWTInvalidToken(t *testing.T) {
	invalidTokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	validJWT := jwt.ValidateJWT(invalidTokenString)
	assert.False(t, validJWT)
}

// TestJWTValidToken validates a valid token.
func TestJWTValidToken(t *testing.T) {
	tokenString := jwt.CreateJWT("John", 2, "user")

	validJWT := jwt.ValidateJWT(tokenString)
	assert.True(t, validJWT)
}

// TestJWTClaims validates 2 claims from the JWT.
func TestJWTClaims(t *testing.T) {
	userId := 2
	userRole := "user"
	tokenString := jwt.CreateJWT("John", userId, userRole)

	uid, audience, err := jwt.GetClaims(tokenString)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	assert.True(t, userId == uid)
	assert.True(t, userRole == audience)
}
