package main

import (
	"context"
	"log"

	"github.com/blackPavlin/shop/app/internal/config"
	"github.com/blackPavlin/shop/app/internal/server"
	"github.com/blackPavlin/shop/app/pkg/logger"
	"github.com/blackPavlin/shop/app/pkg/mongo"
)

func main() {
	// Initialize config
	config, err := config.NewConfig("configs/local.config.yml")
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}

	logger, err := logger.NewLogger(config.Logger.Production)
	if err != nil {
		log.Fatalf("failed to create logger: %+v", err)
	}

	// Connect to mongoDB
	mongo, err := mongo.NewMongoDB(context.Background(), config.Mongo.URL, config.Mongo.Name)
	if err != nil {
		log.Fatalf("failed to connect mongodb: %+v", err)
	}

	// Start server
	if err := server.NewServer(config, logger, mongo).Start(); err != nil {
		log.Fatalf("failed to shutdown server: %+v", err)
	}
}
