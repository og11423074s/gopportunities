package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Get All openings
// @Description GET All openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} GetAllOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error getting openings")
		return
	}

	sendSuccess(ctx, "Get openings", openings)
}
