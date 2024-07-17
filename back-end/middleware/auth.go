package middleware

import (
	"back-end/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("jwt")

		_, err = userService.GetCurrentUser(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
