package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func NotasRouter(r *fiber.App, db *mongo.Database) {

	notas := r.Group("/notas")

	notas.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreateOneNota(c, db)
	})

	notas.Get("/", func(c *fiber.Ctx) error {
		return controllers.FindAllNota(c, db)
	})

	notas.Get(":nota", func(c *fiber.Ctx) error {
		return controllers.FindOneNota(c, db)
	})

	notas.Put(":nota", func(c *fiber.Ctx) error {
		return controllers.UpdateOneNota(c, db)
	})

	notas.Delete(":nota", func(c *fiber.Ctx) error {
		return controllers.DeleteOneNota(c, db)
	})

}
