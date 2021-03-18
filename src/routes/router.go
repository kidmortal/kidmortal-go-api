package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigRouter retorna router com rotas configuradas
func ConfigRouter(db *mongo.Client) *fiber.App {
	router := fiber.New()
	ClientesRouter(router, db)
	PedidosRouter(router, db)
	WebhookRouter(router, db)
	return router
}
