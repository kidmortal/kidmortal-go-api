package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kidmortal/kidmortal-go-api/src/db"
	"github.com/kidmortal/kidmortal-go-api/src/models"
	"github.com/kidmortal/kidmortal-go-api/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetClienteById OMIE busca cliente pelo ID, e Retorna os dados
func GetClienteById(ID int) models.ClienteOmie {
	cliente := utils.OmieGetClienteByNcodcli(ID)
	return cliente
}

func GetNfSaidaByPeriod(c *fiber.Ctx, database *mongo.Database) error {
	var request models.NfParam
	c.BodyParser(&request)
	notas := utils.OmieGetNfSaidaByPeriod(request.Demiinicial, request.Demifinal)

	for _, s := range notas {
		db.InsertOrUpdateNota(s, database)
	}

	err := c.Status(200).JSON(&fiber.Map{
		"pedidos": notas,
	})

	return err
}
