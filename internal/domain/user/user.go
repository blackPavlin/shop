// Package user contains user oriented logic.
package user

import "time"

//go:generate go run github.com/dmarkham/enumer@v1.5.9 -sql -linecomment -type Role -trimprefix Role -output role_string.go

const (
	RoleUser  Role = iota + 1 // USER
	RoleAdmin                 // ADMIN
)

// ID represents an id for user.
type ID int64

// IDs represents a slice of ID.
type IDs []ID

// Role represents an user`s roles.
type Role uint8

// Roles represents a slice of Role.
type Roles []Role

// User represents the user.
type User struct {
	ID        ID
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	Props     *Props
}

// Props represents user editable fields.
type Props struct {
	Email    string
	Name     string
	Phone    string
	Password string
}

// ToInt64 convert ID to int64.
func (id ID) ToInt64() int64 {
	return int64(id)
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.ToInt64())
	}
	return result
}
