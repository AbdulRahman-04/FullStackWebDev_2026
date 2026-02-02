package middleware

import "github.com/gin-gonic/gin"

// BasicSecurityHeaders adds essential HTTP security headers
func BasicSecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Prevent MIME sniffing
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		// Prevent clickjacking
		c.Writer.Header().Set("X-Frame-Options", "DENY")

		// XSS protection (legacy but safe)
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		// Referrer control
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")

		// Basic CSP (tight but safe for backend)
		c.Writer.Header().Set(
			"Content-Security-Policy",
			"default-src 'none'; frame-ancestors 'none';",
		)

		c.Next()
	}
}
