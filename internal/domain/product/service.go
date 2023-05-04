package product

import (
	"context"

	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// Service
type Service interface {
	Create(ctx context.Context, props *Props) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// UseCase
type UseCase struct {
	productRepo Repository
	txManager   repositoryx.TxManager
}

// NewUseCase
func NewUseCase(productRepo Repository, txManager repositoryx.TxManager) *UseCase {
	return &UseCase{productRepo: productRepo, txManager: txManager}
}

// Create
func (s *UseCase) Create(ctx context.Context, props *Props) (*Product, error) {
	return s.productRepo.Create(ctx, props)
}

// Query
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	return s.productRepo.Query(ctx, criteria)
}
