package order

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/product"
)

// ProductID
type ProductID int64

// Product
type Product struct {
	ID ProductID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *ProductProps
}

// Products
type Products []Product

// ProductProps
type ProductProps struct {
	OrderID   ID
	ProductID product.ID
	Amount    int64
	Price     int64
}
