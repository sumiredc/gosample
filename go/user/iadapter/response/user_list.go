package response

import "sample/user/domain/entity"

type UserListResponse struct {
	Users []*entity.User `json:"users"`
}
