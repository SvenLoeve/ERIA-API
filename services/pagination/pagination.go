package pagination

import (
	"errors"
	"net/http"
	"strconv"
)

// GetParams Returns query parameters from item and page.
// Example: localhost:8080/users?page=1&items=10 returns -> 1, 10
func GetParams(w http.ResponseWriter, r *http.Request) (int, int, error) {
	var params = r.URL.Query()
	if params.Get("items") == "" || params.Get("page") == "" {
		return 0, 0, errors.New("required params: items and page not found")
	}

	var pageNumber, err = strconv.ParseInt(params.Get("page"), 10, 32)
	if err != nil {
		return 0, 0, errors.New("required params: parsing page went wrong")
	}

	var items, err2 = strconv.ParseInt(params.Get("items"), 10, 32)
	if err2 != nil {
		return 0, 0, errors.New("required params: parsing items went wrong")
	}

	return int(pageNumber), int(items), nil
}

func badRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)         //set http status
	_, err := w.Write([]byte("400 Bad Request")) //set response message
	if err != nil {
		return
	}
	return
}
