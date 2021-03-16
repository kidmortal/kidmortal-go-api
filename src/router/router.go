package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/router/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigRouter retorna router com rotas configuradas
func ConfigRouter(db *mongo.Client) *fiber.App {
	router := fiber.New()
	routes.ClientesRouter(router, db)
	routes.PedidosRouter(router, db)
	routes.OmieWebhookRouter(router, db)
	return router
}
