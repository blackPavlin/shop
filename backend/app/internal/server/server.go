package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackPavlin/shop/app/internal/adapters/mongodb"
	"github.com/blackPavlin/shop/app/internal/config"
	v1 "github.com/blackPavlin/shop/app/internal/controllers/http/v1"
	"github.com/blackPavlin/shop/app/internal/controllers/http/v1/middlewares"
	"github.com/blackPavlin/shop/app/internal/domain/services"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/auth"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/cart"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/category"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/order"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/product"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/user"
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

	// Repository
	userRepository := mongodb.NewUserRepository(mongo)
	cartRepository := mongodb.NewCartRepository(mongo)
	orderPerository := mongodb.NewOrderRepository(mongo)
	productRepository := mongodb.NewProductRepository(mongo)
	categoryRepository := mongodb.NewCategoryRepository(mongo)

	// Services
	userService := services.NewUserService(userRepository)
	cartService := services.NewCartService(cartRepository)
	orderService := services.NewOrderService(orderPerository)
	productService := services.NewProductService(productRepository)
	categoryService := services.NewCategoryService(categoryRepository)

	// UseCases
	userUseCase := user.NewUserUseCase(userService)
	cartUseCase := cart.NewCartUseCase(cartService, productService)
	authUseCase := auth.NewAuthUseCase(userService, cartService, config.Auth)
	orderUseCase := order.NewOrderUseCase(orderService, cartService)
	productUseCase := product.NewProductUseCase(productService, categoryService)
	categoryUseCase := category.NewCategoryUseCase(categoryService)

	// Middlewares
	authMiddleware := middlewares.NewAuthMiddleware(authUseCase)

	// Controllers
	authController := v1.NewAuthHandler(authUseCase)
	userController := v1.NewUserHandler(userUseCase, authMiddleware)
	cartController := v1.NewCartHandler(cartUseCase, authMiddleware)
	orderController := v1.NewOrderHandler(orderUseCase, authMiddleware)
	productController := v1.NewProductHandler(productUseCase, authMiddleware)
	categoryController := v1.NewCategoryHandler(categoryUseCase, authMiddleware)

	api := router.Group("/api")

	{
		v1 := api.Group("/v1")

		authController.Register(v1)
		userController.Register(v1)
		cartController.Register(v1)
		orderController.Register(v1)
		productController.Register(v1)
		categoryController.Register(v1)
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
