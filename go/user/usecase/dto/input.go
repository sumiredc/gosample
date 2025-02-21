package dto

type GetUserListInput struct {
}

func NewGetUserListInput() *GetUserListInput {
	return &GetUserListInput{}
}

type GetUserInput struct {
	userId string
}

func NewGetUserInput(userId string) *GetUserInput {
	return &GetUserInput{
		userId: userId,
	}
}

func (i *GetUserInput) UserId() string {
	return i.userId
}

type CreateUserInput struct {
	Name  string
	Email string
}

func NewCreateUserInput(name, email string) *CreateUserInput {
	return &CreateUserInput{
		Name:  name,
		Email: email,
	}
}

type UpdateUserInput struct {
	userId string
	Name   string
	Email  string
}

func NewUpdateUserInput(userId, name, email string) *UpdateUserInput {
	return &UpdateUserInput{
		userId: userId,
	}
}

func (i *UpdateUserInput) UserId() string {
	return i.userId
}

type DeleteUserInput struct {
	userId string
}

func NewDeleteUserInput(userId string) *DeleteUserInput {
	return &DeleteUserInput{
		userId: userId,
	}
}

func (i *DeleteUserInput) UserId() string {
	return i.userId
}
