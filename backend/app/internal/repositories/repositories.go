package repositories

import (
	mongor "github.com/blackPavlin/shop/app/internal/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repositories
type Repositories struct {
	UserRepository        mongor.UserRepository
	CartRepository        mongor.CartRepository
	AddressRepository     mongor.AddressRepository
	ProductRepository     mongor.ProductRepository
	TransactionRepository mongor.TransactionRepository
}

// NewRepositories
func NewRepositories(mongoDB *mongo.Database) Repositories {
	return Repositories{
		// Mongo
		UserRepository:        mongor.NewUserRepository(mongoDB),
		CartRepository:        mongor.NewCartRepository(mongoDB),
		AddressRepository:     mongor.NewAddressRepository(mongoDB),
		ProductRepository:     mongor.NewProductRepository(mongoDB),
		TransactionRepository: mongor.NewTransactionRepository(mongoDB),
	}
}
