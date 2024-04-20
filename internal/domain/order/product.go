package order

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/product"
)

// ProductID represents an id for order product.
type ProductID int64

// ProductIDs represents a slice of ProductID.
type ProductIDs []ProductID

// Product represents the order product.
type Product struct {
	ID        ProductID
	OrderID   ID
	ProductID product.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	Props     *ProductProps
}

// Products represents slice of Product.
type Products []*Product

// ProductProps represents order product editable fields.
type ProductProps struct {
	Amount int64
	Price  int64
}

// ToInt64 convert ProductID to int64.
func (id ProductID) ToInt64() int64 {
	return int64(id)
}

// ToInt64 convert slice of IDs to slice int64.
func (ids ProductIDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.ToInt64())
	}
	return result
}
