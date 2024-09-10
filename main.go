package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"viseh-api/api/handler"
	"viseh-api/database"
	"viseh-api/services/cryptography/aes"
	"viseh-api/tests/mock"
)

func main() {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// load .env in runtime environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env not found: %v", err)
		return
	}
	//connect to database and migrate
	database.DB()
	//Start daily key gen
	aes.DailyKeyGen()

	// Register the routes and handlers
	mux.HandleFunc("/", handler.CatchAllHandler)
	mux.HandleFunc("GET /users/{uid}", handler.GetUser)
	mux.HandleFunc("GET /users", handler.GetUsers)           // Example: localhost:8080/users?page=1&items=10
	mux.HandleFunc("PUT /users/{uid}", handler.UpdateUser)   // ID cannot be changed. Email can only be changed if is its name is available.
	mux.HandleFunc("GET /users/search", handler.SearchUsers) // localhost:8080/users/search?email=.org
	mux.HandleFunc("PATCH /users/{uid}", handler.PatchUser)

	mux.HandleFunc("POST /register", handler.CreateUser)
	mux.HandleFunc("POST /login", handler.Login)

	mux.HandleFunc("GET /nfc/{uid}", handler.GetNfc)
	mux.HandleFunc("POST /nfc/{uid}", handler.UpdateNFC)

	//medical professionals
	mux.HandleFunc("GET /med/nfc/{med_id}", handler.GetNfcWithMedID)

	mux.HandleFunc("GET /med/users/{uid}", handler.GetMedUser)
	mux.HandleFunc("GET /med/users", handler.GetMedUsers) // Example: localhost:8080/med/users?page=1&items=10
	mux.HandleFunc("PUT /med/users/{uid}", handler.UpdateMedUser)
	mux.HandleFunc("GET /med/users/search", handler.SearchMedUsers) // Example: localhost:8080/med/users/search?email=.org
	mux.HandleFunc("PATCH /med/users/{uid}", handler.PatchMedUser)
	mux.HandleFunc("GET /med/keys", handler.GetMedKeys)

	mux.HandleFunc("POST /med/register", handler.CreateMedUser)
	mux.HandleFunc("POST /med/login", handler.MedLogin)

	mux.HandleFunc("GET /mock/{count}", mock.Users)
	mux.HandleFunc("GET /med/mock/{count}", mock.MedUsers)

	mux.HandleFunc("GET /privacy", handler.Privacy)
	mux.HandleFunc("GET /dataProtection", handler.DataProtection)

	// Run the server
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
