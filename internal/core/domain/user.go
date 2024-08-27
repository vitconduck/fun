package domain

import "time"

type UserRole string

// User is an entity that represents a user
type User struct {
	Id        uint64
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
