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

type UpdateUserInput struct {
	userId string
	name   string
	email  string
}

func NewUpdateUserInput(userId, name, email string) *UpdateUserInput {
	return &UpdateUserInput{
		userId: userId,
		name:   name,
		email:  email,
	}
}

func (i *UpdateUserInput) UserId() string {
	return i.userId
}

func (i *UpdateUserInput) Name() string {
	return i.name
}

func (i *UpdateUserInput) Email() string {
	return i.email
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
