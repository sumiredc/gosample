package usecase_test

import (
	"errors"
	"sample/user/domain/valueobject"
	"sample/user/internal/mockrepository"
	"sample/user/usecase"
	"sample/user/usecase/dto"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUpdateUserUseCase(t *testing.T) {
	t.Run("should update a user successfully", func(t *testing.T) {
		userID := valueobject.NewUserID()
		name := "test name"
		email := "test@test.xxx"

		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("Update", mock.Anything).Return(nil)

		i := dto.NewUpdateUserInput(userID.String(), name, email)
		u := usecase.NewUpdateUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err != nil {
			t.Errorf("failed to update user: %v", err)
		}

		mRepo.AssertNumberOfCalls(t, "Update", 1)
	})

	t.Run("should return to error", func(t *testing.T) {
		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("Update", mock.Anything).Return(errors.New("failed to update user"))

		userID := valueobject.NewUserID()
		name := "test name"
		email := "test@test.xxx"

		i := dto.NewUpdateUserInput(userID.String(), name, email)
		u := usecase.NewUpdateUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		mRepo.AssertNumberOfCalls(t, "Update", 1)
	})
}
