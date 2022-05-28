package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Http struct {
		Port string `yaml:"port" env-required:"true"`
		Mode string `yaml:"mode" env-required:"true"`
	}
	Mongo struct {
		URL  string `yaml:"url" env-required:"true"`
		Name string `yaml:"name" env-required:"true"`
	}
	Auth struct {
		SigningKey string `yaml:"signing_key" env-required:"true"`
		ExpiresIn  int    `yaml:"expires_in" env-required:"true"`
	}
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}

	if err := cleanenv.ReadConfig(path, config); err != nil {
		return nil, err
	}

	return config, nil
}
