package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "service_mock.go" -package "product"

// Service represents product use cases.
type Service interface {
	Get(ctx context.Context, filter *Filter) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
	Create(ctx context.Context, props *Props) (*Product, error)
	Update(ctx context.Context, productID ID, props *Props) (*Product, error)
	Delete(ctx context.Context, productID ID) error
}

// UseCase represents product service.
type UseCase struct {
	productRepo      Repository
	productImageRepo ImageRepository
	imageStorage     image.Storage
	txManager        repositoryx.TxManager
}

// NewUseCase create instance of UseCase.
func NewUseCase(
	productRepo Repository,
	productImageRepo ImageRepository,
	imageStorage image.Storage,
	txManager repositoryx.TxManager,
) *UseCase {
	return &UseCase{
		productRepo:      productRepo,
		productImageRepo: productImageRepo,
		imageStorage:     imageStorage,
		txManager:        txManager,
	}
}

// Get product.
func (s UseCase) Get(ctx context.Context, filter *Filter) (*Product, error) {
	result, err := s.productRepo.Get(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get product error: %w", err)
	}
	return result, nil
}

// Query products.
func (s UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	result, err := s.productRepo.Query(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("query products error: %w", err)
	}
	return result, nil
}

// Create product.
func (s UseCase) Create(ctx context.Context, props *Props) (*Product, error) {
	product, err := s.productRepo.Create(ctx, props)
	if err != nil {
		return nil, fmt.Errorf("create product error: %w", err)
	}
	return product, nil
}

// Update product.
func (s UseCase) Update(ctx context.Context, productID ID, props *Props) (*Product, error) {
	result, err := s.productRepo.Update(ctx, productID, props)
	if err != nil {
		return nil, fmt.Errorf("update product error: %w", err)
	}
	return result, nil
}

// Delete product.
func (s UseCase) Delete(ctx context.Context, productID ID) error {
	images, err := s.productImageRepo.Query(ctx, &ImageQueryCriteria{
		Filter: ImageFilter{
			ProductID: IDFilter{Eq: IDs{productID}},
		},
	})
	if err != nil {
		return fmt.Errorf("query product images error: %w", err)
	}
	err = s.txManager.RunTransaction(ctx, &sql.TxOptions{}, func(ctx context.Context) error {
		if err = s.productRepo.Delete(ctx, &Filter{
			ID: IDFilter{Eq: IDs{productID}},
		}); err != nil {
			return fmt.Errorf("delete product error: %w", err)
		}
		if err = s.imageStorage.BulkRemove(ctx, images.Names()); err != nil {
			return fmt.Errorf("bulkRemove images errro: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("delete product transaction error: %w", err)
	}
	return nil
}
