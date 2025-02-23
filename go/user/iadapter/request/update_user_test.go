package request_test

import (
	"sample/user/iadapter/request"
	"strings"
	"testing"

	"github.com/go-playground/validator"
)

func TestUpdateUserRequest(t *testing.T) {
	v := validator.New()

	t.Run("should valid to request successfully", func(t *testing.T) {
		req := &request.UpdateUserRequest{
			Name:  strings.Repeat("a", 100),
			Email: strings.Repeat("a", 90) + "@test.xxxx",
		}

		err := v.Struct(req)

		if err != nil {
			t.Errorf("valid to UpdateUserRequest: err %v", err)
		}
	})

	type TestCase struct {
		TestDescription string
		Name            string
		Email           string
	}

	testCases := [5]*TestCase{
		{
			TestDescription: "over name length",
			Name:            strings.Repeat("a", 101),
			Email:           "safe@test.xxxx",
		},
		{
			TestDescription: "over email length",
			Name:            "safe name",
			Email:           strings.Repeat("a", 91) + "@test.xxxx",
		},
		{
			TestDescription: "empty name",
			Name:            "",
			Email:           "safe@test.xxxx",
		},
		{
			TestDescription: "empty email",
			Name:            "safe name",
			Email:           "",
		},
		{
			TestDescription: "email missmatch email type",
			Name:            "safe name",
			Email:           "failed#email.xxx",
		},
	}

	for _, tc := range testCases {
		t.Run("should invalid to request", func(t *testing.T) {
			req := &request.UpdateUserRequest{
				Name:  tc.Name,
				Email: tc.Email,
			}

			err := v.Struct(req)

			if err == nil {
				t.Errorf("expected validation error for test case: %q, but got nil", tc.TestDescription)
			}
		})
	}
}
