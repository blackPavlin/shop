package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrderID
type OrderID primitive.ObjectID

// Order
type Order struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userId" bson:"user_id"`
	Address  Address            `json:"address" bson:"address"`
	Products []OrderProduct     `json:"products" bson:"products"`
}

// OrderProduct
type OrderProduct struct {
	Product
	Amount int64 `json:"amount" bson:"count"`
}
