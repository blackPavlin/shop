package category

import (
	"github.com/blackPavlin/shop/app/internal/controllers/rest/middlewares"
	"github.com/gin-gonic/gin"
)

// CategoryUseCase
type CategoryUseCase interface{}

// CategoryController
type CategoryController struct {
	categoryUseCase CategoryUseCase
	authMiddleware  middlewares.AuthMiddleware
}

// NewCategoryController
func NewCategoryController(
	categoryUseCase CategoryUseCase,
	authMiddleware middlewares.AuthMiddleware,
) CategoryController {
	return CategoryController{
		categoryUseCase: categoryUseCase,
		authMiddleware:  authMiddleware,
	}
}

// RegisterRoutes
func (ctrl CategoryController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/category")
	{
		group.GET("")
		group.Use(ctrl.authMiddleware.AuthGuard).POST("")
		group.Use(ctrl.authMiddleware.AuthGuard).PATCH("")
	}
}
