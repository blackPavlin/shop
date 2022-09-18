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

// CategoryRepository
type CategoryRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

// NewCategoryRepository
func NewCategoryRepository(db *mongo.Database, logger *zap.Logger) CategoryRepository {
	collection := db.Collection("categories")

	return CategoryRepository{
		collection: collection,
		logger:     logger,
	}
}

// Create
func (c CategoryRepository) Create(ctx context.Context, category *entities.Category) (entities.CategoryID, error) {
	res, err := c.collection.InsertOne(ctx, category)
	if err != nil {
		c.logger.Error("cailed create category", zap.Error(err))
		return entities.CategoryID{}, errs.ErrInternal
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.logger.Error("failid typecast category id", zap.Error(errs.ErrFailidTypecastObjectID))
		return entities.CategoryID{}, errs.ErrInternal
		
	}
	return entities.CategoryID(oid), nil
}

// Find
func (c CategoryRepository) Find(ctx context.Context) (entities.Categories, error) {
	cursor, err := c.collection.Find(ctx, bson.D{})
	if err != nil {
		c.logger.Error("failid find categories", zap.Error(err))
		return nil, errs.ErrInternal
	}
	defer cursor.Close(ctx)
	categories := &entities.Categories{}
	if err := cursor.All(ctx, categories); err != nil {
		c.logger.Error("failid find categories", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return *categories, nil
}

// FindByID
func (c CategoryRepository) FindByID(ctx context.Context, id entities.CategoryID) (*entities.Category, error) {
	category := &entities.Category{}
	if err := c.collection.FindOne(ctx, bson.M{"_id": primitive.ObjectID(id)}).Decode(category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrCategoryNotFound
		}
		c.logger.Error("failid find category by id", zap.Error(err))
		return nil, errs.ErrInternal
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
		c.logger.Error("failid find category by name", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return category, nil
}

// DeleteByID
func (c CategoryRepository) DeleteByID(ctx context.Context, id entities.CategoryID) error {
	if _, err := c.collection.DeleteOne(ctx, bson.M{"_id": primitive.ObjectID(id)}); err != nil {
		c.logger.Error("failid delete category by id", zap.Error(err))
		return  errs.ErrInternal
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
		c.logger.Error("failid update category", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return c.FindByID(ctx, entities.CategoryID(category.ID))
}
