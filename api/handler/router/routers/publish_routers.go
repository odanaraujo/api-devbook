package routers

import (
	"github.com/odanaraujo/api-devbook/api/controllers/publishController"
	"net/http"
)

var routerPublish = []Router{
	{
		URI:              "/publish",
		Method:           http.MethodPost,
		Function:         publishController.CreaterPublish,
		IsAuthentication: true,
	},
	{
		URI:              "/publish/{publishID}",
		Method:           http.MethodGet,
		Function:         publishController.GetPublish,
		IsAuthentication: true,
	},
}
