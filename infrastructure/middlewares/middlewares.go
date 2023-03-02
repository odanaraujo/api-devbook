package middlewares

import (
	"github.com/odanaraujo/api-devbook/api/response"
	"github.com/odanaraujo/api-devbook/infrastructure/authentication"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			response.Erro(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc(w, r)
	}
}
