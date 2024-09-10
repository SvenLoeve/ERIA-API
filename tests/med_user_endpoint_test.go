package tests

import (
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"viseh-api/database/ent"
	"viseh-api/services/jwt"
	"viseh-api/types"
)

// TestGetMedUserByID calls an endpoint that returns and validates the user.
func TestGetMedUserByID(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/med/users/12", nil)
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "ambulance"))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("An error occured: %v", err)
	}

	body, _err := io.ReadAll(response.Body)
	var user types.MedUser
	_err = json.Unmarshal(body, &user)
	if _err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", _err)
	}

	assert.True(t, user.Email != "" && strings.Contains(user.Email, "@"))
	assert.True(t, user.ID == 12)
	assert.True(t, user.Name != "")
}

// TestUpdateMedUser updates the users email and validates it.
func TestUpdateMedUser(t *testing.T) {
	// Step 1. Gets a med user by ID 7 from the database.
	request, err := http.NewRequest("GET", "http://localhost:8080/med/users/7", nil)
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))

	if err != nil {
		t.Fatalf("Error setting up http request: %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("An error occured: %v", err)
	}

	body, _err := io.ReadAll(response.Body)

	if _err != nil {
		t.Fatalf("Error reading response: %v", _err)
	}
	// Unmarshall JSON data
	var user types.MedUser
	_err = json.Unmarshal(body, &user)
	if _err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", _err)
	}

	// Step 2. Change the email address.
	user.Email = "changed-email-address@johndoe.com"

	jsonData, _err := json.Marshal(user)
	if _err != nil {
		t.Fatalf("Error marshalling JSON: %v", _err)
	}

	// Step 3. Create an HTTP client and request the email address to be changed.
	request, err = http.NewRequest("PUT", "http://localhost:8080/med/users/7", bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))

	if err != nil {
		t.Fatalf("An error occured %v", _err)
	}

	client = &http.Client{}
	response, err = client.Do(request)

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	defer response.Body.Close()

	// Step 4. Check if the email address has been changed.
	assert.True(t, response.StatusCode == http.StatusOK)
	assert.True(t, string(body) == "Med User Updated")
}
func TestSearchMedUser(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/med/users/search?email=johndoe", nil) // There should be at least 1 record.
	if err != nil {
		t.Fatalf("Error setting up http request: %v", err)
	}
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("Error while checking for user: %v", err)
	}
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("Error reading response: %v", err)
	}
	// Unmarshall JSON data
	var users types.MedUsers
	err = json.Unmarshal(body, &users)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Verify total found users is at least 1.
	assert.True(t, len(users) >= 1)
	assert.True(t, response.StatusCode == http.StatusOK)

	// Check if other attributes are available and set.
	var user = users[0]
	assert.True(t, len(user.Name) > 0)
	assert.True(t, len(user.Email) > 0)
	assert.True(t, len(user.PhoneNumber) > 0)
	assert.True(t, len(user.Role) > 0)
	assert.True(t, len(user.Organisation) > 0)
}

// TestChangePasswordMedUser - Patches password.
func TestChangePasswordMedUser(t *testing.T) {
	jsonData, err := json.Marshal(ent.MedUser{
		Password: "NewTestPassword",
	})
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	request, err := http.NewRequest("PATCH", "http://localhost:8080/med/users/4", bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	assert.True(t, string(body) == "Med User Updated")
	assert.True(t, response.StatusCode == http.StatusOK)
}

// TestPatchUser - Patches all fields at once (except for password field).
func TestPatchMedUser(t *testing.T) {

	user1 := ent.MedUser{
		Name:         gofakeit.Name(),
		Email:        gofakeit.Email(),
		PhoneNumber:  gofakeit.PhoneFormatted(),
		Role:         "doctor",
		Organisation: gofakeit.Company(),
	}

	jsonData, err := json.Marshal(user1)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	request, err := http.NewRequest("PATCH", "http://localhost:8080/med/users/4", bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))

	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request) // Sends Patch request to the server

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	assert.True(t, string(body) == "Med User Updated")
	assert.True(t, response.StatusCode == http.StatusOK)

	// Time to verify if the patch request works as expected.
	request, err = http.NewRequest("GET", "http://localhost:8080/med/users/4", nil)
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	client = &http.Client{}
	response, err = client.Do(request)

	if err != nil {
		t.Fatalf("Error while checking for user: %v", err)
	}

	body, _err := io.ReadAll(response.Body)
	var user2 types.MedUser
	_err = json.Unmarshal(body, &user2)
	if _err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", _err)
	}

	// Is everything received correctly from the server? Checks all the fields.
	assert.True(t, user1.Name == user2.Name)
	assert.True(t, user1.Email == user2.Email)
	assert.True(t, user1.PhoneNumber == user2.PhoneNumber)
	assert.True(t, user1.Role.String() == user2.Role)
	assert.True(t, user1.Organisation == user2.Organisation)
}
