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

// CategoryRepository
type CategoryRepository struct {
	collection *mongo.Collection
}

// NewCategoryRepository
func NewCategoryRepository(db *mongo.Database) CategoryRepository {
	collection := db.Collection("categories")

	return CategoryRepository{
		collection: collection,
	}
}

// Create
func (c CategoryRepository) Create(ctx context.Context, category *entities.Category) (entities.CategoryID, error) {
	res, err := c.collection.InsertOne(ctx, category)
	if err != nil {
		return entities.CategoryID{}, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return entities.CategoryID(oid), nil
	}

	return entities.CategoryID{}, ErrFailidTypecastObjectID
}

// Find
func (c CategoryRepository) Find(ctx context.Context) ([]*entities.Category, error) {
	cursor, err := c.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	categories := []*entities.Category{}

	if err := cursor.All(ctx, categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// FindByID
func (c CategoryRepository) FindByID(ctx context.Context, id entities.CategoryID) (*entities.Category, error) {
	category := &entities.Category{}

	if err := c.collection.FindOne(ctx, bson.M{"_id": primitive.ObjectID(id)}).Decode(category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrCategoryNotFound
		}
		return nil, err
	}

	return category, nil
}

// FindByName
func (c CategoryRepository) FindByName(ctx context.Context, name string) (*entities.Category, error) {
	category := &entities.Category{}

	if err := c.collection.FindOne(ctx, bson.M{"name": name}).Decode(category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrCategoryNotFound
		}
		return nil, err
	}

	return category, nil
}

// DeleteByID
func (c CategoryRepository) DeleteByID(ctx context.Context, id entities.CategoryID) error {
	if _, err := c.collection.DeleteOne(ctx, bson.M{"_id": primitive.ObjectID(id)}); err != nil {
		return err
	}

	return nil
}

// Update
func (c CategoryRepository) Update(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	updates := bson.M{
		"$set": bson.M{
			"name": category.Name,
		},
	}

	if _, err := c.collection.UpdateOne(ctx, bson.M{"_id": primitive.ObjectID(category.ID)}, updates); err != nil {
		return nil, err
	}

	return c.FindByID(ctx, entities.CategoryID(category.ID))
}
