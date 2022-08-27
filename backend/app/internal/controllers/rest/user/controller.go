package user

import (
	"context"
	"net/http"

	"github.com/blackPavlin/shop/app/internal/controllers/rest"
	"github.com/blackPavlin/shop/app/internal/controllers/rest/middlewares"
	"github.com/blackPavlin/shop/app/internal/core/entities"
	"github.com/gin-gonic/gin"
)

// UserUseCase
type UserUseCase interface {
	GetUserByID(ctx context.Context, id entities.UserID) (*entities.User, error)
}

// UserController
type UserController struct {
	userUseCase    UserUseCase
	authMiddleware middlewares.AuthMiddleware
}

// NewUserController
func NewUserController(
	userUseCase UserUseCase,
	authMiddleware middlewares.AuthMiddleware,
) UserController {
	return UserController{
		userUseCase:    userUseCase,
		authMiddleware: authMiddleware,
	}
}

// RegisterRoutes
func (ctrl UserController) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/user")
	{
		group.Use(ctrl.authMiddleware.AuthGuard).GET("", ctrl.getUser)
	}
}

// getUser
func (ctrl UserController) getUser(c *gin.Context) {
	userID, err := ctrl.authMiddleware.GetUserID(c)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	user, err := ctrl.userUseCase.GetUserByID(c, userID)
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rest.DomainUserToResponse(user))
}
