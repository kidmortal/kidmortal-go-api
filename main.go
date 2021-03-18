package main

import (
	"fmt"
	"os"

	"github.com/kidmortal/kidmortal-go-api/src/db"
	"github.com/kidmortal/kidmortal-go-api/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env File")
	}

	dbSession := db.GetInstance()
	router := routes.ConfigRouter(dbSession)
	router.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
