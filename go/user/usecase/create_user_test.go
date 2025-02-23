package usecase_test

import (
	"errors"
	"sample/user/domain/entity"
	"sample/user/domain/valueobject"
	"sample/user/internal/mockrepository"
	"sample/user/usecase"
	"sample/user/usecase/dto"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCreateUserUseCase(t *testing.T) {
	t.Run("should create a user successfully", func(t *testing.T) {
		userID := valueobject.NewUserID()
		name := "test name"
		email := "test@test.xxx"
		eUser := entity.NewUser(userID, name, email)

		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("Create", mock.Anything).Return(eUser, nil)

		i := dto.NewCreateUserInput(name, email)
		u := usecase.NewCreateUserUseCase(mRepo)
		o, err := u.Execute(i)

		if err != nil {
			t.Errorf("failed to create user: %v", err)
		}

		user := o.User()

		if user.Name != name {
			t.Errorf("user name in Output missmatch: expected %q, but got %q", name, user.Name)
		}

		if user.Email != email {
			t.Errorf("user email in Output missmatch: expected %q, but got %q", email, user.Email)
		}

		mRepo.AssertNumberOfCalls(t, "Create", 1)
	})

	t.Run("should return to error", func(t *testing.T) {
		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("Create", mock.Anything).Return(nil, errors.New("Failed to create user"))

		i := dto.NewCreateUserInput("name", "email")
		u := usecase.NewCreateUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		mRepo.AssertNumberOfCalls(t, "Create", 1)
	})
}
