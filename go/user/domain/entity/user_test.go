package entity

import (
	"sample/user/domain/valueobject"
	"testing"
)

func TestUserEntity(t *testing.T) {
	userID := valueobject.NewUserID()
	name := "John Doe"
	email := "sample@test.xxx"

	t.Run("should initialize User with expected values", func(t *testing.T) {
		user := NewUser(userID, name, email)

		if user.ID.String() != userID.String() {
			t.Errorf("ID missmatch: expected %q, but got %q", userID.String(), user.ID.String())
		}

		if user.Name != name {
			t.Errorf("Name missmatch: expected %q, but got %q", name, user.Name)
		}

		if user.Email != email {
			t.Errorf("Email missmatch: expected %q, but got %q", email, user.Email)
		}
	})
}
