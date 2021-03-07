package routes

import (
	"net/http"

	"github.com/kidmortal/kidmortal-go-api/src/controllers"
)

var routesPedidos = []Route{
	{
		URI:          "/pedidos",
		Metodo:       http.MethodPost,
		Function:     controllers.CreateOnePedido,
		AuthRequired: true,
	},
	{
		URI:          "/pedidos",
		Metodo:       http.MethodGet,
		Function:     controllers.FindAllPedido,
		AuthRequired: false,
	},
	{
		URI:          "/pedidos/:pedido",
		Metodo:       http.MethodGet,
		Function:     controllers.FindOnePedido,
		AuthRequired: false,
	},
	{
		URI:          "/pedidos/:pedido",
		Metodo:       http.MethodPut,
		Function:     controllers.UpdateOnePedido,
		AuthRequired: true,
	},
	{
		URI:          "/pedidos/:pedido",
		Metodo:       http.MethodDelete,
		Function:     controllers.DeleteOnePedido,
		AuthRequired: true,
	},
}
