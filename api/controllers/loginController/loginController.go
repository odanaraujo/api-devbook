package loginController

import (
	"encoding/json"
	"fmt"
	"github.com/odanaraujo/api-devbook/api/response"
	"github.com/odanaraujo/api-devbook/api/services/loginService"
	"github.com/odanaraujo/api-devbook/infrastructure/authentication"
	"github.com/odanaraujo/api-devbook/infrastructure/security"
	"github.com/odanaraujo/api-devbook/request"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var loginRequest request.LoginRequest

	if err := json.Unmarshal(body, &loginRequest); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := loginRequest.ValidatePassword(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	userWithHash, err := loginService.Login(loginRequest.Email)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(userWithHash.Password, loginRequest.Password); err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userWithHash.ID)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	fmt.Println(token)
}
