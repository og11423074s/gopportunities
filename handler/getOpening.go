package handler

import "github.com/gin-gonic/gin"

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}
