package routes

import (
	"TaskManager/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("get_all_tasks", handlers.GetTasks)
			tasks.POST("", handlers.CreateTask)
			tasks.GET("get_task/:id", handlers.GetTask)
			tasks.PUT("update_task/:id", handlers.UpdateTask)
			tasks.DELETE("delete_task/:id", handlers.DeleteTask)
			tasks.GET("get_by_status", handlers.FilterByStatus)

			//	GET /get_by_status?status=pending
		}
	}
}
