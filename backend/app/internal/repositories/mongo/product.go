package mongo

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository
type ProductRepository struct {
	collection *mongo.Collection
}

// NewProductRepository
func NewProductRepository(db *mongo.Database) ProductRepository {
	collection := db.Collection("products")

	return ProductRepository{
		collection: collection,
	}
}

// Create
func (p ProductRepository) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	res, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		product.ID = entities.ProductID(oid)
		return product, nil
	}

	return nil, ErrFailidTypecastObjectID
}
