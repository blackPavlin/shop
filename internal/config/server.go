package config

import "time"

// HttpServerConfig is HTTP server configuration.
type HttpServerConfig struct {
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
