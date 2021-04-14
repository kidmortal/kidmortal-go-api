package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOneLancamento Cria um Lancamento pelo metodo POST
func CreateOneLancamento(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Lancamento": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindAllLancamento Busca todos Lancamentos no sistema, aceitando alguns filtros para busca
func FindAllLancamento(c *fiber.Ctx, db *mongo.Database) error {
	var Lancamento []models.Lancamento
	collection := db.Collection("Lancamentos")
	cursor, err := collection.Find(c.Context(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(c.Context(), &Lancamento); err != nil {
		log.Fatal(err)
	}

	err = c.Status(200).JSON(&fiber.Map{
		"Lancamento": Lancamento,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindOneLancamento Busca um unico Lancamento usando o numero do Lancamento como parametro
func FindOneLancamento(c *fiber.Ctx, db *mongo.Database) error {
	collection := db.Collection("Lancamentos")
	var Lancamento models.Lancamento
	LancamentoParam := c.Params("Lancamento")
	LancamentoNumero, err := strconv.Atoi(LancamentoParam)
	collection.FindOne(c.Context(), bson.M{"numero": LancamentoNumero}).Decode(&Lancamento)

	err = c.Status(200).JSON(&fiber.Map{
		"Lancamento": Lancamento,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// UpdateOneLancamento Atualiza um unico Lancamento usando o numero do Lancamento como parametro
func UpdateOneLancamento(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Lancamento": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// DeleteOneLancamento Deleta um unico Lancamento usando o numero do Lancamento como parametro
func DeleteOneLancamento(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Lancamento": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
