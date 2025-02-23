package usecase_test

import (
	"errors"
	"sample/user/domain/valueobject"
	"sample/user/internal/mockrepository"
	"sample/user/usecase"
	"sample/user/usecase/dto"
	"testing"
)

func TestDeleteUserUseCase(t *testing.T) {
	t.Run("should delete a user successfully", func(t *testing.T) {
		userID := valueobject.NewUserID()

		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("Delete", userID).Return(nil)

		i := dto.NewDeleteUserInput(userID.String())
		u := usecase.NewDeleteUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err != nil {
			t.Errorf("failed to delete user: %v", err)
		}

		mRepo.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("should return to error", func(t *testing.T) {
		userID := valueobject.NewUserID()

		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("Delete", userID).Return(errors.New("failed to delete user"))

		i := dto.NewDeleteUserInput(userID.String())
		u := usecase.NewDeleteUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		mRepo.AssertNumberOfCalls(t, "Delete", 1)
	})
}
