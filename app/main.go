package main

import (
	"fmt"
	"github.com/odanaraujo/api-devbook/api/handler/router/routers"
	"github.com/odanaraujo/api-devbook/infrastructure/config"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := routers.RouterConfig()
	fmt.Printf("Connected on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
