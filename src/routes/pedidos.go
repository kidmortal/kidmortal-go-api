package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func PedidosRouter(r *fiber.App, db *mongo.Database) {

	pedidos := r.Group("/pedidos")

	pedidos.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreateOnePedido(c, db)
	})

	pedidos.Get("/", func(c *fiber.Ctx) error {
		return controllers.FindAllPedido(c, db)
	})

	pedidos.Get(":pedido", func(c *fiber.Ctx) error {
		return controllers.FindOnePedido(c, db)
	})

	pedidos.Put(":pedido", func(c *fiber.Ctx) error {
		return controllers.UpdateOnePedido(c, db)
	})

	pedidos.Delete(":pedido", func(c *fiber.Ctx) error {
		return controllers.DeleteOnePedido(c, db)
	})

}
