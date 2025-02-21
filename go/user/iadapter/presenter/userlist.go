package presenter

import (
	"sample/user/domain/entity"
	"sample/user/iadapter/response"
)

func (*Presenter) UserList(users []entity.User) response.Response {
	return &response.UserListResponse{
		Users: users,
	}
}
