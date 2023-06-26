package middleware

import (
	"net/http"
	"strings"

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
			ctx.Set("jwt-claims", claims)
			ctx.Next()
		}
	}
}
