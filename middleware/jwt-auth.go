package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anfern777/cicd-dashboard/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func JwtValidate() gin.HandlerFunc {
	userService := service.NewUserService()

	return func(ctx *gin.Context) {
		bearerToken := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		fmt.Println(bearerToken)
		err := userService.JwtValidateToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.Next()
		}
	}
}
