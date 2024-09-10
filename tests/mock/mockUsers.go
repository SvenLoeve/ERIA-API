package mock

import (
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"log"
	"net/http"
	"strconv"
)

func Users(w http.ResponseWriter, r *http.Request) {

	//Get id from request path
	count, err := strconv.Atoi(r.PathValue("count"))
	if err != nil {
		return
	}

	for range count {
		postBody, _ := json.Marshal(map[string]any{
			"name":           gofakeit.Name(),
			"email":          gofakeit.Email(),
			"password":       gofakeit.Password(true, true, true, true, true, 16),
			"phone_number":   gofakeit.PhoneFormatted(),
			"encryption_key": gofakeit.BitcoinPrivateKey(),
			"data":           "test",
			"med_id":         gofakeit.DigitN(20),
			"role":           "user",
		})
		responseBody := bytes.NewBuffer(postBody)

		_, err := http.Post("http://localhost:8080/register", "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
	}
}
