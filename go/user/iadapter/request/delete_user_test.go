package request_test

import (
	"sample/user/iadapter/request"
	"testing"

	"github.com/go-playground/validator"
)

func TestDeleteUserRequest(t *testing.T) {
	v := validator.New()

	t.Run("should valid to request successfully", func(t *testing.T) {
		req := &request.DeleteUserRequest{}

		err := v.Struct(req)

		if err != nil {
			t.Errorf("valid to DeleteUserRequest: err %v", err)
		}
	})
}
