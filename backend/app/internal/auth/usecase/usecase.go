package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/blackPavlin/shop/app/internal/auth"
	"github.com/blackPavlin/shop/app/internal/config"
	"github.com/blackPavlin/shop/app/internal/models"
	"github.com/blackPavlin/shop/app/internal/user"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userUseCase user.UseCase
	config      *config.Config
}

func NewAuthUseCase(userUseCase user.UseCase, config *config.Config) *AuthUseCase {
	return &AuthUseCase{
		userUseCase: userUseCase,
		config:      config,
	}
}

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

// Signin - Авторизация пользователя
func (a *AuthUseCase) Signin(ctx context.Context, email, password string) (string, error) {
	user, err := a.userUseCase.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New(auth.IncorrectEmailOrPassword)
	}

	if isValidPassword := checkPasswordHash(password, user.Password); !isValidPassword {
		return "", errors.New(auth.IncorrectEmailOrPassword)
	}

	return signToken(email, a.config.Auth.SigningKey, a.config.Auth.ExpiresIn)
}

// Signup - Регистрация пользователя
func (a *AuthUseCase) Signup(ctx context.Context, user *models.User) error {
	u, err := a.userUseCase.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if u != nil {
		return errors.New(auth.UserAllreadyExists)
	}

	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = passwordHash
	user.Role = "user"

	return a.userUseCase.Create(ctx, user)
}

// hashPassword - получить хэш пароля
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash - проверить валидность пароля
func checkPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}

// signToken - получить jwt токен
func signToken(email string, signingKey string, expiresIn int) (string, error) {
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
