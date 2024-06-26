// Package product contains product oriented logic.
package product

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/category"
)

// ID represents an id for product.
type ID int64

// IDs represents a slice of ID.
type IDs []ID

// Product represents the product.
type Product struct {
	ID        ID
	CreatedAt time.Time
	UpdatedAt time.Time
	Images    Images
	Props     *Props
}

// Products represents slice of Product.
type Products []*Product

// Props is a user-defined product properties.
type Props struct {
	CategoryID  category.ID
	Name        string
	Description string
	Amount      int64
	Price       int64
}

// ToInt64 convert ID to int64.
func (id ID) ToInt64() int64 {
	return int64(id)
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.ToInt64())
	}
	return result
}
