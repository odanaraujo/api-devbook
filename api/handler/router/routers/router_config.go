package routers

import (
	"github.com/gorilla/mux"
)

func RouterConfig() *mux.Router {
	r := mux.NewRouter()
	return Config(r)
}
