package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateOneCliente(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
func FindAllCliente(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
func FindOneCliente(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
func UpdateOneCliente(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
func DeleteOneCliente(c *fiber.Ctx, db *mongo.Database) error {
	err := c.Status(200).JSON(&fiber.Map{
		"pedido": "eai",
	})

	if err != nil {
		log.Fatal(err)
	}

	return err
}
