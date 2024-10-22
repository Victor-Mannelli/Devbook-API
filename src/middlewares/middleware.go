package middlewares

import (
	"api/src/utils"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Auth Required")
		if err := utils.ValidateToken(r); err != nil {
			utils.HttpErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
