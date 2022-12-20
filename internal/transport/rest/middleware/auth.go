package middleware

import "github.com/blackPavlin/shop/internal/domain/auth"

// AuthMiddleware
type AuthMiddleware struct {
	authService auth.AuthService
}

// NewAuthMiddleware
func NewAuthMiddleware(authService auth.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}
