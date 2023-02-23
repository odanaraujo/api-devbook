package userController

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/database"
	"github.com/odanaraujo/api-devbook/infrastructure/repository"
	"io"
	"net/http"
	"strconv"
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
	var users []domain.Users

	db, err := database.Connection()

	if err != nil {
		w.Write([]byte("error connecting to database"))
		return
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	users, err = repo.GetAll(users)

	if err != nil {
		w.Write([]byte("Error when trying to get all users"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error converting users to json"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting id to uint"))
		return
	}

	db, err := database.Connection()

	if err != nil {
		w.Write([]byte("error connecting to database"))
		return
	}

	defer db.Close()

	repo := repository.NewRepositoryUser(db)
	user, err := repo.GetUserId(ID)

	if err != nil {
		w.Write([]byte("Error when trying to get all users"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error converting users to json"))
		return
	}
}
