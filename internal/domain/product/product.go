package product

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/category"
)

// ID
type ID int64

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
