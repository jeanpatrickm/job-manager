package router

import (
	"github.com/gin-gonic/gin"
)

// Importar o arquivo de rotas
func InitializeRouter() {
	
	// Inicializar router
	router := gin.Default()

	// inicializar rotas
	InitializeRoutes(router)

	// run server
	router.Run(":8080")
}
