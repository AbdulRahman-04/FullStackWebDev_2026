package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {

			if err := recover(); err != nil {
				log.Printf("err is %v", err)

				c.AbortWithStatusJSON(500, gin.H{
					"msg": "internal server error",
				})
			}

		}()

		c.Next()

	}
}
