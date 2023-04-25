package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserRole Role   `json:"user_role" binding:"required"`
}

type Role int

const (
	Client Role = 1
	Author Role = 2
)
