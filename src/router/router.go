package router

import (
	"github.com/kidmortal/kidmortal-go-api/src/router/routes"

	"github.com/gin-gonic/gin"
)

// Retorna router com rotas configuradas
func ConfigRoutes() *gin.Engine {
	r := gin.Default()
	return routes.Configurar(r)
}
