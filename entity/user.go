package entity

import "gorm.io/gorm"

type Role string

type User struct {
	gorm.Model
	Role     Role   `binding:"required" validate:"is-role"`
	Email    string `binding:"required,email" gorm:"unique"`
	Password string `json:"-"`
}
