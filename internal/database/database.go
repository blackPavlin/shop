package database

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/database/ent"
)

func NewDatabase(ctx context.Context, config *config.DatabaseConfig) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(config.ToDataSource())
	if err != nil {
		return nil, fmt.Errorf("parse database pool config: %w", err)
	}

	conf.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	conf.MaxConnIdleTime = config.ConnMaxIdleTime
	conf.MaxConnLifetime = config.ConnMaxLifetime

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping database connection: %w", err)
	}

	return pool, nil
}

func NewClient(pool *pgxpool.Pool) *ent.Client {
	driver := sql.OpenDB(dialect.Postgres, stdlib.OpenDBFromPool(pool))
	client := ent.NewClient(ent.Driver(driver))

	return client
}
