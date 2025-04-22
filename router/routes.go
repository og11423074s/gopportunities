package router

import (
	"github.com/gin-gonic/gin"
	"github.com/og11423074s/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening/:id", handler.GetOpeningHandler)

		v1.PATCH("/opening/:id", handler.UpdateOpeningHandler)

		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.GetOpeningsHandler)

	}
}
