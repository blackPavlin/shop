package category

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/rest"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/middlewares"
	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryUseCase
type CategoryUseCase interface {
	Create(ctx context.Context, category *entities.Category) (*entities.Category, error)
	Update(ctx context.Context, category *entities.Category) (*entities.Category, error)
	Find(ctx context.Context) ([]*entities.Category, error)
}

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
		group.GET("", ctrl.getCategories)
		group.Use(ctrl.authMiddleware.AuthGuard).POST("", ctrl.createCategory)
		group.Use(ctrl.authMiddleware.AuthGuard).PATCH("", ctrl.updateCategory)
	}
}

// getCategories
func (ctrl CategoryController) getCategories(c *gin.Context) {
	categories, err := ctrl.categoryUseCase.Find(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainCategoriesToResponse(categories))
}

// createCategory
func (ctrl CategoryController) createCategory(c *gin.Context) {
	req := &rest.CreateCategoryRequest{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := ctrl.categoryUseCase.Create(c, &entities.Category{Name: req.Name})
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, rest.DomainCategoryToResponse(category))
}

// updateCategory
func (ctrl CategoryController) updateCategory(c *gin.Context) {
	req := &rest.Category{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	category, err := ctrl.categoryUseCase.Update(c, &entities.Category{
		ID:   categoryID,
		Name: req.Name,
	})
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainCategoryToResponse(category))
}
