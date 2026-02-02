package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	green = "\033[32m"
	red = "\033[31m"
	reset = "\033[0m"
)

func SimpleLogger() gin.HandlerFunc {
	return  func (c*gin.Context)  {

		start := time.Now()

		c.Next()

		duration:= time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		color := green
		if status >= 400 {
			color = red
		}

		log.Printf("%s%s %s -> %d (%v)%s",color,method,path,status,duration,reset)
		
	}
}