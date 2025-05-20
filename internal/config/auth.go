package config

// AuthConfig is authentication configuration.
type AuthConfig struct {
	ExpiresIn  int    `envconfig:"EXPIRES_IN"  required:"true"`
	SigningKey string `envconfig:"SIGNING_KEY" required:"true"`
}
