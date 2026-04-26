package main

import (
	"TaskManager/routes"

	"github.com/gin-gonic/gin"
)

// http.StatusOk = 200
func main() {
	server := gin.Default()
	routes.SetupRoutes(server)
	//server.GET("/pinging", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello World",
	//	})
	//})

	server.Run(":8080")
}
