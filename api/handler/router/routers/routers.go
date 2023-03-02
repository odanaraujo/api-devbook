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
	routers := RouterUsers
	routers = append(routers, routerLogin)
	for _, router := range routers {
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
