package valueobject

import (
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

// func TestMarshalJson(t *testing.T) {
// 	t.Run("Marchal to JSON from UserID", func(t *testing.T) {
// 		expected := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
// 		userID, _ := ParseUserID(expected)

// 		type A struct {
// 			UserID UserID `json:"user_id"`
// 		}

// 		e := A{UserID: *userID}

// 		v, err := json.Marshal(e)
// 		if err != nil || fmt.Sprint(string(v)) != expected {
// 			t.Errorf("failed to convert to JSON fron UserID [expected: %s, UserID: %v, err: %v]", expected, userID, err)
// 		}
// 	})
// }
