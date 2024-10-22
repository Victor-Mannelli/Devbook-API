package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HttpJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func HttpErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	// creates a json object that has the string Error as value of the key "error" -> { error: Error }
	HttpJsonResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

func HashString(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}

func ValidateHash(hashedStr, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str))
}
