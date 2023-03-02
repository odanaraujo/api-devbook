package main

import (
	"fmt"
	"github.com/odanaraujo/api-devbook/api/handler/router/routers"
	"github.com/odanaraujo/api-devbook/infrastructure/config"
	"log"
	"net/http"
)

//GERAR A KEY ALEATÓRIA
//func init() {
//	key := make([]byte, 64)
//	if _, err := rand.Read(key); err != nil {
//		log.Fatal(err)
//	}
//
//	stringBase64 := base64.StdEncoding.EncodeToString(key)
//	fmt.Printf(stringBase64)
//}

func main() {
	config.Carregar()
	r := routers.RouterConfig()
	fmt.Printf("Connected on port - %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
