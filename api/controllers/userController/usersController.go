package userController

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-devbook/api/response"
	"github.com/odanaraujo/api-devbook/api/services"
	"github.com/odanaraujo/api-devbook/domain"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user domain.User

	if err := json.Unmarshal(body, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(false); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	user.ID, err = services.SaveUSer(user)

	if err != nil || user.ID == 0 {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	users, err := services.GetAll(nameOrNick)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	user, err := services.GetUserID(ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		response.Erro(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user domain.User

	if err := json.Unmarshal(body, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(true); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := services.UpdateUser(ID, user)

	if newUser.ID == 0 {
		response.Erro(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = services.DeleteUser(ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusNoContent, nil)
}
