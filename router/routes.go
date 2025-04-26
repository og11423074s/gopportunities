package router

import (
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/handler"

	docs "github.com/og11423074s/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	// Initialize handler
	handler.InitHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Description = "Management Openings"
	docs.SwaggerInfo.Version = "1.0.0"

	v1 := router.Group(basePath)

	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.GetOpeningHandler)

		v1.PATCH("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.GetOpeningsHandler)

	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
