package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/restx"
)

// AuthMiddleware
type AuthMiddleware struct {
	authService auth.Service
}

// NewAuthMiddleware
func NewAuthMiddleware(authService auth.Service) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			render.Render(w, r, restx.ErrorResponse{
				HTTPStatusCode: http.StatusUnauthorized,
				Message:        "Unauthorized",
			})
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 {
			render.Render(w, r, restx.ErrorResponse{
				HTTPStatusCode: http.StatusUnauthorized,
				Message:        "Unauthorized",
			})
			return
		}
		if headerParts[0] != "Bearer" {
			render.Render(w, r, restx.ErrorResponse{
				HTTPStatusCode: http.StatusUnauthorized,
				Message:        "Unauthorized",
			})
			return
		}
		userClaims, err := m.authService.ValidateToken(headerParts[1])
		if err != nil {
			render.Render(w, r, restx.ErrorResponse{
				HTTPStatusCode: http.StatusUnauthorized,
				Message:        "Unauthorized",
			})
			return
		}
		user.SetUser(r.Context(), &user.User{
			ID:   userClaims.UserID,
			Role: userClaims.UserRole,
		})
		next.ServeHTTP(w, r)
	})
}
