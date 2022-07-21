package v1

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase    UserUseCase
	authMiddleware AuthMiddleware
}

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
}

func NewUserHandler(userUseCase UserUseCase, authMiddleware AuthMiddleware) *UserHandler {
	return &UserHandler{
		userUseCase:    userUseCase,
		authMiddleware: authMiddleware,
	}
}

func (h *UserHandler) Register(router *gin.RouterGroup) {
	group := router.Group("user")
	{
		group.Use(h.authMiddleware.AuthGuard).GET("", h.GetUser)
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
