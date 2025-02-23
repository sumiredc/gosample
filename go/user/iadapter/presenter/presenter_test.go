package presenter_test

import (
	"fmt"
	"reflect"
	"sample/user/domain/entity"
	"sample/user/domain/valueobject"
	"sample/user/iadapter/presenter"
	"sample/user/iadapter/response"
	"testing"
)

func TestPresenter(t *testing.T) {
	p := presenter.New()

	t.Run("should return to UserList response successfully", func(t *testing.T) {
		var users []*entity.User

		for i := 0; i < 5; i++ {
			user := entity.NewUser(
				valueobject.NewUserID(),
				fmt.Sprintf("test name%d", i),
				fmt.Sprintf("test%d@test.xxx", i),
			)
			users = append(users, user)
		}

		res := p.UserList(users)

		expected := &response.UserListResponse{Users: users}

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("UserListResponse missmatch: expected %v, but got %v", expected, res)
		}
	})

	t.Run("should return to User response successfully", func(t *testing.T) {
		user := entity.NewUser(
			valueobject.NewUserID(),
			"test name",
			"test@test.xxx",
		)

		res := p.User(user)

		expected := &response.UserResponse{User: user}

		if !reflect.DeepEqual(res, expected) {
			t.Errorf("UserResponse missmatch: expected %v, but got %v", expected, res)
		}
	})
}
