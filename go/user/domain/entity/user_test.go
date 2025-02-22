package entity

import (
	"sample/user/domain/valueobject"
	"testing"
)

func TestNewUser(t *testing.T) {
	userID := valueobject.NewUserID()
	name := "John Doe"
	email := "sample@test.xxx"

	t.Run("Make to user entity instance", func(t *testing.T) {
		user := NewUser(userID, name, email)

		if userID.Value().String() != user.ID.Value().String() {
			t.Errorf("don't match `userID` [expected: %v, actual: %v]", userID.Value(), user.ID.Value())
		}

		if name != user.Name {
			t.Errorf("don't match `name` [expected: %v, actual: %v]", name, user.Name)
		}

		if email != user.Email {
			t.Errorf("don't match `email` [expected: %v, actual: %v]", email, user.Email)
		}
	})
}
