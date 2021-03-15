package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

// Route representa todas rotas da api
type Route struct {
	URI          string
	Metodo       string
	Function     func(c *gin.Context)
	AuthRequired bool
}

// Routes representa um array de route
type Routes []Route

// ConfigRoutes faz append de todas rotas no router
func ConfigRoutes(r *fiber.App, db *mongo.Client) *fiber.App {

	r.Get("/pedidos", func(c *fiber.Ctx) error {
		return controllers.FindAllPedido(c, db)
	})

	r.Get("/pedidos/:pedido", func(c *fiber.Ctx) error {
		return controllers.FindOnePedido(c, db)
	})

	return r
}
