package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	"github.com/blackPavlin/shop/internal/domain/product"
	productpg "github.com/blackPavlin/shop/internal/domain/product/pg"
	"github.com/blackPavlin/shop/internal/domain/user"
	userpg "github.com/blackPavlin/shop/internal/domain/user/pg"
	"github.com/blackPavlin/shop/internal/transport/rest/controller"
)

// Server
type Server struct {
	config *config.Config
	logger *zap.Logger
	router *chi.Mux

	database *ent.Client
}

// NewServer
func NewServer(config *config.Config, logger *zap.Logger, database *ent.Client) *Server {
	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Logger,
	)

	router.Use(cors.New(cors.Options{
		AllowedOrigins:     config.Cors.AllowedOrigins,
		AllowedMethods:     config.Cors.AllowedMethods,
		AllowedHeaders:     config.Cors.AllowedHeaders,
		ExposedHeaders:     config.Cors.ExposedHeaders,
		AllowCredentials:   config.Cors.AllowCredentials,
		MaxAge:             config.Cors.MaxAge,
		OptionsPassthrough: config.Cors.OptionsPassthrough,
		Debug:              config.Cors.Debug,
	}).Handler)

	// Repositories
	userRepository := userpg.NewUserRepository(database, logger)
	cartRepository := cartpg.NewCartRepository(database, logger)
	addressRepository := addresspg.NewAddressRepository(database, logger)
	productRepository := productpg.NewProductRepository(database, logger)
	categoryRepository := categorypg.NewCategoryRepository(database, logger)

	// Services
	userService := user.NewUserUseCase(userRepository)
	authService := auth.NewUseCase(logger, config.Auth, userRepository)
	cartService := cart.NewUseCase(cartRepository, productRepository)
	addressService := address.NewUseCase(addressRepository)
	productService := product.NewUseCase(productRepository)
	categoryService := category.NewUseCase(categoryRepository)

	// Controllers
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)
	cartController := controller.NewCartController(cartService)
	addressController := controller.NewAddressController(addressService)
	productController := controller.NewProductController(productService)
	categoryController := controller.NewCategoryController(categoryService)

	// Middlewares

	router.Route("/api", func(r chi.Router) {
		userController.RegisterRoutes(r)     // /api/user
		authController.RegisterRoutes(r)     // /api/auth
		cartController.RegisterRoutes(r)     // /api/cart
		addressController.RegisterRoutes(r)  // /api/address
		productController.RegisterRoutes(r)  // /api/product
		categoryController.RegisterRoutes(r) // /api/category
	})

	return &Server{
		config:   config,
		logger:   logger,
		router:   router,
		database: database,
	}
}

// Run
func (s *Server) Run() error {
	server := http.Server{
		Addr:         s.config.Server.Address,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
		IdleTimeout:  s.config.Server.IdleTimeout,
		Handler:      s.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %+v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
