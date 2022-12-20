package cart

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/product"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "cart"

// CartService
type CartService interface {
	Create(ctx context.Context, props *Props) (*Cart, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// CartUseCase
type CartUseCase struct {
	cartRepo    Repository
	productRepo product.Repository
}

// NewCartUseCase
func NewCartUseCase(cartRepo Repository, productRepo product.Repository) *CartUseCase {
	return &CartUseCase{cartRepo: cartRepo, productRepo: productRepo}
}

// Create
func (s *CartUseCase) Create(ctx context.Context, props *Props) (*Cart, error) {
	// TODO: check product count

	return s.cartRepo.Create(ctx, props)
}

// Query
func (s *CartUseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	return s.cartRepo.Query(ctx, criteria)
}
