package main

import (
	"fmt"
	"os"

	"github.com/kidmortal/kidmortal-go-api/src/db"
	"github.com/kidmortal/kidmortal-go-api/src/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env File")
	}
	dbSession := db.GetInstance()
	router := router.ConfigRouter(dbSession)
	router.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
