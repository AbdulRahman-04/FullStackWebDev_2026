package middleware

import (
	"strings"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.AppConfig.JWT_KEY)


func AuthMiddleware() gin.HandlerFunc {
	return  func (c*gin.Context)  {

		// get auth header 
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"msg": "no token provided",
			})
			c.Abort()
			return 
		}

		// get auth header spilit into parts 
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"msg": "invalid token format",
			})
			c.Abort()
			return 
		}

		//get token from parts 
		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func (r*jwt.Token)(interface{}, error)  {
			return  jwtKey, nil
		})

		if err != nil || !token.Valid {
			 c.JSON(401, gin.H{
				"msg": "invalid or expired token",
			 })
			 c.Abort()
			 return 
		}

		// get claims from token 
		claims := token.Claims.(jwt.MapClaims)

		userID := claims["id"].(string)
		role := claims["role"].(string)

		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
		
	}
}