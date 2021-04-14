package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOnePedido Cria um pedido pelo metodo POST
func OmieWebhookHandler(fiber *fiber.Ctx, db *mongo.Database) {
	var webhook models.OmieWebhook
	fiber.BodyParser(&webhook)
	switch webhook.Topic {

	case "ClienteFornecedor.Alterado":
		clienteFornecedorAlterado(fiber, db)

	case "Financas.ContaPagar.Incluido":
		financasContaPagarIncluido(fiber, db)

	case "Financas.ContaPagar.Alterado":
		financasContaPagarAlterado(fiber, db)

	}

}

func clienteFornecedorAlterado(fiber *fiber.Ctx, db *mongo.Database) {
	var webhook models.ClienteFornecedorAlterado
	fiber.BodyParser(&webhook)
	fmt.Println(webhook.Topic)
	fmt.Println(webhook.Event.RazaoSocial)
}

func financasContaPagarIncluido(fiber *fiber.Ctx, db *mongo.Database) {
	var webhook models.ContaAPagarIncluido
	fiber.BodyParser(&webhook)
	fmt.Println(webhook.Topic)
	fmt.Println(webhook.Event.Observacao)
}

func financasContaPagarAlterado(fiber *fiber.Ctx, db *mongo.Database) {
	var webhook models.ContaAPagarAlterado
	fiber.BodyParser(&webhook)
	fmt.Println(webhook.Topic)

	cliente := GetClienteById(349775276)
	fmt.Println(cliente.RazaoSocial)
}
