package v1

import (
	"context"

	usecase_dto "github.com/blackPavlin/shop/app/internal/domain/usecase/product/dto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase ProductUseCase
}

type ProductUseCase interface {
	CreateProduct(ctx context.Context, dto usecase_dto.CreateProductDTO) error
}

func NewProductHandler(productUseCase ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (h *ProductHandler) Register(router *gin.RouterGroup) {
	group := router.Group("product")
	{
		group.GET("/")
		group.GET("/:product_id", h.GetProductByID)
		group.POST("/", h.CreateProduct)
		group.PUT("/:product_id", h.UpdateProduct)
		group.DELETE("/:product_id", h.DeleteProduct)
	}
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {}

func (h *ProductHandler) CreateProduct(c *gin.Context) {}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {}
