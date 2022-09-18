package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Http   HttpConfig   `yaml:"http" env-required:"true"`
	Mongo  MongoConfig  `yaml:"mongo" env-required:"true"`
	Auth   AuthConfig   `yaml:"auth" env-required:"true"`
	Logger LoggerConfig `yaml:"logger" env-required:"true"`
}

type HttpConfig struct {
	Port string `yaml:"port" env-required:"true"`
	Mode string `yaml:"mode" env-required:"true"`
}

type LoggerConfig struct {
	Production bool `yaml:"production"`
}

type MongoConfig struct {
	URL  string `yaml:"url" env-required:"true"`
	Name string `yaml:"name" env-required:"true"`
}

type AuthConfig struct {
	SigningKey string `yaml:"signing_key" env-required:"true"`
	ExpiresIn  int    `yaml:"expires_in" env-required:"true"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}

	if err := cleanenv.ReadConfig(path, config); err != nil {
		return nil, err
	}

	return config, nil
}
