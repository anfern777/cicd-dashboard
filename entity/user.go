package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Projects []Project
	Role     string `binding:"required" validate:"is-role"`
	Email    string `binding:"required,email" gorm:"unique"`
	Password string `binding:"required"`
}
