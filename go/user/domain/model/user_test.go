package model

import (
	"sample/user/domain/valueobject"
	"testing"
)

func TestUserModel(t *testing.T) {
	userID := valueobject.NewUserID()
	name := "sumire"
	email := "sumire@test.xxx"
	tableName := "user"

	t.Run("should initialize User with expected values", func(t *testing.T) {
		user := NewUser(userID, name, email)

		if user.ID != userID.String() {
			t.Errorf("ID missmatch: expected %q, but got %q", user.ID, userID.String())
		}

		if user.Name != name {
			t.Errorf("Name missmatch: expected %q, but got %q", user.Name, name)
		}

		if user.Email != email {
			t.Errorf("Email missmatch: expected %q, but got %q", user.Email, email)
		}

		if user.TableName() != tableName {
			t.Errorf("TableName missmatch: expected %q, but got %q", tableName, user.TableName())
		}
	})
}
