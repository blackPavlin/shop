package user

import "time"

//go:generate enumer -sql -linecomment -type Role -trimprefix Role -output role_string.go user.go

const (
	RoleUser  Role = iota // USER
	RoleAdmin             // ADMIN
)

// ID
type ID int64

// IDs
type IDs []ID

// Role
type Role uint8

// User
type User struct {
	ID   ID
	Role Role

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

type Props struct {
	Email    string
	Name     string
	Phone    string
	Password string
}

// ToInt64
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
