package configuration

type Config struct {
	Server struct {
		HTTPListenAddr string `yaml:"HTTP_LISTEN_ADDR"`
	} `yaml:"Server"`
	Database struct {
		PostgresHost     string `yaml:"POSTGRES_HOST"`
		PostgresPort     string `yaml:"POSTGRES_PORT"`
		PostgresUser     string `yaml:"POSTGRES_USER"`
		PostgresPassword string `yaml:"POSTGRES_PASSWORD"`
		PostgresName     string `yaml:"POSTGRES_DB"`
	} `yaml:"Database"`

	Cache struct {
		RedisHost     string `yaml:"REDIS_HOST"`
		RedisPort     int    `yaml:"REDIS_PORT"`
		RedisPassword string `yaml:"REDIS_PASSWORD"`
	} `yaml:"Cache"`

	RateLimit struct {
		MaxRequests  int `yaml:"MAX_REQUESTS"`
		TimeInterval int `yaml:"TIME_INTERVAL"`
	} `yaml:"Rate_Limit"`

	JwtToken struct {
		SecretKey string `yaml:"SECRET_KEY"`
		Period    int    `yaml:"PERIOD"` // in days
	} `yaml:"JWT_TOKEN"`
	Stripe struct {
		Key string `yaml:"KEY"`
	}
}

func NewConfig() (*Config, error) {
	return LoadConfig()
}
