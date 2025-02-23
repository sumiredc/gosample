package dto

import (
	"sample/user/domain/entity"
	"sample/user/domain/valueobject"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	userID := "01JMPR93HT4GN3GDBMFP91C8J9"

	t.Run("should initialize GetUserInput with expected values", func(t *testing.T) {
		i := NewGetUserInput(userID)

		if i.UserID() != userID {
			t.Errorf("UserID missmatch: expected %q, but got %q", userID, i.UserID())
		}
	})
}

func TestGetUserOutput(t *testing.T) {
	userID := valueobject.NewUserID()
	name := "test name"
	email := "test@test.xxx"
	user := entity.NewUser(userID, name, email)

	t.Run("should initialize GetUserOutput with expected values", func(t *testing.T) {
		o := NewGetUserOutput(user)
		u := o.User()

		if u.ID.String() != userID.String() {
			t.Errorf("UserID missmatch: expected %q, but got %q", userID.String(), u.ID.String())
		}
	})
}
