package cart

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

var (
	ErrCartNotFound    = errors.New("Cart not found")
	ErrProductNotFound = errors.New("Product not found")
	Err                = errors.New("")
)

type CartUseCase struct {
	cartService    CartService
	productService ProductService
}

type CartService interface {
	GetCartByUserID(ctx context.Context, userID string) (*entities.Cart, error)
	AddProductToCart(ctx context.Context, userID string, product *entities.CartProduct) error
}

type ProductService interface {
	GetProductByID(ctx context.Context, id string) (*entities.Product, error)
}

func NewCartUseCase(cartService CartService, productService ProductService) *CartUseCase {
	return &CartUseCase{
		cartService:    cartService,
		productService: productService,
	}
}

func (c *CartUseCase) GetCart(ctx context.Context, userID string) (*entities.Cart, error) {
	cart, err := c.cartService.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return nil, ErrCartNotFound
	}

	return cart, nil
}

func (c *CartUseCase) AddProductToCart(ctx context.Context, userID string, productID string, productsCount int32) error {
	product, err := c.productService.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	if product.Count < productsCount {
		return Err
	}

	if _, err := c.GetCart(ctx, userID); err != nil {
		return err
	}

	cartProduct := &entities.CartProduct{
		ProductID: product.ID,
		Count:     productsCount,
	}

	return c.cartService.AddProductToCart(ctx, userID, cartProduct)
}
