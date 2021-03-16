package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	webhooks "github.com/kidmortal/kidmortal-go-api/src/models/omie/webhooks"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOnePedido Cria um pedido pelo metodo POST
func OmieWebhookHandler(c *fiber.Ctx, db *mongo.Client) {
	var webhook webhooks.OmieWebhook
	c.BodyParser(&webhook)
	switch webhook.Topic {

	case "Financas.ContaPagar.Incluido":
		financasContaPagarIncluido(c, db)

	case "Financas.ContaPagar.Alterado":
		fmt.Println("Contas a pagar Alterado")

	}

}

func financasContaPagarIncluido(c *fiber.Ctx, db *mongo.Client) {
	var webhook webhooks.ContaAPagarIncluido
	c.BodyParser(&webhook)
	fmt.Println(webhook.Event.Observacao)
}
