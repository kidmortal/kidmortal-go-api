package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigRoutes faz append de todas rotas no router
func ConfigRoutes(r *fiber.App, db *mongo.Client) *fiber.App {

	r.Post("/pedidos", func(c *fiber.Ctx) error {
		return controllers.CreateOnePedido(c, db)
	})

	r.Get("/pedidos", func(c *fiber.Ctx) error {
		return controllers.FindAllPedido(c, db)
	})

	r.Get("/pedidos/:pedido", func(c *fiber.Ctx) error {
		return controllers.FindOnePedido(c, db)
	})

	r.Put("/pedidos/:pedido", func(c *fiber.Ctx) error {
		return controllers.UpdateOnePedido(c, db)
	})

	r.Delete("/pedidos/:pedido", func(c *fiber.Ctx) error {
		return controllers.DeleteOnePedido(c, db)
	})

	r.Post("/clientes", func(c *fiber.Ctx) error {
		return controllers.CreateOneCliente(c, db)
	})

	r.Get("/clientes", func(c *fiber.Ctx) error {
		return controllers.FindAllCliente(c, db)
	})

	r.Get("/clientes/:cliente", func(c *fiber.Ctx) error {
		return controllers.FindOneCliente(c, db)
	})

	r.Put("/clientes/:cliente", func(c *fiber.Ctx) error {
		return controllers.UpdateOneCliente(c, db)
	})

	r.Delete("/clientes/:cliente", func(c *fiber.Ctx) error {
		return controllers.DeleteOneCliente(c, db)
	})

	return r
}
