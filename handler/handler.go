package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}
