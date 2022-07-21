package v1

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/http/dto"
	usecase_dto "github.com/blackPavlin/shop/app/internal/domain/usecase/auth/dto"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase AuthUseCase
}

type AuthUseCase interface {
	Login(ctx context.Context, dto *usecase_dto.LoginDTO) (string, error)
	Registration(ctx context.Context, dto *usecase_dto.RegistartionDTO) error
}

type AuthMiddleware interface {
	AuthGuard(c *gin.Context)
	RoleGuard(c *gin.Context)
}

func NewAuthHandler(authUseCase AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) Register(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/login", h.Login)
		group.POST("/registration", h.Registration)
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	r := &dto.LoginDTO{}

	if err := c.ShouldBind(r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecaseDTO := &usecase_dto.LoginDTO{
		Email:    r.Email,
		Password: r.Password,
	}

	token, err := h.authUseCase.Login(c, usecaseDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Registration(c *gin.Context) {
	r := &dto.RegistartionDTO{}

	if err := c.ShouldBind(r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecaseDTO := &usecase_dto.RegistartionDTO{
		Name:     r.Name,
		Phone:    r.Phone,
		Address:  r.Address,
		Email:    r.Email,
		Password: r.Password,
	}

	if err := h.authUseCase.Registration(c, usecaseDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}
