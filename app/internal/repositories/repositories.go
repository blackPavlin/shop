package repositories

import (
	mongor "github.com/blackPavlin/shop/app/internal/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Repositories
type Repositories struct {
	UserRepository        mongor.UserRepository
	CartRepository        mongor.CartRepository
	AddressRepository     mongor.AddressRepository
	ProductRepository     mongor.ProductRepository
	TransactionRepository mongor.TransactionRepository
	CategoryRepository    mongor.CategoryRepository
	OrderRepository       mongor.OrderRepository
}

// NewRepositories
func NewRepositories(mongoDB *mongo.Database, logger *zap.Logger) Repositories {
	return Repositories{
		// Mongo
		UserRepository:        mongor.NewUserRepository(mongoDB, logger),
		CartRepository:        mongor.NewCartRepository(mongoDB, logger),
		AddressRepository:     mongor.NewAddressRepository(mongoDB, logger),
		ProductRepository:     mongor.NewProductRepository(mongoDB, logger),
		TransactionRepository: mongor.NewTransactionRepository(mongoDB, logger),
		CategoryRepository:    mongor.NewCategoryRepository(mongoDB, logger),
		OrderRepository:       mongor.NewOrderRepository(mongoDB, logger),
	}
}
