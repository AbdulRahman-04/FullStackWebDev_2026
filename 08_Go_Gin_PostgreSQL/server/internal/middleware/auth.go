package middleware

import (
	"strings"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var jwtKey = []byte(config.AppConfig.JWT_KEY)

		// get authHeader
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(400, gin.H{
				"msg": "missing token",
			})
			c.Abort()
			return
		}

		// get auth headers in parts
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"msg": "invaid token format",
			})
			c.Abort()
			return
		}

		tokenStr := parts[1]

		// jwt verify
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtKey, nil

		})

		if err != nil || !token.Valid {
			c.JSON(400, gin.H{
				"msg": "invalid token or expired",
			})
			c.Abort()
			return
		}

		// get claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{
				"msg": "no claims found",
			})
			c.Abort()
			return
		}

		userId, ok := claims["id"].(string)
		if !ok {
			c.JSON(401, gin.H{
				"msg": "no user id found",
			})
			c.Abort()
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(401, gin.H{
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
