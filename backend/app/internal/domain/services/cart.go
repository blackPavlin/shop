package services

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

type CartService struct {
	cartRepo CartRepository
}

type CartRepository interface {
	Create(ctx context.Context, cart *entities.Cart) error
	FindByUserID(ctx context.Context, userID string) (*entities.Cart, error)
	AddProduct(ctx context.Context, userID string, product *entities.CartProduct) error
}

func NewCartService(cartRepo CartRepository) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (c *CartService) CreateCart(ctx context.Context, cart *entities.Cart) error {
	return c.cartRepo.Create(ctx, cart)
}

func (c *CartService) GetCartByUserID(ctx context.Context, userID string) (*entities.Cart, error) {
	return c.cartRepo.FindByUserID(ctx, userID)
}

func (c *CartService) AddProductToCart(ctx context.Context, userID string, product *entities.CartProduct) error {
	return c.cartRepo.AddProduct(ctx, userID, product)
}

func (c *CartService) RemoveProductFromCart(ctx context.Context, userID string) {}

func (c *CartService) UpdateCountProductsFromCart(ctx context.Context, userID string) {}
