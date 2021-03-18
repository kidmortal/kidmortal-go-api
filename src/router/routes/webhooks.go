package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	models "github.com/kidmortal/kidmortal-go-api/src/models/postgres"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func WebhookRouter(r *fiber.App, db *mongo.Client, dbp *gorm.DB) {

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

		newPedido := models.Pedido{
			Numero:   500,
			Nome:     "MATCHELOOO",
			Cnpj:     "29238232323",
			Valor:    599.25,
			Vendedor: "mainha",
			Cidade:   "Caju",
			Estado:   "CU",
			Status:   "sein",
		}
		dbp.Table("Pedidos").Create(&newPedido)
		fmt.Println(newPedido.ID)
		return err
	})

}
