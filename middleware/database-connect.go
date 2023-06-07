package middleware

import (
	"fmt"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DatabaseConnect(db *gorm.DB) gin.HandlerFunc {
	// Auto migrate and update schemas
	db.AutoMigrate(&entity.User{}, &entity.SourceCodeHostIntegration{}, &entity.CloudProviderIntegration{})
	fmt.Println("Database: Schema Update Successful")

	db.Create(&entity.User{Email: "test@email.com", Role: "admin", Password: "test123"})

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
