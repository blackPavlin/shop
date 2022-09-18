package services

import (
	"github.com/blackPavlin/shop/app/internal/services/address"
	"github.com/blackPavlin/shop/app/internal/services/cart"
	"github.com/blackPavlin/shop/app/internal/services/category"
	"github.com/blackPavlin/shop/app/internal/services/order"
	"github.com/blackPavlin/shop/app/internal/services/product"
	"github.com/blackPavlin/shop/app/internal/services/transaction"
	"github.com/blackPavlin/shop/app/internal/services/user"
)

// Services
type Services struct {
	UserService        user.UserService
	CartService        cart.CartService
	AddressService     address.AddressService
	ProductService     product.ProductService
	TransactionService transaction.TransactionService
	CategoryService    category.CategoryService
	OrderService       order.OrderService
}

// NewServices
func NewServices(
	userRepository user.UserRepository,
	cartRepository cart.CartRepository,
	addressRepository address.AddressRepository,
	productRepository product.ProductRepository,
	transactionRepository transaction.TransactionRepository,
	categoryRepository category.CategoryRepository,
	opderRepository order.OrderRepository,
) Services {
	return Services{
		UserService:        user.NewUserService(userRepository),
		CartService:        cart.NewCartService(cartRepository),
		AddressService:     address.NewAddressService(addressRepository),
		ProductService:     product.NewProductService(productRepository),
		TransactionService: transaction.NewTransactionService(transactionRepository),
		CategoryService:    category.NewCategoryService(categoryRepository),
		OrderService:       order.NewOrderService(opderRepository),
	}
}
