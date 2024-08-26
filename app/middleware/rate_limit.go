package middleware

import (
	"go-backed/app/cache"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// RateLimitMiddleware limits the number of requests from a specific IP address
func RateLimitMiddleware(cache cache.Cache, limit int, period time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		redisKey := "rate_limit:" + ip

		// Attempt to set the initial limit in Redis
		set, err := cache.SetNX(redisKey, limit, period)
		if err != nil {
			log.Printf("Error setting key %s: %v", redisKey, err)
			c.Next()
			return
		}

		// If the key was just set, the request is allowed
		if set {
			c.Next()
			return
		}

		// Fetch the current request count from Redis
		currentCount, err := cache.Get(redisKey)
		if err != nil {
			if err == redis.Nil {
				log.Printf("Key %s not found", redisKey)
			}
			c.Next()
			return
		}

		// Convert the current count to an integer
		currentCountInt, err := strconv.Atoi(currentCount)
		if err != nil {
			log.Printf("Error converting count to int: %v", err)

			c.Next()
			return
		}

		// If the current count is greater than 0, decrement it and allow the request
		if currentCountInt > 0 {
			_, err = cache.DecrBy(redisKey, 1)
			if err != nil {
				// Handle error appropriately (e.g., log it)
				c.Next()
				return
			}
			c.Next()
		} else {
			// Rate limit exceeded
			c.Header("Retry-After", strconv.Itoa(int(period.Seconds())))
			c.String(429, "Rate limit exceeded")
			c.Abort()
		}
	}
}
