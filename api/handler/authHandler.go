package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"viseh-api/database/query"
	"viseh-api/services/cryptography/bcrypt"
	"viseh-api/services/jwt"
	"viseh-api/services/validate"
	"viseh-api/types"
)

var l *types.UserLogin

func Login(w http.ResponseWriter, r *http.Request) {
	//decode request body to login type
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !validate.LoginIsValid(l) {
		BadRequestHandler(w)
		return
	}
	//query database for login information
	u, err := query.GetLoginUser(context.Background(), l.Email)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	//check password hash, create jwt token if correct and set in auth header
	if bcrypt.CheckPasswordHash(l.Password, u.Password) { //check password hash
		jwtToken := jwt.CreateJWT(u.Name, u.ID, string(u.Role)) //create jwt token
		if jwtToken != "" {
			w.Header().Set("Authorization", jwtToken)
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte("Login successful!"))
			return
		} else {
			InternalServerErrorHandler(w, err)
			return
		}
	} else {
		UnauthorizedHandler(w)
		return
	}
}

func MedLogin(w http.ResponseWriter, r *http.Request) {
	//decode request body to login type
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !validate.LoginIsValid(l) {
		BadRequestHandler(w)
		return
	}
	//query database for login information
	mu, err := query.GetLoginMedUser(context.Background(), l.Email)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	//check password hash, create jwt token if correct and set in auth header
	if bcrypt.CheckPasswordHash(l.Password, mu.Password) { //check password hash
		jwtToken := jwt.CreateJWT(mu.Name, mu.ID, string(mu.Role)) //create jwt token
		if jwtToken != "" {
			w.Header().Set("Authorization", jwtToken)
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte("login success"))
			return
		} else {
			InternalServerErrorHandler(w, err)
			return
		}
	} else {
		UnauthorizedHandler(w)
		return
	}
}
