package v1

import "github.com/gin-gonic/gin"

type OrderHandler struct {
	orderUseCase OrderUseCase
}

type OrderUseCase interface{}

func NewOrderHandler(orderUseCase OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

func (h *OrderHandler) Register(router *gin.RouterGroup) {}
