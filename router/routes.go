package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Get opening",
			})
		})

		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "POST opening",
			})
		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "PUT opening",
			})
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "DELETE opening",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Get openings",
			})
		})
	}

}