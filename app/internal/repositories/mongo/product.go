package mongo

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/blackPavlin/shop/app/internal/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// ProductRepository
type ProductRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

// NewProductRepository
func NewProductRepository(db *mongo.Database, logger *zap.Logger) ProductRepository {
	collection := db.Collection("products")

	return ProductRepository{
		collection: collection,
		logger:     logger,
	}
}

// Create
func (p ProductRepository) Create(ctx context.Context, product *entities.Product) (entities.ProductID, error) {
	res, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		p.logger.Error("failed create product", zap.Error(err))
		return entities.ProductID{}, errs.ErrInternal
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		p.logger.Error("failid typecast product id", zap.Error(errs.ErrFailidTypecastObjectID))
		return entities.ProductID{}, errs.ErrInternal
	}
	return entities.ProductID(oid), nil
}

// FindByID
func (p ProductRepository) FindByID(ctx context.Context, id entities.ProductID) (*entities.Product, error) {
	product := &entities.Product{}

	if err := p.collection.FindOne(ctx, bson.M{"_id": primitive.ObjectID(id)}).Decode(product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrProductNotFound
		}
		p.logger.Error("failed find product by id", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return product, nil
}
