package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	authHttp "github.com/blackPavlin/shop/app/internal/auth/delivery"
	authUC "github.com/blackPavlin/shop/app/internal/auth/usecase"
	"github.com/blackPavlin/shop/app/internal/config"
	userHttp "github.com/blackPavlin/shop/app/internal/user/delivery"
	userRepo "github.com/blackPavlin/shop/app/internal/user/repository"
	userUC "github.com/blackPavlin/shop/app/internal/user/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Start(config *config.Config, mongo *mongo.Database) error {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	}))

	// User
	userRepo := userRepo.NewUserRepository(mongo)
	userUseCase := userUC.NewUserUseCase(userRepo)
	userHandler := userHttp.NewHandler(userUseCase)

	// Auth
	authUseCase := authUC.NewAuthUseCase(userUseCase, config)
	authHandler := authHttp.NewHandler(authUseCase)

	{
		api := router.Group("/api")

		authHandler.Register(api)

		{
			app := api.Group("/app")

			userHandler.Register(app)
		}
	}

	server := &http.Server{
		Addr:    config.Http.Port,
		Handler: router,
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
