package dto

import "sample/user/domain/entity"

type CreateUserInput struct {
	name  string
	email string
}

func NewCreateUserInput(name, email string) *CreateUserInput {
	return &CreateUserInput{
		name:  name,
		email: email,
	}
}

func (i *CreateUserInput) Name() string {
	return i.name
}

func (i *CreateUserInput) Email() string {
	return i.email
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
