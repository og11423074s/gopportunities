package handler

import "github.com/gin-gonic/gin"

func GetOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get openings",
	})
}
