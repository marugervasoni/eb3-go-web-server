package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//creamos función logger de tipo handler funcion
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verb := ctx.Request.Method
		time := time.Now()
		url := ctx.Request.URL
		var size int

		ctx.Next()

		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		fmt.Printf("\nPath:%s\nVerb:%s\nTime:%v\nRequest size:%d\n", url, verb, time, size)
	}
}