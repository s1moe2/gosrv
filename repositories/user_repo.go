package repositories

import (
	"database/sql"

	"github.com/s1moe2/gosrv/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo returns a configured UserRepo object
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// GetAll fetches all users, returns an empty slice if no user exists
func (r *UserRepo) GetAll() ([]*models.User, error) {
	var users []*models.User
	return users, nil
}

// FindByID finds a user by ID, returns nil if not found
func (r *UserRepo) FindByID(ID string) (*models.User, error) {
	return nil, nil
}

// Create creates a new user, returning the full model
func (r *UserRepo) Create(user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

// Update updates new user, returning the updated model
func (r *UserRepo) Update(user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

// Delete deletes a user, only returns error if action fails
func (r *UserRepo) Delete(ID string) error {
	return nil
}
