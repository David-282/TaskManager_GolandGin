package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	gorm.Model
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Status      Status     `json:"status" gorm:"default: null"`
	DueDate     *time.Time `json:"due_date" gorm:"not null"`
	CreatedAt   time.Time  `json:"created_at"`
	UserID      uint       `json:"user_id" gorm:"not null"` // foreign key to Users
}

// dto
type CreateTaskInput struct {
	Title       string     `json:"title"       binding:"required"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date" binding:"required"`
}

type UpdateTaskInput struct {
	Title       string     `json:"title"       binding:"required"`
	Description string     `json:"description"  binding:"required"`
	Status      Status     `json:"status"   binding:"required"`
	DueDate     *time.Time `json:"due_date" binding:"required"`
	UserID      uint       `json:"user_id" binding:"required"`
}
