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

// CartRepository
type CartRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

// NewCartRepository
func NewCartRepository(db *mongo.Database, logger *zap.Logger) CartRepository {
	collection := db.Collection("carts")

	return CartRepository{
		collection: collection, 
		logger:     logger,
	}
}

// Create
func (c CartRepository) Create(ctx context.Context, cart *entities.Cart) (entities.CartID, error) {
	res, err := c.collection.InsertOne(ctx, cart)
	if err != nil {
		c.logger.Error("failed create cart", zap.Error(err))
		return entities.CartID{}, errs.ErrInternal
	}
	oid, ok := res.InsertedID.(primitive.ObjectID);
	if !ok {
		c.logger.Error("failid typecast cart id", zap.Error(errs.ErrFailidTypecastObjectID))
		return entities.CartID{}, errs.ErrInternal
		
	}
	return entities.CartID(oid), nil
}

// FindByUserID
func (c CartRepository) FindByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	cart := &entities.Cart{}
	if err := c.collection.FindOne(ctx, bson.M{"user_id": primitive.ObjectID(userID)}).Decode(cart); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrCartNotFound
		}
		c.logger.Error("failid find cart by user id", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return cart, nil
}

// DeleteProducts
func (c CartRepository) DeleteProducts(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	updates := bson.M{
		"$set": bson.M{
			"products": []entities.CartProduct{},
		},
	}
	if _, err := c.collection.UpdateOne(ctx, bson.M{"user_id": primitive.ObjectID(userID)}, updates); err != nil {
		c.logger.Error("failid delete products from cart", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return c.FindByUserID(ctx, userID)
}

// DeleteProductByID
func (c CartRepository) DeleteProductByID(ctx context.Context, userID entities.UserID, id entities.ProductID) (*entities.Cart, error) {
	updates := bson.M{
		"$pull": bson.M{
			"product": bson.M{
				"product_id": id,
			},
		},
	}
	if _, err := c.collection.UpdateOne(ctx, bson.M{"user_id": primitive.ObjectID(userID)}, updates); err != nil {
		c.logger.Error("failid delete product from cart by id", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return c.FindByUserID(ctx, userID)
}

// AppendProduct
func (c CartRepository) AppendProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error) {
	updates := bson.M{
		"$push": bson.M{
			"product": product,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, bson.M{"user_id": primitive.ObjectID(userID)}, updates); err != nil {
		c.logger.Error("failid append product to cart", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return c.FindByUserID(ctx, userID)
}

// UpdateProduct
func (c CartRepository) UpdateProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error) {
	filter := bson.M{
		"user_id": primitive.ObjectID(userID),
		"products": bson.M{
			"product_id": primitive.ObjectID(product.ProductID),
		},
	}
	updates := bson.M{
		"$set": bson.M{
			"amount": product.Amount,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, filter, updates); err != nil {
		c.logger.Error("failid update product from cart", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return c.FindByUserID(ctx, userID)
}
