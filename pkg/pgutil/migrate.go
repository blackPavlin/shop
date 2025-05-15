package pgutil

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/blackPavlin/shop/migration"
)

// MakeMigrate applies migrations to the specified database.
func MakeMigrate(ctx context.Context, config *PostgresConfig) error {
	db, err := NewDatabase(ctx, config)
	if err != nil {
		return fmt.Errorf("connect postgres error: %w", err)
	}
	source, err := iofs.New(migration.MigrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("init source error: %w", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("database instance error: %w", err)
	}
	migrator, err := migrate.NewWithInstance("iofs", source, config.DatabaseName, driver)
	if err != nil {
		return fmt.Errorf("migrator instance error: %w", err)
	}
	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrator up error: %w", err)
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("close postgres connection error: %w", err)
	}
	return nil
}
