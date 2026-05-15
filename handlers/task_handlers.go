package handlers

import (
	"TaskManager/db"
	"TaskManager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task

	query := db.DB.Model(&models.Task{})
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})

}

func CreateTask(c *gin.Context) {

	var input models.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.MustGet("currentUser").(models.Users)
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		UserID:      user.ID, // link task to user
		Status:      models.StatusPending,
		DueDate:     input.DueDate,
	}

	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func GetTask(c *gin.Context) {
	var task models.Task

	if err := db.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})

}

func DeleteTask(c *gin.Context) {
	var task models.Task

	user := c.MustGet("currentUser").(models.Users)
	id := c.Param("id")

	if err := db.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	db.DB.Delete(&task)
	c.JSON(http.StatusNoContent, nil)
}

func UpdateTask(c *gin.Context) {
	var input models.UpdateTaskInput
	var task models.Task

	id := c.Param("id")

	user := c.MustGet("currentUser").(models.Users)

	if err := db.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not your task"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status
	task.DueDate = input.DueDate

	db.DB.Save(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}
