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
	product, err := s.productRepo.Get(ctx, &product.Filter{ID: product.IDFilter{
		Eq: product.IDs{props.ProductID}},
	})
	if err != nil {
		return nil, err
	}
	if product.Props.Amount < props.Amount {
		// todo: return error
	}
	return s.cartRepo.Create(ctx, props)
}

// Update
// func (s *UseCase) Update(ctx context.Context, cart *Cart) (*Cart, error) {}

// Query
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	return s.cartRepo.Query(ctx, criteria)
}
