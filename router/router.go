// inicializa o router utilizando as configuracoes Default do Gin
package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run()

}
