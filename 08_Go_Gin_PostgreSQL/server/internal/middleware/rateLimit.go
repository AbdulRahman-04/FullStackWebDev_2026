package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

func RateLimitMiddleware(limit int) gin.HandlerFunc{
	rate := limiter.Rate{
		Period: 1*time.Minute,
		Limit: int64(limit),
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)

	return  func (c*gin.Context)  {
		key := c.FullPath()+"-"+c.ClientIP()

		limitCtx, err := instance.Get(c,key)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"msg": "rate limit err",
			})
			return 
		}

		if limitCtx.Reached {
			c.AbortWithStatusJSON(429, gin.H{
				"msg": "rate limit exceeded",
			})
			return 
		}

		c.Next()
		
	}
}