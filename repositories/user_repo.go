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

// GetAll ..
func (r *UserRepo) GetAll() []*models.User {
	var users []*models.User
	return users
}

// FindByID ..
func (r *UserRepo) FindByID(ID string) (*models.User, error) {
	return &models.User{}, nil
}

// Create ..
func (r *UserRepo) Create(user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

// Update ..
func (r *UserRepo) Update(user *models.User) (*models.User, error) {
	return &models.User{}, nil
}
// Save ..
func (r *UserRepo) Delete(ID string) error {
	return nil
}