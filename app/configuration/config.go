package configuration

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() (*Config, error) {
	var config Config
	var yamlFile []byte
	var err error

	isProd := os.Getenv("GO_ENV") == "production"
	if isProd {
		yamlFile, err = os.ReadFile("./config/prod.yaml")
	} else {
		yamlFile, err = os.ReadFile("./config/dev.yaml")
	}

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config) // Pass the pointer to the config struct
	if err != nil {
		return nil, err
	}

	// Check and set required environment variables
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return nil, errors.New("POSTGRES_PASSWORD environment variable is required")
	}
	config.Database.PostgresPassword = postgresPassword

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		return nil, errors.New("REDIS_PASSWORD environment variable is required")
	}
	config.Cache.RedisPassword = redisPassword

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		return nil, errors.New("JWT_SECRET_KEY environment variable is required")
	}
	config.JwtToken.SecretKey = jwtSecretKey
	stripeKey := os.Getenv("STRIPE_KEY")
	if stripeKey == "" {
		return nil, errors.New("STRIPE_KEY environment variable is required")
	}
	config.Stripe.Key = stripeKey

	return &config, nil
}
