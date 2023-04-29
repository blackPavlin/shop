package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "auth"

// AuthService
type AuthService interface {
	Login(ctx context.Context, props *LoginProps) (string, error)
	Signup(ctx context.Context, props *SignupProps) (string, error)
	ValidateToken(accessToken string) (*UserClaims, error)
}

// AuthUseCase
type AuthUseCase struct {
	logger *zap.Logger
	config *config.AuthConfig

	userRepo user.Repository
}

// NewAuthUseCase
func NewAuthUseCase(
	logger *zap.Logger,
	config *config.AuthConfig,
	userRepo user.Repository,
) *AuthUseCase {
	return &AuthUseCase{
		logger:   logger,
		config:   config,
		userRepo: userRepo,
	}
}

// Login
func (s *AuthUseCase) Login(ctx context.Context, props *LoginProps) (string, error) {
	user, err := s.userRepo.Get(ctx, &user.Filter{Email: strings.ToLower(props.Email)})
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return "", errorx.ErrInvalidLoginOrPassword
		}
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Props.Password),
		[]byte(props.Password),
	); err != nil {
		return "", errorx.ErrInvalidLoginOrPassword
	}
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.config.ExpiresIn) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID:   user.ID,
		UserRole: user.Role,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(
		[]byte(s.config.SigningKey),
	)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Signup
func (s *AuthUseCase) Signup(ctx context.Context, props *SignupProps) (string, error) {
	exist, err := s.userRepo.Exist(ctx, &user.Filter{Email: strings.ToLower(props.Email)})
	if err != nil {
		return "", err
	}
	if exist {
		return "", errorx.ErrAllreadyExists
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(props.Password), 14)
	if err != nil {
		s.logger.Error("generate password hash error:", zap.Error(err))
		return "", errorx.ErrInternal
	}
	userProps := &user.Props{
		Email:    strings.ToLower(props.Email),
		Name:     props.Name,
		Phone:    props.Phone,
		Password: string(passwordHash),
	}
	if _, err := s.userRepo.Create(ctx, userProps); err != nil {
		return "", err
	}
	loginProps := &LoginProps{
		Email:    props.Email,
		Password: props.Password,
	}
	return s.Login(ctx, loginProps)
}

// ValidateToken
func (s *AuthUseCase) ValidateToken(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{},func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(s.config.SigningKey), nil
	})
	if err != nil {
		return nil, errorx.ErrUnauthorized
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errorx.ErrUnauthorized
}
