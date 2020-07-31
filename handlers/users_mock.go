package handlers

import (
	"github.com/s1moe2/gosrv/models"
)

type userRepoMock struct {
	mockUsers       map[string]*models.User
	getAllImpl      func() ([]*models.User, error)
	findByIDImpl    func(ID string) (*models.User, error)
	findByEmailImpl func(email string) (*models.User, error)
	createImpl      func(user *models.User) (*models.User, error)
	updateImpl      func(user *models.User) (*models.User, error)
	deleteImpl      func(ID string) error
}

func newUserRepoMockDefault() *userRepoMock {
	return &userRepoMock{
		mockUsers: map[string]*models.User{
			"1": &models.User{
				ID:    "1",
				Name:  "user1",
				Email: "user1@eml.com",
			},
			"2": &models.User{
				ID:    "2",
				Name:  "user2",
				Email: "user2@eml.com",
			},
		},
	}
}

func newUserRepoMock(data []*models.User) *userRepoMock {
	users := map[string]*models.User{}
	for _, u := range data {
		users[u.ID] = u
	}
	return &userRepoMock{
		mockUsers: users,
	}
}

func (r *userRepoMock) GetAll() ([]*models.User, error) {
	return r.getAllImpl()
}

func (r *userRepoMock) FindByID(id string) (*models.User, error) {
	return r.findByIDImpl(id)
}

func (r *userRepoMock) FindByEmail(email string) (*models.User, error) {
	return r.findByEmailImpl(email)
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
