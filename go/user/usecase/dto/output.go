package dto

import "sample/user/domain/entity"

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

type GetUserOutput struct {
	user entity.User
}

func NewGetUserOutput(user entity.User) *GetUserOutput {
	return &GetUserOutput{
		user: user,
	}
}

func (o *GetUserOutput) User() entity.User {
	return o.user
}

type CreateUserOutput struct {
	user entity.User
}

func NewCreateUserOutput(user entity.User) *CreateUserOutput {
	return &CreateUserOutput{
		user: user,
	}
}

func (o *CreateUserOutput) User() entity.User {
	return o.user
}

type UpdateUserOutput struct {
}

func NewUpdateUserOutput() *UpdateUserOutput {
	return &UpdateUserOutput{}
}

type DeleteUserOutput struct {
}

func NewDeleteUserOutput() *DeleteUserOutput {
	return &DeleteUserOutput{}
}
