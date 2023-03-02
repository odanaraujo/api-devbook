package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	StringConnectionDatabase = ""
	Port                     = 0
	SecretKey                []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error when running godotenv.load %s", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	StringConnectionDatabase = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}
