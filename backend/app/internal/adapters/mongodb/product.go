package mongodb

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (p *ProductRepository) FindByCategoryID(ctx context.Context, categoryID string, offset, limit int64) ([]*entities.Product, error) {
	uid, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, err
	}

	options := &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	}

	cursor, err := p.collection.Find(ctx, bson.M{"category_id": uid}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	products := []*entities.Product{}

	for cursor.Next(ctx) {
		product := &entities.Product{}

		if err := cursor.Decode(product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
