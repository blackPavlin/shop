package order

import (
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

// Service
type Service interface{}

// UseCase
type UseCase struct {
	orderRepo        Repository
	productRepo      product.Repository
	orderProductRepo ProductRepository

	txManager repositoryx.TxManager
}

// NewUseCase
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
