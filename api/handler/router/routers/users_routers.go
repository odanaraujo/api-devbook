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
		IsAuthentication: true,
	},
	{
		URI:              "/user/{id}",
		Method:           http.MethodGet,
		Function:         userController.GetUser,
		IsAuthentication: true,
	},
	{
		URI:              "/user/{id}",
		Method:           http.MethodPut,
		Function:         userController.UpdateUser,
		IsAuthentication: true,
	},
	{
		URI:              "/user/{id}",
		Method:           http.MethodDelete,
		Function:         userController.DeleteUser,
		IsAuthentication: true,
	},
}
