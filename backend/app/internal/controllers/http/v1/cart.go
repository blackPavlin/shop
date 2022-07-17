package v1

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUseCase     CartUseCase
	authMiddlewarre AuthMiddleware
}

type CartUseCase interface {
	GetCart(ctx context.Context, userID string) (*entities.Cart, error)
}

func NewCartHandler(cartUseCase CartUseCase, authMiddlewarre AuthMiddleware) *CartHandler {
	return &CartHandler{
		cartUseCase:     cartUseCase,
		authMiddlewarre: authMiddlewarre,
	}
}

func (h *CartHandler) Register(router *gin.RouterGroup) {
	group := router.Group("catr")
	{
		group.Use(h.authMiddlewarre.AuthGuard).GET("", h.GetCart)
		group.Use(h.authMiddlewarre.AuthGuard).POST("", h.AddProduct)
		group.Use(h.authMiddlewarre.AuthGuard).PATCH("/:productID", h.UpdateProductsCount)
		group.Use(h.authMiddlewarre.AuthGuard).DELETE("/:productID", h.DeleteProductFromCart)
	}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetString("UserID")

	cart, err := h.cartUseCase.GetCart(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"catr": cart})
}

func (h *CartHandler) AddProduct(c *gin.Context) {}

func (h *CartHandler) UpdateProductsCount(c *gin.Context) {}

func (h *CartHandler) DeleteProductFromCart(c *gin.Context) {}
