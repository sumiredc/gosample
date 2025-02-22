package dto

import "sample/user/domain/entity"

type GetUserInput struct {
	userID string
}

func NewGetUserInput(userID string) *GetUserInput {
	return &GetUserInput{
		userID: userID,
	}
}

func (i *GetUserInput) UserID() string {
	return i.userID
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
