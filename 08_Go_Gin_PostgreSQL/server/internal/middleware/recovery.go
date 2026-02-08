package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return  func (c*gin.Context)  {
		defer func ()  {
			if err := recover(); err != nil {
				log.Printf("err %v", err)

				c.AbortWithStatusJSON(401, gin.H{
					"msg": "internal server err",
				})
			}
		}()
		c.Next()
	}
}