package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func ClientesRouter(r *fiber.App, db *mongo.Database) {
	clientes := r.Group("/clientes")

	clientes.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreateOneCliente(c, db)
	})

	clientes.Get("/", func(c *fiber.Ctx) error {
		return controllers.FindAllCliente(c, db)
	})

	clientes.Get(":cliente", func(c *fiber.Ctx) error {
		return controllers.FindOneCliente(c, db)
	})

	clientes.Put(":cliente", func(c *fiber.Ctx) error {
		return controllers.UpdateOneCliente(c, db)
	})

	clientes.Delete(":cliente", func(c *fiber.Ctx) error {
		return controllers.DeleteOneCliente(c, db)
	})

}
