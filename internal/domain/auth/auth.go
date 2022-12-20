package auth

import (
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/golang-jwt/jwt"
)

// UserClaims
type UserClaims struct {
	jwt.StandardClaims
	UserID   user.ID       `json:"userId"`
	UserRole user.UserRole `json:"userRole"`
}

// LoginProps
type LoginProps struct {
	Email    string
	Password string
}

// SignupProps
type SignupProps struct {
	Email    string
	Name     string
	Phone    string
	Password string
}
