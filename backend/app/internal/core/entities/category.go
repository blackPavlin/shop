package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryID primitive.ObjectID

type Category struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
