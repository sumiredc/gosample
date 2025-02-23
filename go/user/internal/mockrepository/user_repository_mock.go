package mockrepository

import (
	"sample/user/domain/entity"
	"sample/user/domain/model"
	"sample/user/domain/valueobject"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) List() ([]*entity.User, error) {
	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).([]*entity.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockUserRepository) Get(id valueobject.UserID) (*entity.User, error) {
	args := m.Called(id)

	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockUserRepository) Create(u *model.User) (*entity.User, error) {
	args := m.Called(u)

	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(u *model.User) error {
	args := m.Called(u)

	return args.Error(0)
}

func (m *MockUserRepository) Delete(id valueobject.UserID) error {
	args := m.Called(id)

	return args.Error(0)
}
