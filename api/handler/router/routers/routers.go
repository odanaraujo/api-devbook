package routers

import (
	"github.com/gorilla/mux"
	"github.com/odanaraujo/api-devbook/infrastructure/middlewares"
	"net/http"
)

type Router struct {
	URI              string
	Method           string
	Function         func(w http.ResponseWriter, r *http.Request)
	IsAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	routersUsers := RouterUsers
	routersUsers = append(routersUsers, routerLogin)
	routersUsers = append(routersUsers, routerPublish...)
	for _, router := range routersUsers {
		if router.IsAuthentication {
			r.HandleFunc(router.URI,
				middlewares.Logger(middlewares.Authenticate(router.Function)),
			).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI, middlewares.Logger(router.Function)).Methods(router.Method)
		}
	}
	return r
}
