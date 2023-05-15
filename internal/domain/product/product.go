package product

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/category"
)

// ID
type ID int64

// IDs
type IDs []ID

// Product
type Product struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Products
type Products []*Product

// Props
type Props struct {
	CategoryID  category.ID
	Name        string
	Description string
	Amount      int64
	Price       int64
}

// ToInt64
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
