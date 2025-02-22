package model

import (
	"sample/user/domain/valueobject"
	"testing"
)

func TestNewUser(t *testing.T) {
	userID := valueobject.NewUserID()
	name := "sumire"
	email := "sumire@test.xxx"

	t.Run("Make to instance of User model", func(t *testing.T) {
		user := NewUser(userID, name, email)

		if user.ID != userID.String() {
			t.Errorf("don't match `userID` [user.ID: %s, userID: %s]", user.ID, userID.String())
		}

		if user.Name != name {
			t.Errorf("don't match `name` [user.Name: %s, name: %s]", user.Name, name)
		}

		if user.Email != email {
			t.Errorf("don't match `email` [user.Email: %s, email: %s]", user.Email, email)
		}
	})
}

func TestTableName(t *testing.T) {
	expected := "user"
	user := NewUser(valueobject.NewUserID(), "name", "email")
	t.Run("Verify to table name in User model", func(t *testing.T) {
		if user.TableName() != expected {
			t.Errorf("don't match `TableName` [expected: %s, user.TableName(): %s]", expected, user.TableName())
		}
	})
}
