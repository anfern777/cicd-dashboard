package middleware

import (
	"net/http"

	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

func RequireRole(roleRequired string) gin.HandlerFunc {
	userService := service.NewUserService()

	return func(ctx *gin.Context) {
		claims := (ctx.Keys["jwt-claims"].(*service.Claims))
		user, err := userService.FindByEmail(ctx, claims.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if user.Role != roleRequired {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Insufficcient permissions"})
		}

		ctx.Next()
	}
}
