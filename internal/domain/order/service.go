package order

import (
	"context"
	"fmt"

	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "service_mock.go" -package "order"

// Service represents order use cases.
type Service interface {
	Get(ctx context.Context, filter *Filter) (*Order, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
	Create(ctx context.Context) (*Order, error)
}

// UseCase represents order service.
type UseCase struct {
	repository             Repository
	cartRepository         cart.Repository
	productRepository      product.Repository
	orderProductRepository ProductRepository
	txManager              repositoryx.TxManager
}

// NewUseCase  create instance of UseCase.
func NewUseCase(
	repository Repository,
	cartRepository cart.Repository,
	productRepository product.Repository,
	orderProductRepository ProductRepository,
	txManager repositoryx.TxManager,
) *UseCase {
	return &UseCase{
		repository:             repository,
		cartRepository:         cartRepository,
		productRepository:      productRepository,
		orderProductRepository: orderProductRepository,
		txManager:              txManager,
	}
}

// Get order.
func (uc UseCase) Get(ctx context.Context, filter *Filter) (*Order, error) {
	result, err := uc.repository.Get(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get order error: %w", err)
	}
	return result, nil
}

// Query orders.
func (uc UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	result, err := uc.repository.Query(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("query orders error: %w", err)
	}
	return result, nil
}

// Create order.
func (uc UseCase) Create(ctx context.Context) (*Order, error) {
	return nil, nil
}
