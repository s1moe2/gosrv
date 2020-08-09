package repositories

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"

	"github.com/s1moe2/gosrv/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *sqlx.DB
}

// NewUserRepo returns a configured UserRepo object
func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// GetAll fetches all users, returns an empty slice if no user exists
func (r *UserRepo) GetAll(ctx context.Context) ([]*models.User, error) {
	users := []*models.User{}
	err := r.db.SelectContext(ctx, &users, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID finds a user by ID, returns nil if not found
func (r *UserRepo) FindByID(ctx context.Context, ID string) (*models.User, error) {
	user := &models.User{}
	err := r.db.GetContext(ctx, user, "SELECT id, name, email FROM users WHERE id = $1", ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// FindByEmail finds a user by email, returns nil if not found
func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	stmt := "SELECT id, name, email FROM users WHERE email = $1"
	err := r.db.GetContext(ctx, user, stmt, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// Create creates a new user, returning the full model
func (r *UserRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	stmt := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	row := r.db.QueryRowContext(ctx, stmt, user.Name, user.Name)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates new user, returning the updated model
func (r *UserRepo) Update(user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

// Delete deletes a user, only returns error if action fails
func (r *UserRepo) Delete(ID string) (string, error) {
	return ID, nil
}
