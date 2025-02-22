package dto

import "testing"

func TestDeleteUserInput(t *testing.T) {
	userID := "01JMPR93HT4GN3GDBMFP91C8J9"

	t.Run("should initialize DeleteUserInput with expected values", func(t *testing.T) {
		i := NewDeleteUserInput(userID)

		if i.UserID() != userID {
			t.Errorf("UserID missmatch: expected %q, but got %q", userID, i.UserID())
		}
	})
}

func TestDeleteUserOutput(t *testing.T) {
	t.Run("should initialize DeleteUserOutput", func(t *testing.T) {
		o := NewDeleteUserOutput()

		if *o != (DeleteUserOutput{}) {
			t.Errorf("failed to instance of DeleteUserOutput: instance %v", *o)
		}
	})
}
