package cart

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
	"github.com/blackPavlin/shop/app/internal/core/errs"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "cart"

// CartService
type CartService interface {
	FindByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
	DeleteProducts(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
	DeleteProductByID(ctx context.Context, userID entities.UserID, id entities.ProductID) (*entities.Cart, error)
	AppendProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error)
	UpdateProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error)
}

// ProductService
type ProductService interface {
	FindByID(ctx context.Context, id entities.ProductID) (*entities.Product, error)
}

// CartUseCase
type CartUseCase struct {
	cartService    CartService
	productService ProductService
}

// NewCartUseCase
func NewCartUseCase(cartService CartService, productService ProductService) CartUseCase {
	return CartUseCase{
		cartService:    cartService,
		productService: productService,
	}
}

// GetCartByUserID
func (c CartUseCase) GetCartByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	return c.cartService.FindByUserID(ctx, userID)
}

// ClearCartByUserID
func (c CartUseCase) ClearCartByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	return c.cartService.DeleteProducts(ctx, userID)
}

// AppendProduct
func (c CartUseCase) AppendProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error) {
	if product.Amount <= 0 {
		return c.cartService.DeleteProductByID(ctx, userID, entities.ProductID(product.ProductID))
	}

	p, err := c.productService.FindByID(ctx, entities.ProductID(product.ProductID))
	if err != nil {
		return nil, err
	}

	if p.Amount < product.Amount {
		return nil, errs.ErrTooManyProducts
	}

	return c.cartService.AppendProduct(ctx, userID, product)
}

// UpdateProduct
func (c CartUseCase) UpdateProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error) {
	if product.Amount <= 0 {
		return c.cartService.DeleteProductByID(ctx, userID, entities.ProductID(product.ProductID))
	}

	p, err := c.productService.FindByID(ctx, entities.ProductID(product.ProductID))
	if err != nil {
		return nil, err
	}

	if p.Amount < product.Amount {
		return nil, errs.ErrTooManyProducts
	}

	return c.cartService.UpdateProduct(ctx, userID, product)
}
