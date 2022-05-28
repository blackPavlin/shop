package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Phone    string             `json:"phone" bson:"phone"`
	Address  string             `json:"address" bson:"address"`
	Email    string             `json:"email" bson:"email"`
	Role     string             `json:"role" bson:"role"`
	Password string             `json:"-" bson:"password"`
}
