package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router
	router := gin.Default()

	// initialize routes
	initializeRoutes(router)

	// run router
	router.Run(":8080")
}
