package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"viseh-api/services/jwt"
	"viseh-api/types"
)

// TestCheckMockData calls the server and checks the amount of users.
// If there are less than 20 users. The test fails.
func TestCheckMockData(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/users?page=1&items=20", nil)
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "admin"))

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		t.Fatalf("An error occured: %v", err)
	}

	if err != nil {
		t.Fatalf("Error while checking for mockdata: %v", err)
	} else if res.StatusCode != http.StatusOK {
		t.Fatalf("Request failed with status code: %v", res.StatusCode)
	}
	body, _err := io.ReadAll(res.Body)
	if _err != nil {
		t.Fatalf("Error reading response body: %v", _err)
	}

	var users types.Users
	_err = json.Unmarshal(body, &users)
	if _err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", _err)
	}

	if !(len(users) >= 20) {
		t.Fatalf("Not enough users. Need 20 users. Got %v users", len(users))
	}
}

// TestCheckMedicalMockData calls the server and checks the amount of users.
// If there are less than 20 medical users. The test fails.
func TestCheckMedicalMockData(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/med/users?page=1&items=20", nil)
	request.Header.Add("Authorization", "Bearer "+jwt.CreateJWT("John", 1, "doctor"))

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		t.Fatalf("An error occured: %v", err)
	}

	if err != nil {
		t.Fatalf("Error while checking for mockdata: %v", err)
	} else if res.StatusCode != http.StatusOK {
		t.Fatalf("Request failed with status code: %v", res.StatusCode)
	}
	body, _err := io.ReadAll(res.Body)
	if _err != nil {
		t.Fatalf("Error reading response body: %v", _err)
	}

	var users types.MedUsers
	_err = json.Unmarshal(body, &users)
	if _err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", _err)
	}

	if !(len(users) >= 20) {
		t.Fatalf("Not enough medical users. Need 20 medical users. Got %v medical users", len(users))
	}
}
