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
}
