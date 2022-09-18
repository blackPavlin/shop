package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// CartID
type CartID primitive.ObjectID

// Cart
type Cart struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userId" bson:"user_id"`
	Products CartProducts       `json:"products" bson:"products"`
}

// CartProduct
type CartProduct struct {
	ProductID primitive.ObjectID `json:"productId" bson:"product_id"`
	Amount    int64              `json:"amount" bson:"amount"`
}

// CartProducts
type CartProducts = []CartProduct
