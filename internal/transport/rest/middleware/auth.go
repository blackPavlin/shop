package middleware

import "github.com/blackPavlin/shop/internal/domain/auth"

// AuthMiddleware
type AuthMiddleware struct {
	authService auth.Service
}

// NewAuthMiddleware
func NewAuthMiddleware(authService auth.Service) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}
