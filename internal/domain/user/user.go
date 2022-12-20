package user

import "time"

//go:generate enumer -sql -linecomment -type UserRole -trimprefix Role -output role_string.go user.go

const (
	RoleUser  UserRole = iota // USER
	RoleAdmin                 // ADMIN
)

// ID
type ID int64

// UserRole
type UserRole uint8

// User
type User struct {
	ID   ID
	Role UserRole

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
