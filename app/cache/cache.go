package cache

import "time"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)
	SetWithExpiration(key string, value interface{}, expiration time.Duration) error
	DecrBy(key string, decrement int64) (int64, error)
}
