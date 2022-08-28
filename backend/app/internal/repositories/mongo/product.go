package mongo

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/blackPavlin/shop/app/internal/errs"
	"go.mongodb.org/mongo-driver/bson"
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
func (p ProductRepository) Create(ctx context.Context, product *entities.Product) (entities.ProductID, error) {
	res, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return entities.ProductID{}, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return entities.ProductID(oid), nil
	}

	return entities.ProductID{}, ErrFailidTypecastObjectID
}

// FindByID
func (p ProductRepository) FindByID(ctx context.Context, id entities.ProductID) (*entities.Product, error) {
	product := &entities.Product{}

	if err := p.collection.FindOne(ctx, bson.M{"_id": primitive.ObjectID(id)}).Decode(product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrProductNotFound
		}
		return nil, err
	}

	return product, nil
}
