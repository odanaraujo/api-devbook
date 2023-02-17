package routers

import (
	"github.com/odanaraujo/api-devbook/api/controllers/userController"
	"net/http"
)

var RouterUsers = []Router{
	{
		URI:              "/users",
		Method:           http.MethodPost,
		Function:         userController.SaveUser,
		IsAuthentication: false,
	},
	{
		URI:              "/users",
		Method:           http.MethodGet,
		Function:         userController.GetUsers,
		IsAuthentication: false,
	},
	{
		URI:              "/user/{i}",
		Method:           http.MethodGet,
		Function:         userController.GetUser,
		IsAuthentication: false,
	},
}
