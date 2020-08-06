package models

import "context"

// User model
type User struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

// UserRepository defines the set of User related methods available
type UserRepository interface {
	GetAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, ID string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(ID string) (string, error)
}
