// Package pgutil contains utilities for working with Postgres.
package pgutil

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// pgx driver.
	_ "github.com/jackc/pgx/v4/stdlib"
)

// PostgresConfig is a Postgres configuration.
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

func (p *PostgresConfig) toDataSourceName() string {
	return fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s 
			sslmode=%s prefer_simple_protocol=%s statement_cache_mode=%s application_name=%s`,
		p.Host, p.Port, p.User, p.Password, p.DatabaseName, p.SSLMode, p.PreferSimpleProtocol,
		p.StatementCacheMode, p.ApplicationName,
	)
}

// NewDatabase creates database connection and checks it.
func NewDatabase(ctx context.Context, config *PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.toDataSourceName())
	if err != nil {
		return nil, fmt.Errorf("connect to database error: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database error: %w", err)
	}
	return db, nil
}
