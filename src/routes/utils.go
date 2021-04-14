package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func UtilsRouter(r *fiber.App, db *mongo.Database) {
	utils := r.Group("/utils")

	utils.Post("/listarnf", func(c *fiber.Ctx) error {
		return controllers.GetNfSaidaByPeriod(c, db)
	})

}
