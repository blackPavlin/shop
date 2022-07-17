package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID primitive.ObjectID `json:"-" bson:"_id,omitempty"`
}
