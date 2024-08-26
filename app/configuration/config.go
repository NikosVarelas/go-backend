package configuration

import (
	"os"
	"time"
)

type Config struct {
	Database       *PGConfig
	Cache          *RedisConfig
	JWTSecretKey   string
	HTTPListenAddr string
	IsDocker       string
	RateLimit      *RateLimitConfig
}

type RateLimitConfig struct {
	MaxRequests int
	Period      time.Duration
}

func NewRateLimitConfig() *RateLimitConfig {
	return &RateLimitConfig{
		MaxRequests: 100,
		Period:      1 * time.Minute,
	}
}

type PGConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func NewPGConfig() *PGConfig {
	return &PGConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}

type RedisConfig struct {
	Host     string
	Password string
	Port     string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Port:     os.Getenv("REDIS_PORT"),
	}
}

func NewConfig() *Config {
	return &Config{
		Database:       NewPGConfig(),
		Cache:          NewRedisConfig(),
		JWTSecretKey:   os.Getenv("JWT_SECRET"),
		HTTPListenAddr: os.Getenv("HTTP_LISTEN_ADDR"),
		IsDocker:       os.Getenv("IS_DOCKER"),
		RateLimit:      NewRateLimitConfig(),
	}
}
