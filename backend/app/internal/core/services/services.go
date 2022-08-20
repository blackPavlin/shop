package services

import (
	"github.com/blackPavlin/shop/app/internal/core/services/address"
	"github.com/blackPavlin/shop/app/internal/core/services/cart"
	"github.com/blackPavlin/shop/app/internal/core/services/product"
	"github.com/blackPavlin/shop/app/internal/core/services/transaction"
	"github.com/blackPavlin/shop/app/internal/core/services/user"
)

// Services
type Services struct {
	UserService        user.UserService
	CartService        cart.CartService
	AddressService     address.AddressService
	ProductService     product.ProductService
	TransactionService transaction.TransactionService
}

// NewServices
func NewServices(
	userRepository user.UserRepository,
	cartRepository cart.CartRepository,
	addressRepository address.AddressRepository,
	productRepository product.ProductRepository,
	transactionRepository transaction.TransactionRepository,
) Services {
	return Services{
		UserService:        user.NewUserService(userRepository),
		CartService:        cart.NewCartService(cartRepository),
		AddressService:     address.NewAddressService(addressRepository),
		ProductService:     product.NewProductService(productRepository),
		TransactionService: transaction.NewTransactionService(transactionRepository),
	}
}
