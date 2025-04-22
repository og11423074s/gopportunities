package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Create opening",
	})
}

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Update opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete opening",
	})
}

func GetOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get openings",
	})
}
