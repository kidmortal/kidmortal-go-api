package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	webhooks "github.com/kidmortal/kidmortal-go-api/src/models/bling/webhooks"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOnePedido Cria um pedido pelo metodo POST
func BlingWebhookHandler(fiber *fiber.Ctx, db *mongo.Client) {
	var webhook webhooks.PedidoStatusAlterado
	fiber.BodyParser(&webhook)
	pedido := webhook.Retorno.Pedidos[0].Pedido
	switch pedido.Situacao {

	case "Concluido":
		pedidoConcluido(fiber, db, pedido)

	case "Aprovado":
		pedidoAprovado(fiber, db, pedido)

	case "Em Andamento":
		pedidoEmAndamento(fiber, db, pedido)

	case "Separado":
		pedidoSeparado(fiber, db, pedido)

	}

}

func pedidoConcluido(fiber *fiber.Ctx, db *mongo.Client, pedido webhooks.PedidoBling) {
	fmt.Println(pedido.Cliente.Nome)
	fmt.Println(pedido.Situacao)
}

func pedidoAprovado(fiber *fiber.Ctx, db *mongo.Client, pedido webhooks.PedidoBling) {
	fmt.Println(pedido.Cliente.Nome)
	fmt.Println(pedido.Situacao)
}
func pedidoEmAndamento(fiber *fiber.Ctx, db *mongo.Client, pedido webhooks.PedidoBling) {
	fmt.Println(pedido.Cliente.Nome)
	fmt.Println(pedido.Situacao)
}
func pedidoSeparado(fiber *fiber.Ctx, db *mongo.Client, pedido webhooks.PedidoBling) {
	fmt.Println(pedido.Cliente.Nome)
	fmt.Println(pedido.Situacao)
}
