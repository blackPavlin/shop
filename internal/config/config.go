// Package config contains application configs.
package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/blackPavlin/shop/pkg/zapx"
)

// Config is application configuration.
type Config struct {
	Logger   *zapx.Config
	Server   *HTTPServerConfig   `envconfig:"HTTP_SERVER" required:"true"`
	Cors     *CorsConfig         `envconfig:"CORS"        required:"true"`
	Postgres *DatabaseConfig     `envconfig:"POSTGRES"    required:"true"`
	Auth     *AuthConfig         `envconfig:"AUTH"        required:"true"`
	S3       *S3Config           `envconfig:"S3"          required:"true"`
	Search   *SearchEngineConfig `envconfig:"SEARCH"      required:"true"`
}

// NewConfig create application config.
func NewConfig() (*Config, error) {
	config := new(Config)

	if err := envconfig.Process("", config); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return config, nil
}
