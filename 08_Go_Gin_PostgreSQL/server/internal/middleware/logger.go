package middleware

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	
)

var (
	green = "\033[32m"
	red = "\033[31m"
	reset =  "\033[0m"
)

func InitSimpleLogger() *log.Logger {
  
	if _, err := os.Stat("logs"); os.IsNotExist(err){
		os.Mkdir("logs", 0755)
	}

	file, err := os.OpenFile("logs/server/log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
		panic(err)
	}

	multi := io.MultiWriter(os.Stdout, file)
	return  log.New(multi, "", log.LstdFlags)
}

func SimpleLogger(logger *log.Logger) gin.HandlerFunc {
	return func (c*gin.Context)  {

		start := time.Now

		c.Next()


		duration := time.Since(start())
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		color := green 
		if status >= 400 {
			color = red
		}

		logger.Printf("%s%s %s -> %d (%v)%s",
	color,method,path,status,duration,reset)
		
	}
}