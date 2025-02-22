package dto

import (
	"testing"
)

func TestUpdateUserInput(t *testing.T) {
	userID := "01JMPR93HT4GN3GDBMFP91C8J9"
	name := "test name"
	email := "test@test.xxx"

	t.Run("should initialize UpdateUserInput with expected values", func(t *testing.T) {
		i := NewUpdateUserInput(userID, name, email)

		if i.UserID() != userID {
			t.Errorf("UserID missmatch: expected %q, but got %q", userID, i.UserID())
		}

		if i.Name() != name {
			t.Errorf("Name missmatch: expected %q, but got %q", name, i.Name())
		}

		if i.Email() != email {
			t.Errorf("Email missmatch: expected %q, but got %q", email, i.Email())
		}
	})
}

func TestUpdateUserOutput(t *testing.T) {
	t.Run("should initialize UpdateUserOutput", func(t *testing.T) {
		o := NewUpdateUserOutput()

		if *o != (UpdateUserOutput{}) {
			t.Errorf("failed to instance of DeleteUserOutput: instance %v", *o)
		}
	})
}
