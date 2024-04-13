package handlers

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ProcessImage(ctx *gin.Context) {
	imageFile, err := ctx.FormFile("imageFile")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	src, err := imageFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tempFile, err := os.CreateTemp("", "processed_*.jpg")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	err = jpeg.Encode(tempFile, img, &jpeg.Options{Quality: 5})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.File(tempFile.Name())
}
