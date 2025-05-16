package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/migration"
)

func MakeMigrations(pool *pgxpool.Pool, config *config.DatabaseConfig) error {
	source, err := iofs.New(migration.MigrationsFS, "postgres")
	if err != nil {
		return fmt.Errorf("create iofs source: %w", err)
	}

	instance, err := postgres.WithInstance(stdlib.OpenDBFromPool(pool), &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create postgres instance: %w", err)
	}

	migration, err := migrate.NewWithInstance("iofs", source, config.Name, instance)
	if err != nil {
		return fmt.Errorf("create migration with instance: %w", err)
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migratorion up: %w", err)
	}

	return nil
}
