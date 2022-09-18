package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackPavlin/shop/app/internal/config"
	"github.com/blackPavlin/shop/app/internal/repositories"
	"github.com/blackPavlin/shop/app/internal/services"
	"github.com/blackPavlin/shop/app/internal/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Server
type Server struct {
	config *config.Config
	mongo  *mongo.Database
	logger *zap.Logger
	router *gin.Engine

	repositories repositories.Repositories
	services     services.Services
	usecases     usecases.UseCases
}

// NewSever
func NewServer(
	config *config.Config,
	logger *zap.Logger,
	mongo *mongo.Database,
) Server {
	s := Server{
		config: config,
		logger: logger,
		mongo:  mongo,
	}

	s.repositories = s.initRepositories()
	s.services = s.initServices()
	s.usecases = s.initUseCases()
	s.router = s.initRoutes()

	return s
}

// Start
func (s Server) Start() error {
	server := &http.Server{
		Handler:      s.router,
		Addr:         s.config.Http.Port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
