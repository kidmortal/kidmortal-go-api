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

func CreateOnePedido(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Criando Pedido",
	})
}
func FindAllPedido(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Buscando todos Pedidos",
	})
}
func FindOnePedido(c *gin.Context) {
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
	collection := client.Database("pyramid").Collection("pedidos")
	var pedido bson.M

	pedidoParam, _ := c.Params.Get("pedido")
	pedidoNumero, err := strconv.ParseInt(pedidoParam, 10, 64)

	collection.FindOne(ctx, bson.M{"numero": pedidoNumero}).Decode(&pedido)

	c.JSON(200, gin.H{
		"pedido": pedido,
	})
}
func UpdateOnePedido(c *gin.Context) {
	pedido, _ := c.Params.Get("pedido")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Atualizando o Pedido %v", pedido),
	})
}
func DeleteOnePedido(c *gin.Context) {
	pedido, _ := c.Params.Get("pedido")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Deletando o Pedido %v", pedido),
	})
}
