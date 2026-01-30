package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c*gin.Context) (string, error) {
	// get file from formfile
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	// create uploads folder 
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	// create filename and path
	fileName := fmt.Sprint(time.Now().Unix()) + "_" + file.Filename
	filePath := filepath.Join("uploads", fileName)

	// save changes
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		return  "", err
	}

	return  filePath, nil
}