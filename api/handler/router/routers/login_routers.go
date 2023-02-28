package routers

import (
	"github.com/odanaraujo/api-devbook/api/controllers/loginController"
	"net/http"
)

var routerLogin = Router{
	URI:              "/login",
	Method:           http.MethodPost,
	Function:         loginController.Login,
	IsAuthentication: false,
}
