package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartID primitive.ObjectID

type Cart struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userId" bson:"user_id"`
	Products []CartProduct      `json:"products" bson:"products"`
}

type CartProduct struct {
	ProductID primitive.ObjectID `json:"productId" bson:"product_id"`
	Amount    int64              `json:"amount" bson:"amount"`
}
