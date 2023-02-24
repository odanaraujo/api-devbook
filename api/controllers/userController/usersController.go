package userController

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-devbook/api/services"
	"github.com/odanaraujo/api-devbook/domain"
	"io"
	"log"
	"net/http"
	"strconv"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Unable to read request"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	var user domain.User

	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("error when trying to convert user"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	ID, err := services.SaveUSer(user)

	if err != nil || ID == 0 {
		w.Write([]byte("Error when trying to save"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fmt.Sprintf("Usuário inserido com sucesso %d", ID))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := services.GetAll()

	if err != nil {
		w.Write([]byte("Error when trying to get all users"))
		log.Fatalf("GetUsers %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error converting users to json"))
		log.Fatalf("GetUsers %s", err)
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

	user, err := services.GetUserID(ID)

	if err != nil {
		w.Write([]byte("Error when trying to get user"))
		log.Fatalf("GetUser %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error converting users to json "))
		log.Fatalf("GetUser %s", err)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting id to uint"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Unable to read request"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	var user domain.User

	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("error when trying to convert user"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	newUser, err := services.UpdateUser(ID, user)

	if err != nil {
		w.Write([]byte("Error when trying to update user"))
		log.Fatalf("UpdateUser %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		w.Write([]byte("Error converting users to json "))
		log.Fatalf("GetUser %s", err)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting id to uint"))
		log.Fatalf("SaveUser %s", err)
		return
	}

	err = services.DeleteUser(ID)

	w.WriteHeader(http.StatusNoContent)
}
