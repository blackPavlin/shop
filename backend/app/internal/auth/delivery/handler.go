package delivery

import (
	"net/http"

	"github.com/blackPavlin/shop/app/internal/auth"
	"github.com/blackPavlin/shop/app/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/signin", h.Signin())
		group.POST("/signup", h.Signup())
	}
}

func (h *Handler) Signin() gin.HandlerFunc {
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(c *gin.Context) {
		r := &Request{}

		if err := c.ShouldBind(r); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := h.useCase.Signin(c, r.Email, r.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (h *Handler) Signup() gin.HandlerFunc {
	type Request struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Address  string `json:"address" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(c *gin.Context) {
		r := &Request{}

		if err := c.ShouldBind(r); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := &models.User{
			Name:     r.Name,
			Phone:    r.Phone,
			Address:  r.Address,
			Email:    r.Email,
			Password: r.Password,
		}
		
		if err := h.useCase.Signup(c, user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"success": true})
	}
}
