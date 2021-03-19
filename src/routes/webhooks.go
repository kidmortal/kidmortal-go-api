package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func WebhookRouter(r *fiber.App, db *mongo.Database) {

	webhook := r.Group("/webhook")

	webhook.Post("/omie/pyramid", func(c *fiber.Ctx) error {
		controllers.OmieWebhookHandler(c, db)
		err := c.Status(200).JSON(&fiber.Map{
			"message": "Recebido",
		})
		if err != nil {
			log.Fatal(err)
		}
		return err
	})

	webhook.Post("/bling", func(c *fiber.Ctx) error {
		controllers.BlingWebhookHandler(c, db)
		err := c.Status(200).JSON(&fiber.Map{
			"message": "Recebido",
		})
		if err != nil {
			log.Fatal(err)
		}
		return err
	})

	webhook.Get("/teste", func(c *fiber.Ctx) error {

		err := c.Status(200).JSON(&fiber.Map{
			"message": "Recebido",
		})
		if err != nil {
			log.Fatal(err)
		}
		return err
	})

}
