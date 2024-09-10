package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"viseh-api/database/query"
	"viseh-api/services/cryptography/bcrypt"
	"viseh-api/services/jwt"
	"viseh-api/services/pagination"
	"viseh-api/services/validate"
	"viseh-api/types"
)

func CreateMedUser(w http.ResponseWriter, r *http.Request) {
	//decode request body to user type
	err := json.NewDecoder(r.Body).Decode(&mu)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !validate.MedUserIsValid(mu) {
		BadRequestHandler(w)
		return
	}
	//hash password
	mu.Password, err = bcrypt.HashPassword(mu.Password)
	//Create user database query
	err = query.CreateMedUser(context.Background(), *mu)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("user created"))
		return
	}
}

func GetMedUser(w http.ResponseWriter, r *http.Request) {
	//Get id from request path
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleAmbulance, uid) {
		return
	}
	//query database for user
	medUser, err := query.GetMedUser(context.Background(), uid)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	//Set content type
	w.Header().Set("Content-Type", "application/json")
	//Return user data to api request
	err = json.NewEncoder(w).Encode(medUser)
	if err != nil {
		return
	}
}

func GetMedUsers(w http.ResponseWriter, r *http.Request) {
	if !jwt.Handler(w, r, jwt.RoleAmbulance, 0) {
		return
	}
	pageNumber, items, err := pagination.GetParams(w, r)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	// Query database for all users
	users, err := query.GetMedUsers(context.Background(), pageNumber, items)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	totalUsers, totalPages, err := query.TotalMedUsersAndTotalPages(context.Background(), items)
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
func UpdateMedUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleDoctor, uid) {
		return
	}

	var user types.MedUser
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	user.ID = uid
	err = query.UpdateMedUser(context.Background(), user)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Med User Updated"))
	return
}
func SearchMedUsers(w http.ResponseWriter, r *http.Request) {
	if !jwt.Handler(w, r, jwt.RoleAmbulance, 0) {
		return
	}
	var params = r.URL.Query()
	if !params.Has("email") {
		BadRequestHandler(w)
		return
	}
	var email = params.Get("email")
	// Query database to search user by email
	users, err := query.SearchMedUsers(context.Background(), email)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	totalUsers, err := query.GetTotalMedUsers(context.Background())

	if err != nil {
		InternalServerErrorHandler(w, err)
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
func PatchMedUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleDoctor, uid) {
		return
	}

	var user types.FullMedUser
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
	emailExists, err := query.MedEmailExists(context.Background(), user.Email)
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

	err = query.PatchMedUser(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		InternalServerErrorHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Med User Updated"))
}

func GetMedKeys(w http.ResponseWriter, r *http.Request) {
	if !jwt.Handler(w, r, jwt.RoleAmbulance, 0) {
		return
	}
	//query database for user
	keys, err := query.GetAllKeys(context.Background())
	if err != nil {
		return
	}
	//Set content type
	w.Header().Set("Content-Type", "application/json")
	//Return user data to api request
	err = json.NewEncoder(w).Encode(keys)
	if err != nil {
		return
	}
}
