package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/server"
	"github.com/blackPavlin/shop/pkg/logger"
	"github.com/blackPavlin/shop/pkg/pgutil"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}

	logger, err := logger.NewLogger(config.Logger)
	if err != nil {
		log.Fatalf("failed to create logger: %+v", err)
	}

	if err := pgutil.MakeMigrate(context.Background(), config.Postgres); err != nil {
		log.Fatalf("failed to make migrations: %+v", err)
	}

	db, err := pgutil.NewDatabase(context.Background(), config.Postgres)
	if err != nil {
		log.Fatalf("failed to connect database: %+v", err)
	}

	driver := sql.OpenDB(dialect.Postgres, db.DB)
	database := ent.NewClient(ent.Driver(driver))

	if err := server.NewServer(config, logger, database).Run(); err != nil {
		log.Fatalf("failed to shutdown server: %+v", err)
	}
}
