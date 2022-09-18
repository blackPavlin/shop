package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/blackPavlin/shop/app/internal/config"
	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/blackPavlin/shop/app/internal/errs"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "auth"

// UserService
type UserService interface {
	Create(ctx context.Context, user *entities.User) (entities.UserID, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

// CartService
type CartService interface {
	Create(ctx context.Context, cart *entities.Cart) (entities.CartID, error)
}

// TransactionService
type TransactionService interface {
	RunTransaction(ctx context.Context, callback func(context context.Context) error) error
}

// AuthUseCase
type AuthUseCase struct {
	userService        UserService
	cartService        CartService
	transactionService TransactionService
	authConfig         config.AuthConfig
}

// UserClaims
type UserClaims struct {
	jwt.StandardClaims
	UserID   string            `json:"userId"`
	UserRole entities.UserRole `json:"userRole"`
}

// NewAuthUseCase
func NewAuthUseCase(
	userService UserService,
	cartService CartService,
	transactionService TransactionService,
	authConfig config.AuthConfig,
) AuthUseCase {
	return AuthUseCase{
		userService:        userService,
		cartService:        cartService,
		transactionService: transactionService,
		authConfig:         authConfig,
	}
}

// Login
func (a AuthUseCase) Login(ctx context.Context, dto *LoginUserDTO) (string, error) {
	user, err := a.userService.FindByEmail(ctx, strings.ToLower(dto.Email))
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return "", errs.ErrInvalidLoginOrPassword
		}

		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return "", errs.ErrInvalidLoginOrPassword
	}

	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(a.authConfig.ExpiresIn) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID:   user.ID.Hex(),
		UserRole: user.Role,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(a.authConfig.SigningKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

// Signup
func (a AuthUseCase) Signup(ctx context.Context, dto *SignupUserDTO) (entities.UserID, error) {
	exists, err := a.userService.ExistsByEmail(ctx, strings.ToLower(dto.Email))
	if err != nil {
		return entities.UserID{}, err
	}

	if exists {
		return entities.UserID{}, errs.ErrUserAllreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 14)
	if err != nil {
		return entities.UserID{}, err
	}

	user := &entities.User{
		Name:     dto.Name,
		Phone:    dto.Phone,
		Email:    strings.ToLower(dto.Email),
		Role:     entities.RoleUser,
		Password: string(passwordHash),
	}

	var userID entities.UserID

	if err := a.transactionService.RunTransaction(ctx, func(sctx context.Context) error {
		userID, err = a.userService.Create(sctx, user)
		if err != nil {
			return err
		}

		cart := &entities.Cart{
			UserID:   primitive.ObjectID(userID),
			Products: []entities.CartProduct{},
		}

		if _, err := a.cartService.Create(sctx, cart); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return entities.UserID{}, err
	}

	return userID, nil
}

// ValidateToken
func (a AuthUseCase) ValidateToken(ctx context.Context, token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrUnauthorized
		}

		return []byte(a.authConfig.SigningKey), nil
	})
	if err != nil {
		// TODO: Log
		return nil, errs.ErrUnauthorized
	}

	if !t.Valid {
		return nil, errs.ErrUnauthorized
	}

	claims, ok := t.Claims.(*UserClaims)
	if !ok {
		return nil, errs.ErrUnauthorized
	}

	return claims, nil
}
