package mongodb

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository struct {
	collection *mongo.Collection
}

func NewCartRepository(db *mongo.Database) *CartRepository {
	collection := db.Collection("carts")

	return &CartRepository{
		collection: collection,
	}
}

func (c *CartRepository) Create(ctx context.Context, cart *entities.Cart) error {
	if _, err := c.collection.InsertOne(ctx, cart); err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) FindByUserID(ctx context.Context, userID string) (*entities.Cart, error) {
	cart := &entities.Cart{}

	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	if err := c.collection.FindOne(ctx, bson.M{"user_id": uid}).Decode(cart); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return cart, nil
}

func (c *CartRepository) AddProduct(ctx context.Context, userID string, product *entities.CartProduct) error {
	if _, err := c.collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$push": bson.M{"products": product}}); err != nil {
		return err
	}

	return nil
}
