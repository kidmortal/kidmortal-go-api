package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pedido struct {
	numero  int32
	cliente string
	valor   float32
}

func route1(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
func route2(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "ping pong",
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	client, err := mongo.NewClient(options.Client().ApplyURI("<ATLAS_URI_HERE>"))
	router.GET("/ping", route1)
	router.GET("/pong", route2)
	router.Run(":2500") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
