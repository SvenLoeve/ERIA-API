package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"viseh-api/database/query"
	"viseh-api/services/cryptography/aes"
	"viseh-api/services/cryptography/bcrypt"
	"viseh-api/services/jwt"
	"viseh-api/services/pagination"
	"viseh-api/services/validate"
	"viseh-api/types"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//decode request body to user type
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !validate.UserIsValid(u) {
		BadRequestHandler(w)
		return
	}
	//hash password
	u.Password, err = bcrypt.HashPassword(u.Password)
	// generate encryption key
	u.EncryptionKey = aes.GenerateKey() //todo mock use later
	//Create user database query
	err = query.CreateUser(context.Background(), *u)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("user created"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	//Get ID from request path
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleUser, uid) {
		return
	}
	// Query database for user
	user, err := query.GetUser(context.Background(), uid)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	//Set content type
	w.Header().Set("Content-Type", "application/json")
	//Return user data to api request
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if !jwt.Handler(w, r, jwt.RoleAdmin, 0) {
		return
	}
	pageNumber, items, err := pagination.GetParams(w, r)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	// Query database for all users
	users, err := query.GetUsers(context.Background(), pageNumber, items)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	totalUsers, totalPages, err := query.TotalUsersAndTotalPages(context.Background(), items)
	//Set content type
	w.Header().Set("Content-Type", "application/json")

	w.Header().Add("X-Page", strconv.Itoa(pageNumber))        // The current page index
	w.Header().Add("X-Per-Page", strconv.Itoa(items))         // The total number of results per page
	w.Header().Add("X-Total-Count", strconv.Itoa(totalUsers)) // The total number of records
	w.Header().Add("X-Total-Pages", strconv.Itoa(totalPages)) // The total number of pages

	//Return user data to API request
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleUser, uid) {
		return
	}

	var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	user.ID = uid
	err = query.UpdateUser(context.Background(), user, jwt.IsAdministrator(r))
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("User Updated"))
	return
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	if !jwt.Handler(w, r, jwt.RoleAdmin, 0) {
		return
	}

	var params = r.URL.Query()
	if !params.Has("email") {
		BadRequestHandler(w)
		return
	}
	var email = params.Get("email")
	// Query database to search user by email
	users, err := query.SearchUsers(context.Background(), email)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	totalUsers, err := query.GetTotalUsers(context.Background())
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	//Set content type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-Total-Count", strconv.Itoa(totalUsers))
	//Return user data to API request
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}
func PatchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	if !jwt.Handler(w, r, jwt.RoleUser, uid) {
		return
	}

	var user types.FullUser
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if user.ID != 0 && user.ID != uid { // ID cannot be changed.
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("ID cannot be changed."))
		return
	}
	user.ID = uid
	emailExists, err := query.EmailExists(context.Background(), user.Email)
	if emailExists { // Email is unique, so it needs to be checked.
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Email already exists."))
		return
	}
	if user.Password != "" {
		user.Password, err = bcrypt.HashPassword(user.Password)
		if err != nil {
			InternalServerErrorHandler(w, err)
			return
		}
	}

	err = query.PatchUser(context.Background(), user, jwt.IsAdministrator(r))
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("User Updated"))
}
