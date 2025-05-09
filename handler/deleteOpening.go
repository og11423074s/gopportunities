package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Delete opening
// @Description Delete a new opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")

	if !exists || id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id").Error())
		return
	}

	opening := schemas.Opening{}

	// Get opening by id
	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening: %v", err.Error()))
		return
	}

	sendSuccess(ctx, "Delete opening", opening)

}
