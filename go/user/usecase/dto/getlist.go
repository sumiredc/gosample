package dto

import "sample/user/domain/entity"

type GetUserListInput struct {
}

func NewGetUserListInput() *GetUserListInput {
	return &GetUserListInput{}
}

type GetUserListOutput struct {
	users []entity.User
}

func NewGetUserListOutput(users []entity.User) *GetUserListOutput {
	return &GetUserListOutput{
		users: users,
	}
}

func (o *GetUserListOutput) Users() []entity.User {
	return o.users
}
