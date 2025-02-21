package presenter

import (
	"sample/user/domain/entity"
	"sample/user/iadapter/response"
)

func (*Presenter) User(user entity.User) response.Response {
	return &response.UserResponse{
		User: user,
	}
}
