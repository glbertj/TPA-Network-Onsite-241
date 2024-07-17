package middleware

import (
	"back-end/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RoleMiddleware(userService services.UserService, role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("jwt")

		user, err := userService.GetCurrentUser(token)
		if err != nil || strings.ToUpper(user.Role) != strings.ToUpper(role) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
