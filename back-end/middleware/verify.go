package middleware

import (
	"back-end/services"
	"github.com/gin-gonic/gin"
)

func VerifyMiddleware(userService services.UserService, role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//implement verify email
		ctx.Next()
	}
}
