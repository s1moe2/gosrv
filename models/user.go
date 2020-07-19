package models

// User model
type User struct {
	ID    string `json:"id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

// UserRepository defines the set of User related methods available
type UserRepository interface {
	GetAll() ([]*User, error)
	FindByID(ID string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(ID string) error
}
