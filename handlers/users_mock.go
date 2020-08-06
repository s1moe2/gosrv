package handlers

import (
	"github.com/s1moe2/gosrv/models"
)

type userRepoMock struct {
	getAllImpl      func() ([]*models.User, error)
	findByIDImpl    func(ID string) (*models.User, error)
	findByEmailImpl func(email string) (*models.User, error)
	createImpl      func(user *models.User) (*models.User, error)
	updateImpl      func(user *models.User) (*models.User, error)
	deleteImpl      func(ID string) (string, error)
}

func newUserRepoMockDefault() *userRepoMock {
	return &userRepoMock{}
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

func (r *userRepoMock) Delete(id string) (string, error) {
	return r.deleteImpl(id)
}
