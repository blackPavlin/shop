package middlewares

import (
	"net/http"
	"strings"

	"github.com/blackPavlin/shop/app/internal/domain/usecase/auth"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authUseCase AuthUseCase
}

type AuthUseCase interface {
	ParseAuthToken(acessToken string) (*auth.UserClaims, error)
	CheckPermissions(role string) bool
}

func NewAuthMiddleware(authUseCase AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		authUseCase: authUseCase,
	}
}

func (a *AuthMiddleware) AuthGuard(c *gin.Context) {
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

	claim, err := a.authUseCase.ParseAuthToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("UserID", claim.UserID)
	c.Set("UserRole", claim.UserRole)
}

func (a *AuthMiddleware) RoleGuard(c *gin.Context) {
	userRole := c.GetString("UserRole")

	if ok := a.authUseCase.CheckPermissions(userRole); !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
