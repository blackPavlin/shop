package cart

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/domain/user"
)

// ID
type ID int64

// Cart
type Cart struct {
	ID     ID
	UserID user.ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Carts
type Carts []*Cart

// Props
type Props struct {
	ProductID product.ID
	Amount    int64
}
