package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {

	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  15,
	}

	store := memory.NewStore()

	instance := limiter.New(store, rate)

	return ginmw.NewMiddleware(instance)
}
