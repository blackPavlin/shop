package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"-" bson:"user_id"`
	Products []CartProduct      `json:"-" bson:"products"`
}

type CartProduct struct {
	ProductID primitive.ObjectID `json:"-" bson:"product_id"`
	Count     int64              `json:"-" bson:"count"`
}
