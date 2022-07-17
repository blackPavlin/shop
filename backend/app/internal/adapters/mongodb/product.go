package mongodb

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	collection := db.Collection("products")

	return &ProductRepository{
		collection: collection,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *entities.Product) error {
	if _, err := p.collection.InsertOne(ctx, product); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) FindByID(ctx context.Context, id string) (*entities.Product, error) {
	product := &entities.Product{}

	if err := p.collection.FindOne(ctx, bson.M{"_id": id}).Decode(product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return product, nil
}
