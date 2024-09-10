package handler

import (
	"io"
	"net/http"
	"os"
)

func Privacy(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/Privacy_Verklaring_Final.pdf")
	if err != nil {
		InternalServerErrorHandler(w, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	w.Header().Set("Content-Type", "application/pdf")

	if _, err := io.Copy(w, f); err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
}
func DataProtection(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile("static/dataProtection.html")
	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(dat)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
}
