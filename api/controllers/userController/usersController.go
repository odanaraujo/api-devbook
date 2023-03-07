package userController

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-devbook/api/response"
	"github.com/odanaraujo/api-devbook/api/services/userService"
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/authentication"
	"github.com/odanaraujo/api-devbook/infrastructure/security"
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

	user.ID, err = userService.SaveUSer(user)

	if err != nil || user.ID == 0 {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	users, err := userService.GetAll(nameOrNick)

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

	user, err := userService.GetUserID(ID)

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

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != ID {
		response.Erro(w, http.StatusForbidden, errors.New("Unable to update a user other than your me"))
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

	newUser, err := userService.UpdateUser(ID, user)

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

	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if tokenUserID != ID {
		response.Erro(w, http.StatusForbidden, errors.New("Unable to delete a user other than your me"))
		return
	}

	err = userService.DeleteUser(ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followID, err := authentication.ExtractUserID(r)

	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["userId"], 10, 32)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if followID == ID {
		response.Erro(w, http.StatusForbidden, errors.New("Unable to follow yourself"))
		return
	}

	if err := userService.FollowUser(ID, followID); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followID, err := authentication.ExtractUserID(r)

	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if followID == ID {
		response.Erro(w, http.StatusForbidden, errors.New("Unable to unfollow yourself"))
		return
	}

	if err := userService.UnfollowUser(ID, followID); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetFollowersUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	users, err := userService.GetFollowersUser(ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)

}

func GetFollowingUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	users, err := userService.GetFollowingUser(ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	userToken, err := authentication.ExtractUserID(r)

	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["userId"], 10, 64)

	if userToken != ID {
		response.Erro(w, http.StatusForbidden, errors.New("You are not allowed to update third party password"))
		return
	}

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var request domain.PasswordRequest
	if err := json.Unmarshal(body, &request); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	validateOldPassword(w, userToken, request.OldPassword)

	passwordHash, err := security.Hash(request.NewPassword)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := userService.UpdatePassword(userToken, string(passwordHash)); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func validateOldPassword(w http.ResponseWriter, userToken uint64, oldPassword string) {
	password, err := userService.GetPassword(userToken)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(password, oldPassword); err != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("Please enter a valid old password"))
		return
	}
}
