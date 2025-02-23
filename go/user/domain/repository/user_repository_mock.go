package repository

import (
	"sample/user/domain/entity"
	"sample/user/domain/model"
	"sample/user/domain/valueobject"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) List() ([]entity.User, error) {
	return nil, nil
}

func (m *MockUserRepository) Get(id valueobject.UserID) (*entity.User, error) {
	return nil, nil
}

func (m *MockUserRepository) Create(u *model.User) (*entity.User, error) {
	args := m.Called(u)

	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(u *model.User) error {
	return nil
}

func (m *MockUserRepository) Delete(id valueobject.UserID) error {
	return nil
}
