package cart

import (
	"context"
	"fmt"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "cart"

// Service represents cart use cases.
type Service interface {
	Save(ctx context.Context, props *Props) (*Cart, error)
	Delete(ctx context.Context, filter *Filter) error
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// UseCase represents cart service.
type UseCase struct {
	cartRepo    Repository
	productRepo product.Repository
}

// NewUseCase create instance of UseCase.
func NewUseCase(cartRepo Repository, productRepo product.Repository) *UseCase {
	return &UseCase{cartRepo: cartRepo, productRepo: productRepo}
}

// Save cart.
func (s *UseCase) Save(ctx context.Context, props *Props) (*Cart, error) {
	p, err := s.productRepo.Get(ctx, &product.Filter{ID: product.IDFilter{
		Eq: product.IDs{props.ProductID}},
	})
	if err != nil {
		return nil, fmt.Errorf("get cart error: %w", err)
	}
	if p.Props.Amount < props.Amount {
		return nil, errorx.NewConflictError(
			"quantity of goods in the basket must not exceed the quantity in stock")
	}
	result, err := s.cartRepo.Save(ctx, props)
	if err != nil {
		return nil, fmt.Errorf("create cart error: %w", err)
	}
	return result, nil
}

// Delete carts.
func (s *UseCase) Delete(ctx context.Context, filter *Filter) error {
	if err := s.cartRepo.Delete(ctx, filter); err != nil {
		return fmt.Errorf("delete carts error: %w", err)
	}
	return nil
}

// Query carts.
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	result, err := s.cartRepo.Query(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("query carts error: %w", err)
	}
	return result, nil
}
