package entity

import "gorm.io/gorm"

type Role string
const(
	Guest Role = "Guest"
	Admin Role = "Admin"
)

type User struct {
	gorm.Model
	Privilege Role `json:"privilege"`
	Email string `json:"email"`
	Password string `json:"-"`
}

