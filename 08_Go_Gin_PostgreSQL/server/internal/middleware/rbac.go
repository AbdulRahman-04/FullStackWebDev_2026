package middleware

import "github.com/gin-gonic/gin"

func OnlyAdmins() gin.HandlerFunc{
	return func (c*gin.Context)  {
		role, exists := c.Get("role")
		if !exists || role != "admin"{
			c.JSON(400, gin.H{
				"msg": "Only Admins Can access",
			})
			c.Abort()
			return 
		}
		c.Next()
		
	}
}

func OnlyUsers() gin.HandlerFunc{
	return  func (c*gin.Context)  {
		role, exists := c.Get("role")
		if !exists || role != "user"{
			c.JSON(400, gin.H{
				"msg": "only users allwoed",
			})
			c.Abort()
			return 
		}
		c.Next()
		
	}
}