// Package user contains user oriented logic.
package user

import "time"

//go:generate go run github.com/dmarkham/enumer@v1.5.9 -sql -linecomment -type Role -trimprefix Role -output role_string.go

const (
	RoleUser  Role = iota // USER
	RoleAdmin             // ADMIN
)

// ID represents an id for user.
type ID int64

// IDs defines a slice of ID.
type IDs []ID

// Role represents an user`s roles.
type Role uint8

// Roles defines a slice of Role.
type Roles []Role

// User represents the user.
type User struct {
	ID   ID
	Role Role

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Props represents user editable fields.
type Props struct {
	Email    string
	Name     string
	Phone    string
	Password string
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
