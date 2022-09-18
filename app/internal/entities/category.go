package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// CategoryID
type CategoryID primitive.ObjectID

// Category
type Category struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

// Сategories
type Categories = []*Category
