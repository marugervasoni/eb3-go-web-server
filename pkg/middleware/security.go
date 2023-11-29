package middleware

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)


func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//traemos logica de validación de token 		
		tokenHeader := ctx.GetHeader("tokenPostman")
		tokenEnv := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		} else {
			ctx.Next()
		}

	}
}