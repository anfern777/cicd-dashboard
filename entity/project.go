package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UserID      uint   `json:"userID"`
	Name        string `binding:"required"`
	Description string `json:"description"`
}
