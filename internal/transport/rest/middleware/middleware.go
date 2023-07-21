// Package middleware contains rest middlewares
package middleware

import "github.com/blackPavlin/shop/internal/domain/auth"

// Middleware represents rest middleware.
type Middleware struct {
	authService auth.Service
}

// NewMiddleware create instance of Middleware.
func NewMiddleware(authService auth.Service) *Middleware {
	return &Middleware{authService: authService}
}
