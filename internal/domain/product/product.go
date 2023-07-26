// Package product contains product oriented logic.
package product

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/category"
)

// ID represents an id for product.
type ID int64

// IDs defines a slice of ID.
type IDs []ID

// Product represents the product.
type Product struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Images Images
	Props  *Props
}

// Products slice of Product.
type Products []*Product

// Props is a user-defined product properties.
type Props struct {
	CategoryID  category.ID
	Name        string
	Description string
	Amount      int64
	Price       int64
	Images      Images
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
