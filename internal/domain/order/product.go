package order

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/product"
)

// ProductID represents an id for order product.
type ProductID int64

// ProductIDs defines a slice of ProductID.
type ProductIDs []ProductID

// Product represents the order product.
type Product struct {
	ID ProductID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *ProductProps
}

// Products slice of Product.
type Products []*Product

// ProductProps represents order product editable fields.
type ProductProps struct {
	OrderID   ID
	ProductID product.ID
	Amount    int64
	Price     int64
}

// ToInt64 convert slice of IDs to slice int64.
func (ids ProductIDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
