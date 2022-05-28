package delivery

import (
	"net/http"

	"github.com/blackPavlin/shop/app/internal/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase user.UseCase
}

func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	group := router.Group("/user")
	{
		group.GET("", h.GetUser())
	}
}

func (h *Handler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("Email")

		u, err := h.useCase.FindByEmail(c, email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if u == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": user.UserNotFound})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": u})
	}
}
