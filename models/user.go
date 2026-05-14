package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name"     gorm:"not null"`
	Email    string `json:"email"    gorm:"uniqueIndex;not null"`
	Password string `json:"-"`
}

type RegisterUserInput struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserInput struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
