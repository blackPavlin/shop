// Package auth contains authorization oriented logic.
package auth

import (
	"github.com/golang-jwt/jwt"

	"github.com/blackPavlin/shop/internal/domain/user"
)

// UserClaims represents jwt claims.
type UserClaims struct {
	jwt.StandardClaims
	UserID   user.ID   `json:"userId"`
	UserRole user.Role `json:"userRole"`
}

// LoginProps represents login props.
type LoginProps struct {
	Email    string
	Password string
}

// SignupProps represents signup props.
type SignupProps struct {
	Email    string
	Name     string
	Phone    string
	Password string
}
