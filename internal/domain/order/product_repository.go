package order

import "context"

//go:generate mockgen -source $GOFILE -destination "product_repository_mock.go" -package "order"

// ProductRepository
type ProductRepository interface {
	BulkCreateTx(ctx context.Context) (Products, error)
}
