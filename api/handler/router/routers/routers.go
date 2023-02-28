package routers

import (
	"github.com/gorilla/mux"
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
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}
	return r
}
