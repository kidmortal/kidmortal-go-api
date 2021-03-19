package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	mongodb "github.com/kidmortal/kidmortal-go-api/src/models/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOnePedido Cria um pedido pelo metodo POST
func CreateOnePedido(c *fiber.Ctx, db *mongo.Database) error {
	var pedido mongodb.Pedido
	validate := validator.New()
	err := validate.Struct(pedido)

	if err != nil {
		err = c.Status(200).JSON(&fiber.Map{
			"pedido": "Dados incompletos",
		})
		return err
	}
	result, err := db.Collection("pedidos").InsertOne(context.TODO(), pedido)

	if err != nil {
		log.Fatal(err)
	}

	err = c.Status(200).JSON(&fiber.Map{
		"pedido": result,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindAllPedido Busca todos pedidos no sistema, aceitando alguns filtros para busca
func FindAllPedido(c *fiber.Ctx, db *mongo.Database) error {
	var pedido []mongodb.Pedido
	collection := db.Collection("pedidos")
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
func FindOnePedido(c *fiber.Ctx, db *mongo.Database) error {
	collection := db.Collection("pedidos")
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
func UpdateOnePedido(c *fiber.Ctx, db *mongo.Database) error {
	numero, err := strconv.Atoi(c.Params("pedido"))
	if err != nil {
		err = c.Status(200).JSON(&fiber.Map{
			"pedido": "Pedido invalido",
		})
		return err
	}

	var pedido mongodb.Pedido
	c.BodyParser(&pedido)
	validate := validator.New()
	err = validate.Struct(pedido)

	if err != nil {
		fmt.Println(err)
		err = c.Status(200).JSON(&fiber.Map{
			"pedido": "Dados incompletos",
		})

		return err
	}

	result := db.Collection("pedidos").FindOneAndReplace(context.TODO(), mongodb.Pedido{Numero: numero}, pedido)

	err = c.Status(200).JSON(&fiber.Map{
		"pedido": result,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// DeleteOnePedido Deleta um unico pedido usando o numero do pedido como parametro
func DeleteOnePedido(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
