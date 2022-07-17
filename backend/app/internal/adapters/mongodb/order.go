package mongodb

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	collection := db.Collection("orders")

	return &OrderRepository{
		collection: collection,
	}
}

func (o *OrderRepository) Create(ctx context.Context, order *entities.Order) error {
	if _, err := o.collection.InsertOne(ctx, order); err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) FindByUserID(ctx context.Context, userID string) ([]*entities.Order, error) {
	cursor, err := o.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	orders := []*entities.Order{}

	for cursor.Next(ctx) {
		order := &entities.Order{}

		if err := cursor.Decode(order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
