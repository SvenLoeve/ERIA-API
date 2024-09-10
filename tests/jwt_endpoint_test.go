package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"viseh-api/database/ent"
	"viseh-api/services/jwt"
)

// TestJWTEndpointNoTokenUnauthorized tests if access is blocked when no token could be found.
func TestJWTEndpointNoTokenUnauthorized(t *testing.T) {

	jsonData, err := json.Marshal(ent.User{
		Name: "John",
	})
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	request, err := http.NewRequest("PATCH", "http://localhost:8080/users/4", bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	assert.True(t, string(body) == "Bearer token missing.")
	assert.True(t, response.StatusCode == http.StatusUnauthorized)

}

// TestJWTEndpointTokenUnauthorized tests if a malicious JWT against works against the system.
func TestJWTEndpointTokenUnauthorized(t *testing.T) {
	invalidJWT := "InvalidToken"

	jsonData, err := json.Marshal(ent.User{
		Name: "John",
	})
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	request, err := http.NewRequest("PATCH", "http://localhost:8080/users/4", bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+invalidJWT)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	assert.True(t, string(body) == "Token as either expired or already revoked.")
	assert.True(t, response.StatusCode == http.StatusUnauthorized)
}

// TestJWTEndpointTokenRole tests the role permissions of the JSON Web Token.
func TestJWTEndpointTokenRole(t *testing.T) {
	jwtToken := jwt.CreateJWT("John", 1, "user")

	jsonData, err := json.Marshal(ent.User{
		Name: "John",
	})
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	request, err := http.NewRequest("PATCH", "http://localhost:8080/users/4", bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwtToken)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	assert.True(t, string(body) == "Invalid audience. Not authorized.")
	assert.True(t, response.StatusCode == http.StatusUnauthorized)
}
