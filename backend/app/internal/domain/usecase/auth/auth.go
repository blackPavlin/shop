package auth

import (
	"context"
	"errors"
	"time"

	"github.com/blackPavlin/shop/app/internal/config"
	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/auth/dto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAllreadyExists       = errors.New("User allready exists")
	ErrIncorrectLoginOrPassword = errors.New("Incorrect login or password")
	ErrUnauthorized             = errors.New("ErrUnauthorized")
)

type AuthUseCase struct {
	userService UserService
	cartService CartService
	config      config.AuthConfig
}

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User) ([12]byte, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}

type CartService interface {
	CreateCart(ctx context.Context, cart *entities.Cart) error
}

type UserClaims struct {
	jwt.StandardClaims
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

func NewAuthUseCase(userService UserService, cartService CartService, config config.AuthConfig) *AuthUseCase {
	return &AuthUseCase{
		userService: userService,
		cartService: cartService,
		config:      config,
	}
}

func (a *AuthUseCase) Login(ctx context.Context, dto *dto.LoginDTO) (string, error) {
	user, err := a.userService.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ErrIncorrectLoginOrPassword
	}

	if isValidPassword := checkPasswordHash(dto.Password, user.Password); !isValidPassword {
		return "", ErrIncorrectLoginOrPassword
	}

	token, err := signToken(user.ID.Hex(), user.Role, a.config.SigningKey, a.config.ExpiresIn)
	if err != nil {
		return "", ErrIncorrectLoginOrPassword
	}

	return token, nil
}

func (a *AuthUseCase) Registration(ctx context.Context, dto *dto.RegistartionDTO) error {
	user, err := a.userService.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return ErrUserAllreadyExists
	}

	passwordHash, err := hashPassword(dto.Password)
	if err != nil {
		return err
	}

	u := &entities.User{
		Name:     dto.Name,
		Phone:    dto.Phone,
		Address:  dto.Address,
		Email:    dto.Email,
		Role:     entities.UserRole,
		Password: passwordHash,
	}

	userID, err := a.userService.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	c := &entities.Cart{
		Products: []entities.CartProduct{},
		UserID:   userID,
	}

	return a.cartService.CreateCart(ctx, c)
}

func (a *AuthUseCase) ParseAuthToken(acessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(acessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnauthorized
		}

		return []byte(a.config.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok && !token.Valid {
		return nil, ErrUnauthorized
	}

	return claims, nil
}


func (a *AuthUseCase) CheckPermissions(role string) bool {
	return role == entities.AdminRole
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}

func signToken(userID string, userRole string, signingKey string, expiresIn int) (string, error) {
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID:   userID,
		UserRole: userRole,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
