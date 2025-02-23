package usecase_test

import (
	"errors"
	"sample/user/domain/errkit"
	"sample/user/domain/repository/testdata"
	"sample/user/domain/valueobject"
	"sample/user/usecase"
	"sample/user/usecase/dto"
	"testing"
)

func TestDeleteUserUseCase(t *testing.T) {
	t.Run("should delete a user successfully", func(t *testing.T) {
		userID := valueobject.NewUserID()

		mRepo := new(testdata.MockUserRepository)
		mRepo.On("Delete", userID).Return(nil)

		i := dto.NewDeleteUserInput(userID.String())
		u := usecase.NewDeleteUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err != nil {
			t.Errorf("failed to delete user: %v", err)
		}

		mRepo.AssertNumberOfCalls(t, "Delete", 1)
	})

	t.Run("should return to failed to parse UserID error", func(t *testing.T) {
		userID := "NO-ULID"

		mRepo := new(testdata.MockUserRepository)

		i := dto.NewDeleteUserInput(userID)
		u := usecase.NewDeleteUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		if !errors.Is(err, errkit.ErrMissingParameter) {
			t.Errorf("err missmatch: %v", err)
		}

		mRepo.AssertNumberOfCalls(t, "Delete", 0)
	})

	t.Run("should return to failed to delete error", func(t *testing.T) {
		userID := valueobject.NewUserID()

		mRepo := new(testdata.MockUserRepository)
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
