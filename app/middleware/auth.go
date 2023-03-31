package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckUserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
