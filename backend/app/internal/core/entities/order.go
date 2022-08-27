package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderID primitive.ObjectID

type Order struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userId" bson:"user_id"`
	Address  Address            `json:"address" bson:"address"`
	Products []OrderProduct     `json:"products" bson:"products"`
}

type OrderProduct struct {
	ProductID primitive.ObjectID `json:"productId" bson:"product_id"`
	Amount    int64              `json:"amount" bson:"count"`
}
