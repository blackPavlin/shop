package mongodb

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	collection := db.Collection("categories")

	return &CategoryRepository{
		collection: collection,
	}
}

func (c *CategoryRepository) Create(ctx context.Context, category *entities.Category) error {
	if _, err := c.collection.InsertOne(ctx, category); err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) Find(ctx context.Context) ([]*entities.Category, error) {
	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	categories := []*entities.Category{}

	for cursor.Next(ctx) {
		category := &entities.Category{}

		if err := cursor.Decode(category); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategoryRepository) FindByID(ctx context.Context, id string) (*entities.Category, error) {
	category := &entities.Category{}

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := c.collection.FindOne(ctx, bson.M{"_id": uid}).Decode(category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return category, nil
}

func (c *CategoryRepository) FindByName(ctx context.Context, name string) (*entities.Category, error) {
	category := &entities.Category{}

	if err := c.collection.FindOne(ctx, bson.M{"name": name}).Decode(category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return category, nil
}
