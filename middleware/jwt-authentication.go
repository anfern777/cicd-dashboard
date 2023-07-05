package middleware

import (
	"net/http"
	"strings"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

func JwtValidate() gin.HandlerFunc {
	userService := service.NewUserService()

	return func(ctx *gin.Context) {
		bearerToken := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		claims, err := userService.JwtValidateToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			var user entity.User
			email := claims.Email
			err := service.GetDB(ctx).Where("email = ?", email).First(&user).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error: ": "could not find user from jwt"})
			}
			ctx.Set("jwt-user", user)
			ctx.Next()
		}
	}
}
