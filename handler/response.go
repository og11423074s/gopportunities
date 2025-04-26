package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"errCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s success", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrCode int    `json:"errCode"`
}

type GetAllOpeningResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type OpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
