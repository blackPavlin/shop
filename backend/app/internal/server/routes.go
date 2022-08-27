package server

import (
	"github.com/blackPavlin/shop/app/internal/controllers/rest/auth"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/cart"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/middlewares"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s Server) initRoutes() *gin.Engine {
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

	authMiddleware := middlewares.NewAuthMiddleware(s.usecases.AuthUseCase)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")

		auth.NewAuthController(s.usecases.AuthUseCase).RegisterRoutes(v1)
		user.NewUserController(s.usecases.UserUseCase, authMiddleware).RegisterRoutes(v1)
		cart.NewCartController(s.usecases.CartUseCase, authMiddleware).RegisterRoutes(v1)
	}

	return router
}
