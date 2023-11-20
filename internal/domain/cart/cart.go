// Package cart contains user cart oriented logic.
package cart

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/domain/user"
)

// ID represents an id for cart.
type ID int64

// IDs defines a slice of ID.
type IDs []ID

// Cart represents the cart.
type Cart struct {
	ID     ID
	UserID user.ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props   *Props
	Product *product.Product
}

// Carts slice of Cart.
type Carts []*Cart

// Props represents cart editable fields.
type Props struct {
	ProductID product.ID
	Amount    int64
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
