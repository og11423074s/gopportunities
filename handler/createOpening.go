package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Create opening
// @Description Create a new opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param opening body CreateOpeningRequest true "Create opening"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("Error binding opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Error binding opening")
		return
	}

	err = request.Validate()
	if err != nil {
		logger.Errorf("Error validating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("Error validating opening: %v", err.Error()))
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening")
		return
	}

	sendSuccess(ctx, "Create opening", opening)
}
