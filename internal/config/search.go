package config

type SearchEngineConfig struct {
	Host   string `envconfig:"HOST"    required:"true"`
	APIKey string `envconfig:"API_KEY" required:"true"`
}
