package cart

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

var (
	ErrAmountMustBePositive = errors.New("Product amount must be positive")
)

// CartService
type CartService struct {
	cartRepository CartRepository
}

// NewCartService
func NewCartService(cartRepository CartRepository) CartService {
	return CartService{
		cartRepository: cartRepository,
	}
}

// Create
func (s CartService) Create(ctx context.Context, cart *entities.Cart) (entities.CartID, error) {
	return s.cartRepository.Create(ctx, cart)
}

// FindByUserID
func (s CartService) FindByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	return s.cartRepository.FindByUserID(ctx, userID)
}

// DeleteProducts
func (s CartService) DeleteProducts(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	return s.cartRepository.DeleteProducts(ctx, userID)
}

// DeleteProductByID
func (s CartService) DeleteProductByID(ctx context.Context, userID entities.UserID, id entities.ProductID) (*entities.Cart, error) {
	return s.cartRepository.DeleteProductByID(ctx, userID, id)
}

// AppendProduct
func (s CartService) AppendProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error) {
	if product.Amount <= 0 {
		return nil, ErrAmountMustBePositive
	}

	return s.cartRepository.AppendProduct(ctx, userID, product)
}

// UpdateProduct
func (s CartService) UpdateProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error) {
	if product.Amount <= 0 {
		return nil, ErrAmountMustBePositive
	}

	return s.cartRepository.UpdateProduct(ctx, userID, product)
}
