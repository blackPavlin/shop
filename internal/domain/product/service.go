package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// Service represents product use cases.
type Service interface {
	Create(ctx context.Context, props *Props) (*Product, error)
	Update(ctx context.Context, productID ID, props *Props) (*Product, error)
	Get(ctx context.Context, filter *Filter) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// UseCase represents product service.
type UseCase struct {
	productRepo Repository
	txManager   repositoryx.TxManager
}

// NewUseCase create instance of UseCase.
func NewUseCase(
	productRepo Repository,
	txManager repositoryx.TxManager,
) *UseCase {
	return &UseCase{productRepo: productRepo, txManager: txManager}
}

// Create product.
func (s *UseCase) Create(ctx context.Context, props *Props) (*Product, error) {
	var (
		product *Product
		err     error
	)
	err = s.txManager.RunTransaction(ctx, &sql.TxOptions{}, func(ctx context.Context) error {
		product, err = s.productRepo.CreateTx(ctx, props)
		if err != nil {
			return fmt.Errorf("create product error: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create product transaction error: %w", err)
	}
	return product, nil
}

// Update product.
func (s *UseCase) Update(ctx context.Context, productID ID, props *Props) (*Product, error) {
	result, err := s.productRepo.Update(ctx, productID, props)
	if err != nil {
		return nil, fmt.Errorf("update product error: %w", err)
	}
	return result, nil
}

// Get product.
func (s *UseCase) Get(ctx context.Context, filter *Filter) (*Product, error) {
	result, err := s.productRepo.Get(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get product error: %w", err)
	}
	return result, nil
}

// Query products.
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	result, err := s.productRepo.Query(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("query products error: %w", err)
	}
	return result, nil
}
