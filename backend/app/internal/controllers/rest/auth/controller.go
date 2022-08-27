package auth

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/rest"
	"github.com/blackPavlin/shop/app/internal/core/entities"
	"github.com/blackPavlin/shop/app/internal/core/usecases/auth"
	"github.com/gin-gonic/gin"
)

// AuthUseCase
type AuthUseCase interface {
	Login(ctx context.Context, dto *auth.LoginUserDTO) (string, error)
	Signup(ctx context.Context, dto *auth.SignupUserDTO) (entities.UserID, error)
}

// AuthController
type AuthController struct {
	authUseCase AuthUseCase
}

// NewAuthController
func NewAuthController(authUseCase AuthUseCase) AuthController {
	return AuthController{
		authUseCase: authUseCase,
	}
}

// RegisterRoutes
func (ctrl AuthController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/login", ctrl.login)
		group.POST("/signup", ctrl.signup)
	}
}

// login
func (ctrl AuthController) login(c *gin.Context) {
	req := &rest.LoginRequest{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginDTO := &auth.LoginUserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := ctrl.authUseCase.Login(c, loginDTO)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	res := rest.LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, res)
}

// signup
func (ctrl AuthController) signup(c *gin.Context) {
	req := &rest.SignupRequest{}

	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	signupDTO := &auth.SignupUserDTO{
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: req.Password,
	}

	if _, err := ctrl.authUseCase.Signup(c, signupDTO); err != nil {
		rest.HandleError(c, err)
		return
	}

	loginDTO := &auth.LoginUserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := ctrl.authUseCase.Login(c, loginDTO)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	res := rest.LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, res)
}
