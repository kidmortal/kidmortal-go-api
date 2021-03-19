package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	mongodb "github.com/kidmortal/kidmortal-go-api/src/models/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOneNota Cria um Nota pelo metodo POST
func CreateOneNota(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Nota": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindAllNota Busca todos Notas no sistema, aceitando alguns filtros para busca
func FindAllNota(c *fiber.Ctx, db *mongo.Database) error {
	var Nota []mongodb.Nota
	collection := db.Collection("Notas")
	cursor, err := collection.Find(c.Context(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(c.Context(), &Nota); err != nil {
		log.Fatal(err)
	}

	err = c.Status(200).JSON(&fiber.Map{
		"Nota": Nota,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindOneNota Busca um unico Nota usando o numero do Nota como parametro
func FindOneNota(c *fiber.Ctx, db *mongo.Database) error {
	collection := db.Collection("Notas")
	var Nota mongodb.Nota
	NotaParam := c.Params("Nota")
	NotaNumero, err := strconv.Atoi(NotaParam)
	collection.FindOne(c.Context(), bson.M{"numero": NotaNumero}).Decode(&Nota)

	err = c.Status(200).JSON(&fiber.Map{
		"Nota": Nota,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// UpdateOneNota Atualiza um unico Nota usando o numero do Nota como parametro
func UpdateOneNota(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Nota": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// DeleteOneNota Deleta um unico Nota usando o numero do Nota como parametro
func DeleteOneNota(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Nota": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
