package order

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/product"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "product_repository_mock.go" -package "order"

// ProductRepository represents order product repository.
type ProductRepository interface {
	BulkCreateTx(ctx context.Context, products Products) (Products, error)
}

// ProductQueryCriteria represents a criteria for order product query.
type ProductQueryCriteria struct {
	Filter ProductFilter
}

// ProductFilter represents order product filter.
type ProductFilter struct {
	ID        ProductIDFilter
	OrderID   IDFilter
	ProductID product.IDFilter
}

// ProductIDFilter represents ProductID filter.
type ProductIDFilter struct {
	Eq  ProductIDs
	Neq ProductIDs
}
