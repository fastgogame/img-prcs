package main

import (
	"img-prcs/internal/handlers"
	"img-prcs/internal/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20

	router.Use(middleware.Logger())

	router.LoadHTMLGlob("internal/templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/process-image", handlers.ProcessImage)

	router.Run(":8888")
}
