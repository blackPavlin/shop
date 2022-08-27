package middlewares

import (
	"net/http"
	"strings"

	"github.com/blackPavlin/shop/app/internal/controllers/rest"
	"github.com/blackPavlin/shop/app/internal/core/entities"
	"github.com/blackPavlin/shop/app/internal/core/usecases/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthMiddleware
type AuthMiddleware struct {
	authUseCase auth.AuthUseCase
}

// NewAuthMiddleware
func NewAuthMiddleware(authUseCase auth.AuthUseCase) AuthMiddleware {
	return AuthMiddleware{
		authUseCase: authUseCase,
	}
}

// AuthGuard
func (a AuthMiddleware) AuthGuard(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claim, err := a.authUseCase.ValidateToken(c, headerParts[1])
	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.Set("UserID", claim.UserID)
	c.Set("UserRole", claim.UserRole)
}

// GetUserID
func (a AuthMiddleware) GetUserID(c *gin.Context) (entities.UserID, error) {
	userID, err := primitive.ObjectIDFromHex(c.GetString("UserID"))
	if err != nil {
		return entities.UserID{}, err
	}

	return entities.UserID(userID), nil
}
