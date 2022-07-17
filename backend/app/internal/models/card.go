package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"-" bson:"user_id"`
	Products []struct {
		ProductID primitive.ObjectID `json:"-" bson:"product_id"`
		Count     int32              `json:"-" bson:"count"`
	} `json:"-" bson:"products"`
}
