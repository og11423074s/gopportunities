package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router
	router := gin.Default()

	// initialize routes
	initializeRoutes(router)

	// run router
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
