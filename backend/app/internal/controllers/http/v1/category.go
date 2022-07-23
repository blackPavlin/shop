package v1

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/http/dto"
	usecase_dto "github.com/blackPavlin/shop/app/internal/domain/usecase/category/dto"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUseCase CategoryUseCase
	authMiddleware  AuthMiddleware
}

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, dto *usecase_dto.CreateCategoryDTO) error
	GetCategoryList(ctx context.Context) ([]string, error)
}

func NewCategoryHandler(categoryUseCase CategoryUseCase, authMiddleware AuthMiddleware) *CategoryHandler {
	return &CategoryHandler{
		categoryUseCase: categoryUseCase,
		authMiddleware:  authMiddleware,
	}
}

func (h *CategoryHandler) Register(router *gin.RouterGroup) {
	group := router.Group("/category")
	{
		group.GET("", h.GetCategories)
		group.Use(h.authMiddleware.AuthGuard, h.authMiddleware.RoleGuard).POST("", h.CreateCategory)
	}
}

// GetCategories - Получить список категорий
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryUseCase.GetCategoryList(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// CreateCategory - Создать категорию
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	r := &dto.CreateCategoryDTO{}

	if err := c.ShouldBind(r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecaseDTO := &usecase_dto.CreateCategoryDTO{
		Name: r.Name,
	}

	if err := h.categoryUseCase.CreateCategory(c, usecaseDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}
