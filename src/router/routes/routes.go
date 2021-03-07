package routes

import "github.com/gin-gonic/gin"

// Route representa todas rotas da api
type Route struct {
	URI          string
	Metodo       string
	Function     func(c *gin.Context)
	AuthRequired bool
}

// Configurar faz a conexao de todas rotas ao router
func Configurar(r *gin.Engine) *gin.Engine {
	rotas := routesPedidos

	for _, route := range rotas {
		r.Handle(route.Metodo, route.URI, route.Function)
	}

	return r
}
