package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	mongodb "github.com/kidmortal/kidmortal-go-api/src/models/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOnePedido Cria um pedido pelo metodo POST
func CreateOnePedido(c *fiber.Ctx, db *mongo.Client) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindAllPedido Busca todos pedidos no sistema, aceitando alguns filtros para busca
func FindAllPedido(c *fiber.Ctx, db *mongo.Client) error {
	var pedido []mongodb.Pedido
	collection := db.Database("pyramid").Collection("pedidos")
	cursor, err := collection.Find(c.Context(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(c.Context(), &pedido); err != nil {
		log.Fatal(err)
	}

	err = c.Status(200).JSON(&fiber.Map{
		"pedido": pedido,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindOnePedido Busca um unico pedido usando o numero do pedido como parametro
func FindOnePedido(c *fiber.Ctx, db *mongo.Client) error {
	collection := db.Database("pyramid").Collection("pedidos")
	var pedido mongodb.Pedido
	pedidoParam := c.Params("pedido")
	pedidoNumero, err := strconv.Atoi(pedidoParam)
	collection.FindOne(c.Context(), bson.M{"numero": pedidoNumero}).Decode(&pedido)

	err = c.Status(200).JSON(&fiber.Map{
		"pedido": pedido,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// UpdateOnePedido Atualiza um unico pedido usando o numero do pedido como parametro
func UpdateOnePedido(c *fiber.Ctx, db *mongo.Client) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// DeleteOnePedido Deleta um unico pedido usando o numero do pedido como parametro
func DeleteOnePedido(c *fiber.Ctx, db *mongo.Client) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
