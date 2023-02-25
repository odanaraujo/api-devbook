package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func Erro(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, struct {
		Error string `json:"message"`
	}{
		Error: err.Error(),
	})
}
