package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

// UserID
type UserID primitive.ObjectID

// UserRole
type UserRole string

// User
type User struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Phone    string             `json:"phone" bson:"phone"`
	Email    string             `json:"email" bson:"email"`
	Role     UserRole           `json:"role" bson:"role"`
	Password string             `json:"-" bson:"password"`
}
