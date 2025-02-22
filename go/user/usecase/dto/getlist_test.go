package dto

import (
	"fmt"
	"sample/user/domain/entity"
	"sample/user/domain/valueobject"
	"testing"
)

func TestGetUserListInput(t *testing.T) {
	t.Run("should initialize GetUserListInput", func(t *testing.T) {
		i := NewGetUserListInput()

		if *i != (GetUserListInput{}) {
			t.Errorf("failed to instance of GetUserListInput: instance %v", *i)
		}
	})
}

func TestGetUserListOutput(t *testing.T) {
	expected := 5
	var users []entity.User

	for i := 0; i < expected; i++ {
		userID := valueobject.NewUserID()
		name := fmt.Sprintf("user%d", i)
		email := fmt.Sprintf("user%d@test.xxx", i)
		users = append(users, entity.NewUser(userID, name, email))
	}

	t.Run("should initialize GetUserListOutput with expected values", func(t *testing.T) {
		o := NewGetUserListOutput(users)

		if len(o.Users()) != expected {
			t.Errorf("Users length missmatch: expected %q, but got %q", expected, len(o.Users()))
		}
	})

}
