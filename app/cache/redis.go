package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "admin",
		}),
	}
}

// Get retrieves a value from the cache as a string
func (c *RedisCache) Get(key string) (string, error) {
	val, err := c.client.Get(key).Result()
	return val, err
}


func (c *RedisCache) Set(key string, value interface{}) error {
	return c.client.Set(key, value, 0).Err()
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}

func (c *RedisCache) Delete(key string) error {
	return c.client.Del(key).Err()
}

func (c *RedisCache) Ping() (string, error) {
	return c.client.Ping().Result()
}

func (c *RedisCache) SetWithExpiration(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(key, value, expiration).Err()
}

func (c *RedisCache) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.client.SetNX(key, value, expiration).Result()
}

func (c *RedisCache) DecrBy(key string, decrement int64) (int64, error) {
	return c.client.DecrBy(key, decrement).Result()
}