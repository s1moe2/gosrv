package handlers

import (
	"github.com/s1moe2/gosrv/models"
)

type userRepoMock struct {
	mockUsers    []*models.User
	getAllImpl   func() ([]*models.User, error)
	findByIDImpl func(ID string) (*models.User, error)
	createImpl   func(user *models.User) (*models.User, error)
	updateImpl   func(user *models.User) (*models.User, error)
	deleteImpl   func(ID string) error
}

func newUserRepoMockDefault() *userRepoMock {
	return &userRepoMock{
		mockUsers: []*models.User{
			&models.User{
				ID:    "1",
				Name:  "user1",
				Email: "user1@eml.com",
			},
			&models.User{
				ID:    "2",
				Name:  "user2",
				Email: "user2@eml.com",
			},
		},
	}
}

func newUserRepoMock(data []*models.User) *userRepoMock {
	return &userRepoMock{
		mockUsers: data,
	}
}

func (r *userRepoMock) GetAll() ([]*models.User, error) {
	return r.getAllImpl()
}

func (r *userRepoMock) FindByID(id string) (*models.User, error) {
	return r.findByIDImpl(id)
}

func (r *userRepoMock) Create(user *models.User) (*models.User, error) {
	return r.createImpl(user)
}

func (r *userRepoMock) Update(user *models.User) (*models.User, error) {
	return r.updateImpl(user)
}

func (r *userRepoMock) Delete(id string) error {
	return r.deleteImpl(id)
}
