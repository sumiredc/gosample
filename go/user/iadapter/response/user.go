package response

import "sample/user/domain/entity"

type UserResponse struct {
	User *entity.User `json:"user"`
}
