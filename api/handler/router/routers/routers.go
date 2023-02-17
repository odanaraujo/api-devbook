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

	for _, routerUsers := range routers {
		r.HandleFunc(routerUsers.URI, routerUsers.Function).Methods(routerUsers.Method)
	}
	return r
}
