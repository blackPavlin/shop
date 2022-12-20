package pgutil

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// PostgresConfig
type PostgresConfig struct {
	Host                 string        `envconfig:"HOST" required:"true"`
	Port                 int           `envconfig:"PORT" required:"true"`
	User                 string        `envconfig:"USER" required:"true"`
	Password             string        `envconfig:"PASSWORD" required:"true"`
	DatabaseName         string        `envconfig:"DATABASE_NAME" required:"true"`
	SSLMode              string        `envconfig:"SSL_MODE" required:"true"`
	MaxOpenConns         int           `envconfig:"MAX_OPEN_CONNS" required:"true"`
	ConnMaxLifetime      time.Duration `envconfig:"CONN_MAX_LIFETIME" required:"true"`
	MaxIdleConns         int           `envconfig:"MAX_IDLE_CONNS" required:"true"`
	ConnMaxIdleTime      time.Duration `envconfig:"CONN_MAX_IDLE_TIME"`
	PreferSimpleProtocol string        `envconfig:"PREFER_SIMPLE_PROTOCOL" required:"true"`
	StatementCacheMode   string        `envconfig:"STATEMENT_CACHE_MODE" required:"true"`
	ApplicationName      string        `envconfig:"APPLICATION_NAME" required:"true"`
}

// ToDataSourceName
func (p *PostgresConfig) ToDataSourceName() string {
	return fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s 
			sslmode=%s prefer_simple_protocol=%s statement_cache_mode=%s application_name=%s`,
		p.Host, p.Port, p.User, p.Password, p.DatabaseName, p.SSLMode, p.PreferSimpleProtocol,
		p.StatementCacheMode, p.ApplicationName,
	)
}

// NewDatabase
func NewDatabase(ctx context.Context, config *PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", config.ToDataSourceName())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
