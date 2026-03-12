package middleware

import (
	"strings"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/09Backend_Practice/server/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return  func (c*gin.Context)  {

		var jwtKey = []byte(config.AppConfig.JWT_KEY)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(400, gin.H{
				"msg": "missing token",
			})
			c.Abort()
			return 
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(400, gin.H{
				"msg":"invalid token format",
			})
			c.Abort()
			return 
		}

		tokenStr := parts[1]

		// token verify 
		token, err := jwt.Parse(tokenStr, func (t*jwt.Token) (interface{}, error)  {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return  nil , jwt.ErrSignatureInvalid
			}

			return jwtKey, nil
		})

		if err != nil {
			c.JSON(400, gin.H{
				"msg": "invalid or expired  token",
			})
			c.Abort()
			return 
		}

		// get claims from token 
		claims,ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no claims found in token",
			})
			c.Abort()
			return 
		}

		userId, ok := claims["id"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no userId found",
			})
			c.Abort()
			return 
		}

		role,ok := claims["role"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no role found",
			})
			c.Abort()
			return 
		}

		c.Set("userId", userId)
		c.Set("role", role)

		c.Next()
		
	}
}