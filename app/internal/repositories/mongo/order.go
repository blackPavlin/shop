package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// OrderRepository
type OrderRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

// NewOrderRepository
func NewOrderRepository(db *mongo.Database, logger *zap.Logger) OrderRepository {
	collection := db.Collection("orders")

	return OrderRepository{
		collection: collection,
		logger:     logger,
	}
}
