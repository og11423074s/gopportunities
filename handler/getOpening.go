package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

func GetOpeningHandler(ctx *gin.Context) {

	opening := schemas.Opening{}

	id, exists := ctx.GetQuery("id")

	if !exists || id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id").Error())
		return
	}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
	}

	sendSuccess(ctx, "Get opening", opening)
}
