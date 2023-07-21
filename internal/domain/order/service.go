package order

import (
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

// Service represents order use cases.
type Service interface{}

// UseCase represents order service.
type UseCase struct {
	orderRepo        Repository
	productRepo      product.Repository
	orderProductRepo ProductRepository

	txManager repositoryx.TxManager
}

// NewUseCase  create instance of UseCase.
func NewUseCase(
	orderRepo Repository,
	productRepo product.Repository,
	orderProductRepo ProductRepository,
	txManager repositoryx.TxManager,
) *UseCase {
	return &UseCase{
		orderRepo:        orderRepo,
		productRepo:      productRepo,
		orderProductRepo: orderProductRepo,
		txManager:        txManager,
	}
}
