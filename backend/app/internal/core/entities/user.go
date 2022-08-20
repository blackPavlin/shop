package entities

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type UserID [12]byte

type UserRole string

type User struct {
	ID       UserID   `json:"-" bson:"_id,omitempty"`
	Name     string   `json:"name" bson:"name"`
	Phone    string   `json:"phone" bson:"phone"`
	Email    string   `json:"email" bson:"email"`
	Role     UserRole `json:"role" bson:"role"`
	Password string   `json:"-" bson:"password"`
}
