// Package config contains application configs.
package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/blackPavlin/shop/pkg/s3x"
	"github.com/blackPavlin/shop/pkg/searchx"
	"github.com/blackPavlin/shop/pkg/zapx"
)

// Config is application configuration.
type Config struct {
	Logger   *zapx.Config
	Server   *ServerConfig         `envconfig:"REST_SRV" required:"true"`
	Cors     *CorsConfig           `envconfig:"CORS" required:"true"`
	Postgres *DatabaseConfig       `envconfig:"POSTGRES" required:"true"`
	Auth     *AuthConfig           `envconfig:"AUTH" required:"true"`
	S3       *s3x.S3Config         `envconfig:"S3" required:"true"`
	Search   *searchx.SearchConfig `envconfig:"SEARCH" required:"true"`
}

// ServerConfig is HTTP server configuration.
type ServerConfig struct {
	Address      string        `envconfig:"ADDRESS" required:"true"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" required:"true"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" required:"true"`
	IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT" required:"true"`
}

// CorsConfig defines parameters for CORS related request handler.
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

// AuthConfig is authentication configuration.
type AuthConfig struct {
	ExpiresIn  int    `envconfig:"EXPIRES_IN" required:"true"`
	SigningKey string `envconfig:"SIGNING_KEY" required:"true"`
}

// NewConfig create application config.
func NewConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}
	return config, nil
}
