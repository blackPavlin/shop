package config

import (
	"fmt"
	"time"
)

type DatabaseConfig struct {
	Host            string        `envconfig:"HOST" required:"true"`
	Port            int           `envconfig:"PORT" required:"true"`
	User            string        `envconfig:"USER" required:"true"`
	Password        string        `envconfig:"PASSWORD" required:"true"`
	Name            string        `envconfig:"NAME" required:"true"`
	Schema          string        `envconfig:"SCHEMA" required:"true"`
	SSLMode         string        `envconfig:"SSL_MODE" required:"true"`
	MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" required:"true"`
	MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" required:"true"`
	ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" required:"true"`
	ConnMaxIdleTime time.Duration `envconfig:"CONN_MAX_IDLE_TIME" required:"true"`
}

func (c *DatabaseConfig) ToDataSource() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name,
		c.Schema,
		c.SSLMode,
	)
}
