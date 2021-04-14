package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigRouter retorna router com rotas configuradas
func ConfigRouter(db *mongo.Database) *fiber.App {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: false,
	}))
	ClientesRouter(router, db)
	PedidosRouter(router, db)
	NotasRouter(router, db)
	WebhookRouter(router, db)
	WebsocketManager(router, db)
	UtilsRouter(router, db)
	return router
}
