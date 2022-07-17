package v1

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase     UserUseCase
	authMiddlewarre AuthMiddleware
}

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
}

type AuthMiddleware interface {
	AuthGuard(c *gin.Context)
}

func NewUserHandler(userUseCase UserUseCase, authMiddlewarre AuthMiddleware) *UserHandler {
	return &UserHandler{
		userUseCase:     userUseCase,
		authMiddlewarre: authMiddlewarre,
	}
}

func (h *UserHandler) Register(router *gin.RouterGroup) {
	group := router.Group("user")
	{
		group.Use(h.authMiddlewarre.AuthGuard).GET("", h.GetUser)
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.GetString("UserID")

	user, err := h.userUseCase.GetUser(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
