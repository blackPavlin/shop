// Package config contains application configs.
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/blackPavlin/shop/pkg/s3x"
	"github.com/blackPavlin/shop/pkg/searchx"
	"github.com/blackPavlin/shop/pkg/zapx"
)

// Config is application configuration.
type Config struct {
	Logger   *zapx.Config
	Server   *HttpServerConfig     `envconfig:"HTTP_SERVER" required:"true"`
	Cors     *CorsConfig           `envconfig:"CORS" required:"true"`
	Postgres *DatabaseConfig       `envconfig:"POSTGRES" required:"true"`
	Auth     *AuthConfig           `envconfig:"AUTH" required:"true"`
	S3       *s3x.S3Config         `envconfig:"S3" required:"true"`
	Search   *searchx.SearchConfig `envconfig:"SEARCH" required:"true"`
}

// AuthConfig is authentication configuration.
type AuthConfig struct {
	ExpiresIn  int    `envconfig:"EXPIRES_IN" required:"true"`
	SigningKey string `envconfig:"SIGNING_KEY" required:"true"`
}

// NewConfig create application config.
func NewConfig() (*Config, error) {
	config := new(Config)

	if err := envconfig.Process("", config); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return config, nil
}
