// Package server contains implementations for http server.
package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/domain/address"
	addresspg "github.com/blackPavlin/shop/internal/domain/address/pg"
	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/domain/cart"
	cartpg "github.com/blackPavlin/shop/internal/domain/cart/pg"
	"github.com/blackPavlin/shop/internal/domain/category"
	categorypg "github.com/blackPavlin/shop/internal/domain/category/pg"
	imagestorage "github.com/blackPavlin/shop/internal/domain/image/storage"
	"github.com/blackPavlin/shop/internal/domain/product"
	productpg "github.com/blackPavlin/shop/internal/domain/product/pg"
	"github.com/blackPavlin/shop/internal/domain/user"
	userpg "github.com/blackPavlin/shop/internal/domain/user/pg"
	"github.com/blackPavlin/shop/internal/transport/rest/controller"
	restmiddleware "github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
)

// Server represents http server.
type Server struct {
	config *config.Config
	logger *zap.Logger
	router *chi.Mux

	database *ent.Client
	storage  *minio.Client
}

// NewServer create instance of http server.
func NewServer(
	conf *config.Config,
	logger *zap.Logger,
	database *ent.Client,
	storage *minio.Client,
) *Server {
	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Logger,
	)

	router.Use(cors.New(cors.Options{
		AllowedOrigins:     conf.Cors.AllowedOrigins,
		AllowedMethods:     conf.Cors.AllowedMethods,
		AllowedHeaders:     conf.Cors.AllowedHeaders,
		ExposedHeaders:     conf.Cors.ExposedHeaders,
		AllowCredentials:   conf.Cors.AllowCredentials,
		MaxAge:             conf.Cors.MaxAge,
		OptionsPassthrough: conf.Cors.OptionsPassthrough,
		Debug:              conf.Cors.Debug,
	}).Handler)

	// Storages
	imageStorage := imagestorage.NewStorage(storage, conf.S3, logger)

	// Repositories
	txManager := pg.NewTxManager(database, logger)
	userRepository := userpg.NewUserRepository(database, logger)
	cartRepository := cartpg.NewCartRepository(database, logger)
	addressRepository := addresspg.NewAddressRepository(database, logger)
	productRepository := productpg.NewProductRepository(database, logger)
	imageRepository := productpg.NewImageRepository(database, logger)
	categoryRepository := categorypg.NewCategoryRepository(database, logger)

	// Services
	userService := user.NewUseCase(userRepository)
	authService := auth.NewUseCase(logger, conf.Auth, userRepository)
	cartService := cart.NewUseCase(cartRepository, productRepository)
	addressService := address.NewUseCase(addressRepository)
	productService := product.NewUseCase(productRepository, imageRepository, imageStorage, txManager)
	imageService := product.NewImageUseCase(logger, productRepository, imageRepository, imageStorage, txManager)
	categoryService := category.NewUseCase(categoryRepository)

	// Middlewares
	authMiddleware := restmiddleware.NewMiddleware(authService)

	// Controllers
	userController := controller.NewUserController(userService, authMiddleware)
	authController := controller.NewAuthController(authService)
	cartController := controller.NewCartController(cartService, authMiddleware)
	addressController := controller.NewAddressController(addressService, authMiddleware)
	productController := controller.NewProductController(productService, imageService, authMiddleware)
	categoryController := controller.NewCategoryController(categoryService, authMiddleware)

	router.Route("/api", func(r chi.Router) {
		userController.RegisterRoutes(r)     // /api/user
		authController.RegisterRoutes(r)     // /api/auth
		cartController.RegisterRoutes(r)     // /api/cart
		addressController.RegisterRoutes(r)  // /api/address
		productController.RegisterRoutes(r)  // /api/product
		categoryController.RegisterRoutes(r) // /api/category
	})

	return &Server{conf, logger, router, database, storage}
}

// Run http server.
func (s *Server) Run() error {
	server := http.Server{
		Addr:         s.config.Server.Address,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
		IdleTimeout:  s.config.Server.IdleTimeout,
		Handler:      s.router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %+v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server error: %w", err)
	}
	if err := s.database.Close(); err != nil {
		return fmt.Errorf("close database connection error: %w", err)
	}
	return nil
}
