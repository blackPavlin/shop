package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/blackPavlin/shop/pkg/logger"
	"github.com/blackPavlin/shop/pkg/pgutil"
	"github.com/blackPavlin/shop/pkg/redisx"
	"github.com/blackPavlin/shop/pkg/s3x"
)

// Config
type Config struct {
	Logger   *logger.LoggerConfig
	Server   *ServerConfig          `envconfig:"REST_SRV" required:"true"`
	Cors     *CorsConfig            `envconfig:"CORS" required:"true"`
	Postgres *pgutil.PostgresConfig `envconfig:"POSTGRES" required:"true"`
	Auth     *AuthConfig            `envconfig:"AUTH" required:"true"`
	S3       *s3x.S3Config          `envconfig:"S3" required:"true"`
	Redis    *redisx.RedisConfig    `envconfig:"REDIS" required:"true"`
}

// ServerConfig
type ServerConfig struct {
	Address      string        `envconfig:"ADDRESS" required:"true"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" required:"true"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" required:"true"`
	IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT" required:"true"`
}

// CorsConfig
type CorsConfig struct {
	AllowedOrigins     []string `envconfig:"ALLOWED_ORIGINS" required:"true"`
	AllowedMethods     []string `envconfig:"ALLOWED_METHODS" required:"true"`
	AllowedHeaders     []string `envconfig:"ALLOWED_HEADERS" required:"true"`
	ExposedHeaders     []string `envconfig:"EXPOSED_HEADERS" required:"true"`
	AllowCredentials   bool     `envconfig:"ALLOW_CREDENTIALS" required:"true"`
	MaxAge             int      `envconfig:"MAX_AGE" required:"true"`
	OptionsPassthrough bool     `envconfig:"OPTIONS_PASSTHROUGH" required:"true"`
	Debug              bool     `envconfig:"DEBUG" required:"true"`
}

// AuthConfig
type AuthConfig struct {
	ExpiresIn  int    `envconfig:"EXPIRES_IN" required:"true"`
	SigningKey string `envconfig:"SIGNING_KEY" required:"true"`
}

// NewConfig
func NewConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
