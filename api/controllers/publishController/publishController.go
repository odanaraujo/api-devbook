package publishController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-devbook/api/response"
	"github.com/odanaraujo/api-devbook/api/services/publishService"
	"github.com/odanaraujo/api-devbook/domain"
	"github.com/odanaraujo/api-devbook/infrastructure/authentication"
	"io"
	"net/http"
	"strconv"
)

func CreaterPublish(w http.ResponseWriter, r *http.Request) {
	tokenID, err := authentication.ExtractUserID(r)

	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	var publish domain.Publish
	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal(body, &publish); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := publish.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	publish.AuthorID = tokenID

	publish.ID, err = publishService.CreaterPublish(publish)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publish)
}

func GetPublish(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["publishId"], 10, 64)

	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	publish, err := publishService.GetPublish(ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publish)
}
