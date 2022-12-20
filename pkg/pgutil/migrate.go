package pgutil

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/ent"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// MakeMigrate
func MakeMigrate(ctx context.Context, config *PostgresConfig) error {
	db, err := NewDatabase(ctx, config)
	if err != nil {
		return err
	}

	source, err := iofs.New(ent.Migrations, "migrations")
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("iofs", source, config.DatabaseName, driver)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
