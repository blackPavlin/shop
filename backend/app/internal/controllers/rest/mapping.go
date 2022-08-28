package rest

import "github.com/blackPavlin/shop/app/internal/entities"

// DomainUserToResponse
func DomainUserToResponse(user *entities.User) User {
	return User{
		Id:    user.ID.Hex(),
		Email: user.Email,
		Name:  user.Name,
		Phone: user.Phone,
		// Role: ,
	}
}

// DomainCartToResponse
func DomainCartToResponse(cart *entities.Cart) Cart {
	return Cart{
		Id:       cart.ID.Hex(),
		UserId:   cart.UserID.Hex(),
		Products: DomainCartProductsToResponse(cart.Products),
	}
}

// DomainCartProductsToResponse
func DomainCartProductsToResponse(cartProducts []entities.CartProduct) CartProductList {
	products := make(CartProductList, 0, len(cartProducts))

	for _, product := range cartProducts {
		products = append(products, CartProduct{
			ProductId: product.ProductID.Hex(),
			Amount:    int(product.Amount),
		})
	}

	return products
}
