package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/router/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigRouter retorna router com rotas configuradas
func ConfigRouter(client *mongo.Client) *fiber.App {
	r := fiber.New()
	return routes.ConfigRoutes(r, client)
}
