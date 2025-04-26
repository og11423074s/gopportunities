package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Get by ID opening
// @Description GET by ID opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
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
