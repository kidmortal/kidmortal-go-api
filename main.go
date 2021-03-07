package main

import (
	"log"

	"github.com/kidmortal/kidmortal-go-api/src/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := router.ConfigRoutes()
	router.Run(":2500")

}
