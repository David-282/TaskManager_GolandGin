// @title Task Manager API
// @version 1.0
// @description Task Manager backend API
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"TaskManager/config"
	"TaskManager/db"
	"TaskManager/routes"
	"fmt"
		_ "TaskManager/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	godotenv.Load()
	

	cfg := config.Load()

	db.Connect(cfg)
	fmt.Println("USER:", cfg.DBUser)
	fmt.Println("DB:", cfg.DBName)

	fmt.Println("DB nil?", db.DB == nil)

	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(server)

	server.Run(":8080")
}