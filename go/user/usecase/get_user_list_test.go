package usecase_test

import (
	"errors"
	"fmt"
	"sample/user/domain/entity"
	"sample/user/domain/valueobject"
	"sample/user/internal/mockrepository"
	"sample/user/usecase"
	"sample/user/usecase/dto"
	"testing"
)

func TestGetUserListUseCase(t *testing.T) {
	t.Run("should get users successfully", func(t *testing.T) {
		expected := 5
		var users []*entity.User

		for i := 0; i < expected; i++ {
			user := entity.NewUser(
				valueobject.NewUserID(),
				fmt.Sprintf("test name%d", i),
				fmt.Sprintf("test%d@test.xxx", i),
			)
			users = append(users, user)
		}

		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("List").Return(users, nil)

		i := dto.NewGetUserListInput()
		u := usecase.NewGetUserListUseCase(mRepo)
		o, err := u.Execute(i)

		if err != nil {
			t.Errorf("failed to get user list: %v", err)
		}

		if len(o.Users()) != expected {
			t.Errorf("failed to return users length: expected %q, but got %q", expected, len(o.Users()))
		}

		mRepo.AssertNumberOfCalls(t, "List", 1)
	})

	t.Run("should return to error", func(t *testing.T) {
		mRepo := new(mockrepository.MockUserRepository)
		mRepo.On("List").Return(nil, errors.New("failed to get user list"))

		i := dto.NewGetUserListInput()
		u := usecase.NewGetUserListUseCase(mRepo)
		_, err := u.Execute(i)

		if err == nil {
			t.Errorf("failed to return not error")
		}

		mRepo.AssertNumberOfCalls(t, "List", 1)
	})
}
