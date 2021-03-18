package controllers

import (
	requests "github.com/kidmortal/kidmortal-go-api/src/models/omie/requests"
	"github.com/kidmortal/kidmortal-go-api/src/utils"
)

// GetClienteById OMIE busca cliente pelo ID, e Retorna os dados
func GetClienteById(ID int) requests.ClienteOmie {
	cliente := utils.OmieGetClienteById(ID)
	return cliente
}
