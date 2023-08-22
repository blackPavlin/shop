package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateCartResponse transform domain entity to rest response.
func CreateCartResponse(c *cart.Cart) rest.Cart {
	return rest.Cart{
		Id:      int64(c.ID),
		UserId:  int64(c.UserID),
		Amount:  c.Props.Amount,
		Product: CreateProductResponse(c.Product),
	}
}

// CreateCartListResponse transform domain entity to rest response.
func CreateCartListResponse(carts cart.Carts) rest.CartList {
	result := make(rest.CartList, 0, len(carts))
	for _, c := range carts {
		result = append(result, CreateCartResponse(c))
	}
	return result
}
