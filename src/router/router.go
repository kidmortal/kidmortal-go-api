package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/router/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// ConfigRouter retorna router com rotas configuradas
func ConfigRouter(db *mongo.Client, dbp *gorm.DB) *fiber.App {
	router := fiber.New()
	routes.ClientesRouter(router, db)
	routes.PedidosRouter(router, db)
	routes.WebhookRouter(router, db, dbp)
	return router
}
