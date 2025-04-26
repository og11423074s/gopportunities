package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Update opening
// @Description Update a opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Param opening body UpdateOpeningRequest true "Update opening"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [patch]
func UpdateOpeningHandler(ctx *gin.Context) {

	request := UpdateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Error binding opening")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, exists := ctx.GetQuery("id")

	if !exists || id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	sendSuccess(ctx, "Update opening", opening)
}
