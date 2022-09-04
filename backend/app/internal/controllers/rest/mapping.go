package rest

import "github.com/blackPavlin/shop/app/internal/entities"

// DomainUserToResponse
func DomainUserToResponse(user *entities.User) User {
	return User{
		Id:    user.ID.Hex(),
		Email: user.Email,
		Name:  user.Name,
		Phone: user.Phone,
		Role:  UserRole(user.Role),
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

// DomainCategoryToResponse
func DomainCategoryToResponse(category *entities.Category) Category {
	return Category{
		Id:   category.ID.Hex(),
		Name: category.Name,
	}
}

// DomainCategoriesToResponse
func DomainCategoriesToResponse(categories []*entities.Category) CategoryList {
	result := make(CategoryList, 0, len(categories))

	for _, category := range categories {
		result = append(result, DomainCategoryToResponse(category))
	}

	return result
}

// DomainAddressToResponse
func DomainAddressToResponse(address *entities.Address) Address {
	return Address{
		Id:       address.ID.Hex(),
		City:     address.City,
		Country:  address.Country,
		Flat:     address.Flat,
		House:    address.House,
		Letter:   address.Letter,
		Postcode: address.Postcode,
		Street:   address.Street,
	}
}

// DomainAddressesToResponse
func DomainAddressesToResponse(addresses []*entities.Address) AddressList {
	result := make(AddressList, 0, len(addresses))

	for _, address := range addresses {
		result = append(result, DomainAddressToResponse(address))
	}

	return result
}
