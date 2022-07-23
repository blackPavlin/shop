package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CategoryID  primitive.ObjectID `json:"category_id" bson:"category_id"`
	Quantity    int64              `json:"quantity" bson:"quantity"`
}
