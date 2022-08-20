package server

import (
	"github.com/blackPavlin/shop/app/internal/controllers/rest/auth"
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

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")

		auth.NewAuthController(s.usecases.AuthUseCase).RegisterRoutes(v1)
	}

	return router
}
