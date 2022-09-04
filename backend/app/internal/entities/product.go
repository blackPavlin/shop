package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductID primitive.ObjectID

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CategoryID  primitive.ObjectID `json:"categoryId" bson:"category_id"`
	Amount      int64              `json:"amount" bson:"amount"`
	Price       float64            `json:"price" bson:"price"`
}
