package v1

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/http/dto"
	usecase_dto "github.com/blackPavlin/shop/app/internal/domain/usecase/product/dto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase ProductUseCase
	authMiddleware AuthMiddleware
}

type ProductUseCase interface {
	CreateProduct(ctx context.Context, dto *usecase_dto.CreateProductDTO) error
	GetProductsByCategory(ctx context.Context, dto *usecase_dto.GetProductDTO) ([]*usecase_dto.GetProductOutDTO, error)
	GetProductByID(ctx context.Context, id string) (*usecase_dto.GetProductOutDTO, error)
}

func NewProductHandler(productUseCase ProductUseCase, authMiddleware AuthMiddleware) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
		authMiddleware: authMiddleware,
	}
}

func (h *ProductHandler) Register(router *gin.RouterGroup) {
	group := router.Group("/product")
	{
		group.GET("/", h.GetProductsByCategory)
		group.GET("/:product_id", h.GetProductByID)
		group.Use(h.authMiddleware.AuthGuard, h.authMiddleware.RoleGuard).POST("/", h.CreateProduct)
		group.Use(h.authMiddleware.AuthGuard, h.authMiddleware.RoleGuard).PUT("/:product_id", h.UpdateProduct)
		group.Use(h.authMiddleware.AuthGuard, h.authMiddleware.RoleGuard).DELETE("/:product_id", h.DeleteProduct)
	}
}

// GetProductsByCategory - получить список товаров в категории
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	query := &dto.GetProductQueryDTO{}

	if err := c.BindQuery(query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecaseDTO := &usecase_dto.GetProductDTO{
		Category: query.Category,
		Limit:    query.Limit,
		Offset:   query.Offset,
	}

	products, err := h.productUseCase.GetProductsByCategory(c, usecaseDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductByID - получить товар по ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("product_id")

	product, err := h.productUseCase.GetProductByID(c, productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// CreateProduct - создать товар
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	r := &dto.CreateProductDTO{}

	if err := c.ShouldBind(r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecaseDTO := &usecase_dto.CreateProductDTO{
		Name:        r.Name,
		Description: r.Description,
		Category:    r.Category,
		Quantity:    r.Quantity,
	}

	if err := h.productUseCase.CreateProduct(c, usecaseDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {}
