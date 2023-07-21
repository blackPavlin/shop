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

	Props *Props
}

// Carts slice of Cart.
type Carts []*Cart

// Props represents cart editable fields.
type Props struct {
	ProductID product.ID
	Amount    int64
}
