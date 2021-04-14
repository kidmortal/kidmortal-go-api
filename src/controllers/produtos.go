package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOneProduto Cria um Produto pelo metodo POST
func CreateOneProduto(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Produto": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindAllProduto Busca todos Produtos no sistema, aceitando alguns filtros para busca
func FindAllProduto(c *fiber.Ctx, db *mongo.Database) error {
	var Produto []models.Produto
	collection := db.Collection("Produtos")
	cursor, err := collection.Find(c.Context(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(c.Context(), &Produto); err != nil {
		log.Fatal(err)
	}

	err = c.Status(200).JSON(&fiber.Map{
		"Produto": Produto,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// FindOneProduto Busca um unico Produto usando o numero do Produto como parametro
func FindOneProduto(c *fiber.Ctx, db *mongo.Database) error {
	collection := db.Collection("Produtos")
	var Produto models.Produto
	ProdutoParam := c.Params("Produto")
	ProdutoNumero, err := strconv.Atoi(ProdutoParam)
	collection.FindOne(c.Context(), bson.M{"numero": ProdutoNumero}).Decode(&Produto)

	err = c.Status(200).JSON(&fiber.Map{
		"Produto": Produto,
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// UpdateOneProduto Atualiza um unico Produto usando o numero do Produto como parametro
func UpdateOneProduto(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Produto": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}

// DeleteOneProduto Deleta um unico Produto usando o numero do Produto como parametro
func DeleteOneProduto(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"Produto": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
