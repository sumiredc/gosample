package valueobject

import (
	"encoding/json"
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestUserID(t *testing.T) {
	t.Run("should initialize UserID with expected values", func(t *testing.T) {
		userID := NewUserID()
		_, err := ulid.Parse(userID.String())

		if err != nil {
			t.Errorf("failed to parse UserID as ULID: value %q, error %v", userID.String(), err)
		}
	})

	t.Run("should parse UserID as ULID", func(t *testing.T) {
		expected := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		userID, err := ParseUserID(expected)

		if err != nil {
			t.Errorf("failed to parse UserID as ULID: value %q, err %v", expected, err)
		}

		if userID.Value().String() != expected {
			t.Errorf("UserID.Value() missmatch: expected %q, but got %q", expected, userID.Value())
		}

		if userID.String() != expected {
			t.Errorf("UserID.String() missmatch: expected %q, but got %q", expected, userID.String())
		}
	})

	testCases := [3]string{"d22bea62-00a2-4887-b06b-6ddef9f31499", "1", "string"}

	for _, c := range testCases {
		t.Run("should return an error when UserID is not a valid ULID", func(t *testing.T) {
			_, err := ParseUserID(c)

			if err == nil {
				t.Errorf("expected an error when parsing invalid UserID, but got none: value %q", c)
			}
		})
	}

	t.Run("should convert JSON as UserID", func(t *testing.T) {
		id := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		expected := `"` + id + `"`
		userID, _ := ParseUserID(id)

		v, err := json.Marshal(userID)

		if err != nil {
			t.Errorf("failed to convert JSON as UserID: err %v", err)
		}

		if string(v) != expected {
			t.Errorf("failed to convert JSON as UserID: expected %q, but got %q", expected, string(v))
		}
	})
}
