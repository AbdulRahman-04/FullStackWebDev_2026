package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprint(time.Now().Unix()) + "_" + file.Filename
	filePath := filepath.Join("uploads", fileName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "", err
	}
	return filePath, err
}
