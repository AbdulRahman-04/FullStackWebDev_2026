package middleware

import "github.com/gin-gonic/gin"

func OnlyUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "user" {
			c.JSON(400, gin.H{
				"msg": "only users allowed",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func OnlyAdmins() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(401, gin.H{
				"msg": "only admins allwoed",
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
