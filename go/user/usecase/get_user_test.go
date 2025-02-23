package usecase_test

import (
	"errors"
	"sample/user/domain/entity"
	"sample/user/domain/errkit"
	"sample/user/domain/repository/testdata"
	"sample/user/domain/valueobject"
	"sample/user/usecase"
	"sample/user/usecase/dto"
	"testing"
)

func TestGetUserUseCase(t *testing.T) {
	t.Run("should get a user successfully", func(t *testing.T) {
		userID := valueobject.NewUserID()
		name := "test name"
		email := "test@test.xxx"
		eUser := entity.NewUser(userID, name, email)

		mRepo := new(testdata.MockUserRepository)
		mRepo.On("Get", userID).Return(eUser, nil)

		i := dto.NewGetUserInput(userID.String())
		u := usecase.NewGetUserUseCase(mRepo)
		o, err := u.Execute(i)

		if err != nil {
			t.Errorf("failed to get a user: %v", err)
		}

		user := o.User()

		if user.ID.String() != userID.String() {
			t.Errorf("UserID missmatch: expected %q, but got %q", userID.String(), user.ID.String())
		}

		mRepo.AssertNumberOfCalls(t, "Get", 1)
	})

	t.Run("should return to failed to parse UserID error", func(t *testing.T) {
		userID := "NO-ULID"

		mRepo := new(testdata.MockUserRepository)

		i := dto.NewGetUserInput(userID)
		u := usecase.NewGetUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		if !errors.Is(err, errkit.ErrMissingParameter) {
			t.Errorf("err missmatch: %v", err)
		}

		mRepo.AssertNumberOfCalls(t, "Get", 0)
	})

	t.Run("should return to failed to get error", func(t *testing.T) {
		userID := valueobject.NewUserID()

		mRepo := new(testdata.MockUserRepository)
		mRepo.On("Get", userID).Return(nil, errors.New("failed to get a user"))

		i := dto.NewGetUserInput(userID.String())
		u := usecase.NewGetUserUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		mRepo.AssertNumberOfCalls(t, "Get", 1)
	})
}
