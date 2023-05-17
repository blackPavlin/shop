package product

import (
	"context"
	"database/sql"

	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// Service
type Service interface {
	Create(ctx context.Context, props *Props) (*Product, error)
	Get(ctx context.Context, filter *Filter) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// UseCase
type UseCase struct {
	productRepo Repository
	imageRepo   ImageRepository
	txManager   repositoryx.TxManager
}

// NewUseCase
func NewUseCase(
	productRepo Repository,
	imageRepo ImageRepository,
	txManager repositoryx.TxManager,
) *UseCase {
	return &UseCase{productRepo: productRepo, imageRepo: imageRepo, txManager: txManager}
}

// Create
func (s *UseCase) Create(ctx context.Context, props *Props) (*Product, error) {
	var (
		product *Product
		err error
	)
	err = s.txManager.RunTransaction(ctx, &sql.TxOptions{}, func(ctx context.Context) error {
		product, err = s.productRepo.CreateTx(ctx, props)
		if err != nil {
			return err
		}
		for _, image := range props.Images {
			image.Props.ProductID = product.ID
		}
		product.Props.Images, err = s.imageRepo.BulkCreateTx(ctx, props.Images)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get
func (s *UseCase) Get(ctx context.Context, filter *Filter) (*Product, error) {
	return s.productRepo.Get(ctx, filter)
}

// Query
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error) {
	return s.productRepo.Query(ctx, criteria)
}
