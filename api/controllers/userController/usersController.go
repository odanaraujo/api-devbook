package userController

import (
	"encoding/json"
	"fmt"
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/database"
	"github.com/odanaraujo/api-devbook/infrastructure/repository"
	"io"
	"net/http"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Unable to read request"))
		return
	}

	var user domain.Users

	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("error when trying to convert user"))
		return
	}

	db, err := database.Connection()

	if err != nil {
		w.Write([]byte("error connecting to database"))
		return
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	userId, err := repo.Save(user)

	if err != nil || userId == 0 {
		w.Write([]byte("Error when trying to save"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fmt.Sprintf("Usuário inserido com sucesso %d", userId))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
