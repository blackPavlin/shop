package redisx

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// RedisConfig
type RedisConfig struct {
	Addr     string `envconfig:"ADDRESS" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DB       int    `envconfig:"DB" required:"true"`
}

// NewClient
func NewClient(ctx context.Context, config *RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return client, nil
}
