package middleware

import (
	"fmt"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BasicAuth(db *gorm.DB) gin.HandlerFunc {
	var users []entity.User
	err := db.Find(&users).Error
	if err != nil {
		fmt.Println("MIDDLEWARE ERROR: Could not fetch users")
	}

	var accounts = make(gin.Accounts)
	for _, user := range users {
		accounts[user.Email] = user.Password
	}
	return gin.BasicAuth(accounts)
}
