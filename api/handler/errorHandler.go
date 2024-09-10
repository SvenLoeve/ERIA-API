package handler

import (
	"net/http"
)

func InternalServerErrorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError) //set http status
	_, err = w.Write([]byte(err.Error()))         //set response message
	if err != nil {
		return
	}
	return
}

func NotFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)         //set http status
	_, err := w.Write([]byte("404 Not Found")) //set response message
	if err != nil {
		return
	}
	return
}

func BadRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)         //set http status
	_, err := w.Write([]byte("400 Bad Request")) //set response message
	if err != nil {
		return
	}
	return
}

func UnprocessableEntityHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity) //set http status
	_, err = w.Write([]byte(err.Error()))         //set response message
	if err != nil {
		return
	}
	return
}

func UnauthorizedHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)    //set http status
	_, err := w.Write([]byte("Unauthorized")) //set response message
	if err != nil {
		return
	}
	return
}
