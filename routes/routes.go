package routes

import (
	"TaskManager/handlers"
	"TaskManager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}

	    // Protected routes — RequireAuth middleware runs first
    tasks := api.Group("/tasks")
    tasks.Use(middleware.RequireAuth)
		{
			tasks.GET("get_all_tasks", handlers.GetTasks)
			tasks.POST("", handlers.CreateTask)
			tasks.GET("get_task/:id", handlers.GetTask)
			tasks.PUT("update_task/:id", handlers.UpdateTask)
			tasks.DELETE("delete_task/:id", handlers.DeleteTask)
			//tasks.GET("get_by_status", handlers.FilterByStatus)

		}
	
}
