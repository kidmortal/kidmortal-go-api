package main

import (
	"log"

	"github.com/kidmortal/kidmortal-go-api/src/db"
	"github.com/kidmortal/kidmortal-go-api/src/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	dbSession := db.GetInstance()
	router := router.ConfigRouter(dbSession)
	router.Listen(":3000")

}
