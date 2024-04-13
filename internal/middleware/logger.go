package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		reqMethod := ctx.Request.Method
		reqURI := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		fmt.Printf("| %3d | %13v | %15s | %s | %s |\n",
			statusCode,
			latency,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
