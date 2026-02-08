package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

var store = memory.NewStore()

func RateLimitMiddleware(limit int, period time.Duration) gin.HandlerFunc {
	rate := limiter.Rate{
		Period: period,
		Limit:  int64(limit),
	}

	instance := limiter.New(store, rate)

	return func(c *gin.Context) {
		key := c.FullPath() + "-" + c.ClientIP()

		context, err := instance.Get(c, key)
		if err != nil || context.Reached {
			c.Abort()
			return
		}

		c.Next()
	}
}
