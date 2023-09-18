package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "auth"

const passwordCost = 14

// Service represents auth use cases.
type Service interface {
	Login(ctx context.Context, props *LoginProps) (string, error)
	Signup(ctx context.Context, props *SignupProps) (string, error)
	ValidateToken(accessToken string) (*UserClaims, error)
}

// UseCase represents auth service.
type UseCase struct {
	logger   *zap.Logger
	config   *config.AuthConfig
	userRepo user.Repository
}

// NewUseCase create instance of UseCase.
func NewUseCase(logger *zap.Logger, conf *config.AuthConfig, userRepo user.Repository) *UseCase {
	return &UseCase{logger, conf, userRepo}
}

// Login user.
func (s *UseCase) Login(ctx context.Context, props *LoginProps) (string, error) {
	usr, err := s.userRepo.Get(ctx, &user.Filter{Email: user.EmailFilter{
		Eq: []string{strings.ToLower(props.Email)}},
	})
	if err != nil {
		if errors.Is(err, errorx.ErrNotFound) {
			return "", errorx.ErrInvalidLoginOrPassword
		}
		return "", errorx.ErrInternal
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(usr.Props.Password),
		[]byte(props.Password),
	); err != nil {
		return "", errorx.ErrInvalidLoginOrPassword
	}
	token, err := s.SignToken(usr)
	if err != nil {
		return "", errorx.ErrInternal
	}
	return token, nil
}

// Signup user.
func (s *UseCase) Signup(ctx context.Context, props *SignupProps) (string, error) {
	exist, err := s.userRepo.Exist(ctx, &user.Filter{Email: user.EmailFilter{
		Eq: []string{strings.ToLower(props.Email)}},
	})
	if err != nil {
		return "", fmt.Errorf("check exists user error: %w", err)
	}
	if exist {
		return "", errorx.ErrAlreadyExists
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(props.Password), passwordCost)
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
		return "", fmt.Errorf("create user error: %w", err)
	}
	loginProps := &LoginProps{
		Email:    props.Email,
		Password: props.Password,
	}
	return s.Login(ctx, loginProps)
}

// SignToken sign authorization token.
func (s *UseCase) SignToken(usr *user.User) (string, error) {
	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.config.ExpiresIn) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:   usr.ID,
		UserRole: usr.Role,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(
		[]byte(s.config.SigningKey),
	)
	if err != nil {
		s.logger.Error("sign auth token error:", zap.Error(err))
		return "", errorx.ErrInternal
	}
	return token, nil
}

// ValidateToken check and parse authorization token.
func (s *UseCase) ValidateToken(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
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
