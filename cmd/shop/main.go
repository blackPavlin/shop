package main

import (
	"context"
	"log"

	_ "go.uber.org/automaxprocs"

	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/database"
	"github.com/blackPavlin/shop/internal/s3"
	"github.com/blackPavlin/shop/internal/search"
	"github.com/blackPavlin/shop/internal/server"
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

	db, err := database.NewDatabase(context.Background(), conf.Postgres)
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
	}

	if err := database.MakeMigrations(db, conf.Postgres); err != nil {
		log.Fatalf("failed to make migrations: %+v", err)
	}

	dbClient := database.NewClient(db)

	storage, err := s3.NewClient(context.Background(), conf.S3)
	if err != nil {
		log.Fatalf("failed to connect file storage: %+v", err)
	}

	search, err := search.NewSearchEngine(conf.Search)
	if err != nil {
		log.Fatalf("failed to connect search engine: %+v", err)
	}

	if err := server.NewServer(conf, logger, dbClient, storage, search).Run(); err != nil {
		log.Fatalf("failed to shutdown server: %+v", err)
	}
}
