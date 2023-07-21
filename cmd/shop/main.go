package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/server"
	"github.com/blackPavlin/shop/pkg/pgutil"
	"github.com/blackPavlin/shop/pkg/s3x"
	"github.com/blackPavlin/shop/pkg/zapx"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load conf: %+v", err)
	}

	logger, err := zapx.NewLogger(conf.Logger)
	if err != nil {
		log.Fatalf("failed to create zapx: %+v", err)
	}

	if err := pgutil.MakeMigrate(context.Background(), conf.Postgres); err != nil {
		log.Fatalf("failed to make migrations: %+v", err)
	}

	db, err := pgutil.NewDatabase(context.Background(), conf.Postgres)
	if err != nil {
		log.Fatalf("failed to connect database: %+v", err)
	}

	driver := sql.OpenDB(dialect.Postgres, db.DB)
	database := ent.NewClient(ent.Driver(driver))

	storage, err := s3x.NewStorage(context.Background(), conf.S3)
	if err != nil {
		log.Fatalf("failed to connect file storage: %+v", err)
	}

	if err := server.NewServer(conf, logger, database, storage).Run(); err != nil {
		log.Fatalf("failed to shutdown server: %+v", err)
	}
}
