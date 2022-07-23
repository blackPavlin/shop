package v1

import "github.com/gin-gonic/gin"

type OrderHandler struct {
	orderUseCase   OrderUseCase
	authMiddleware AuthMiddleware
}

type OrderUseCase interface{}

func NewOrderHandler(orderUseCase OrderUseCase, authMiddleware AuthMiddleware) *OrderHandler {
	return &OrderHandler{
		orderUseCase:   orderUseCase,
		authMiddleware: authMiddleware,
	}
}

func (h *OrderHandler) Register(router *gin.RouterGroup) {
	group := router.Group("/order")
	{
		group.Use(h.authMiddleware.AuthGuard).GET("", h.GetOrders)
		group.Use(h.authMiddleware.AuthGuard).GET("/:order_id", h.GetOrderByID)
		group.Use(h.authMiddleware.AuthGuard).POST("", h.CreateOrder)
	}
}

func (h *OrderHandler) GetOrders(c *gin.Context) {}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {}

func (h *OrderHandler) CreateOrder(c *gin.Context) {}
