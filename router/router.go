package router

import "github.com/gin-gonic/gin"

// Setup and run the server
func Initialize() {
	router := gin.Default()
	
	initializeRoutes(router)

	router.Run(":8080")
}
