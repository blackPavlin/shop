package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// AddressID
type AddressID primitive.ObjectID

// Address
type Address struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userId" bson:"user_id"`
	City     string             `json:"city" bson:"city"`
	Country  string             `json:"country" bson:"country"`
	Flat     *int               `json:"flat,omitempty" bson:"flat,omitempty"`
	House    int                `json:"house" bson:"house"`
	Letter   *string            `json:"letter,omitempty" bson:"letter,omitempty"`
	Postcode int                `json:"postcode" bson:"postcode"`
	Street   string             `json:"street" bson:"street"`
}
