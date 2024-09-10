package handler

import (
	"net/http"
	"viseh-api/database/ent"
)

var u *ent.User
var mu *ent.MedUser

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
	_, err := w.Write([]byte("Bad endpoint"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
}
