package dto

import (
	"sample/user/domain/entity"
	"sample/user/domain/valueobject"
	"testing"
)

func TestCreateUserInput(t *testing.T) {
	name := "test name"
	email := "test@test.xxx"

	t.Run("should initialize CreateUserInput with expected values", func(t *testing.T) {
		i := NewCreateUserInput(name, email)

		if i.Name() != name {
			t.Errorf("Name mismatch: expected %q, but got %q", name, i.Name())
		}

		if i.Email() != email {
			t.Errorf("Email mismatch: expected %q, but got %q", email, i.Email())
		}
	})
}

func TestCreateUserOutput(t *testing.T) {
	userID := valueobject.NewUserID()
	name := "test name"
	email := "test@test.xxx"
	user := entity.NewUser(userID, name, email)

	t.Run("should initialize CreateUserOutput with expected values", func(t *testing.T) {
		o := NewCreateUserOutput(user)
		u := o.User()

		if u.ID.String() != userID.String() {
			t.Errorf("User mismatch: expected %q, but got %q", userID.String(), u.ID.String())
		}
	})
}
