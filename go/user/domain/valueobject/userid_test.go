package valueobject

import (
	"encoding/json"
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestNewUserID(t *testing.T) {
	t.Run("Make to UserID value object", func(t *testing.T) {
		userID := NewUserID()
		_, err := ulid.Parse(userID.Value().String())

		if err != nil {
			t.Errorf("UserID is not ulid [UserID: %v, err: %v]", userID.Value().String(), err)
		}
	})
}

func TestParseUserID(t *testing.T) {
	t.Run("Parse to UserID value object", func(t *testing.T) {
		expected := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		_, err := ParseUserID(expected)

		if err != nil {
			t.Errorf("failed to parse ulid to UserID [expected: %s, err: %v]", expected, err)
		}
	})
}

func TestValue(t *testing.T) {
	t.Run("Value to UserID", func(t *testing.T) {
		expected := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		userID, err := ParseUserID(expected)

		if err != nil || userID.Value().String() != expected {
			t.Errorf("failed to convert to UserID from ULID [expected: %s, UserID: %v, err: %v]", expected, userID, err)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("String to UserID", func(t *testing.T) {
		expected := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		userID, err := ParseUserID(expected)

		if err != nil || userID.String() != expected {
			t.Errorf("failed to convert to UserID from ULID [expected: %s, UserID: %v, err: %v]", expected, userID, err)
		}
	})
}

func TestMarshalJSON(t *testing.T) {
	t.Run("Marchal to JSON from UserID", func(t *testing.T) {
		id := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		expected := `"` + id + `"`
		userID, _ := ParseUserID(id)

		v, err := json.Marshal(userID)

		if err != nil {
			t.Errorf("failed to convert to JSON fron UserID [err: %v]", err)
		}

		if string(v) != expected {
			t.Errorf("failed to convert to JSON fron UserID [expected: %s, UserID: %v]", expected, string(v))
		}
	})
}
