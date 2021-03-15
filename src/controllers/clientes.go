package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOneCliente(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Criando Cliente",
	})
}
func FindAllCliente(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Buscando todos Clientes",
	})
}
func FindOneCliente(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("pyramid").Collection("Clientes")
	var Cliente bson.M

	ClienteParam, _ := c.Params.Get("Cliente")
	ClienteNumero, err := strconv.ParseInt(ClienteParam, 10, 64)

	collection.FindOne(ctx, bson.M{"numero": ClienteNumero}).Decode(&Cliente)

	c.JSON(200, gin.H{
		"Cliente": Cliente,
	})
}
func UpdateOneCliente(c *gin.Context) {
	Cliente, _ := c.Params.Get("Cliente")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Atualizando o Cliente %v", Cliente),
	})
}
func DeleteOneCliente(c *gin.Context) {
	Cliente, _ := c.Params.Get("Cliente")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Deletando o Cliente %v", Cliente),
	})
}
