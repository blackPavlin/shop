package cart

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/product"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "cart"

// Service
type Service interface {
	Create(ctx context.Context, props *Props) (*Cart, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// UseCase
type UseCase struct {
	cartRepo    Repository
	productRepo product.Repository
}

// NewUseCase
func NewUseCase(cartRepo Repository, productRepo product.Repository) *UseCase {
	return &UseCase{cartRepo: cartRepo, productRepo: productRepo}
}

// Create
func (s *UseCase) Create(ctx context.Context, props *Props) (*Cart, error) {
	// TODO: check product count

	return s.cartRepo.Create(ctx, props)
}

// Query
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	return s.cartRepo.Query(ctx, criteria)
}
