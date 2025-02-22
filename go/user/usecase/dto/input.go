package dto

type GetUserListInput struct {
}

func NewGetUserListInput() *GetUserListInput {
	return &GetUserListInput{}
}

type GetUserInput struct {
	userID string
}

func NewGetUserInput(userID string) *GetUserInput {
	return &GetUserInput{
		userID: userID,
	}
}

func (i *GetUserInput) UserId() string {
	return i.userID
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
	userID string
	name   string
	email  string
}

func NewUpdateUserInput(userID, name, email string) *UpdateUserInput {
	return &UpdateUserInput{
		userID: userID,
		name:   name,
		email:  email,
	}
}

func (i *UpdateUserInput) UserId() string {
	return i.userID
}

func (i *UpdateUserInput) Name() string {
	return i.name
}

func (i *UpdateUserInput) Email() string {
	return i.email
}

type DeleteUserInput struct {
	userID string
}

func NewDeleteUserInput(userID string) *DeleteUserInput {
	return &DeleteUserInput{
		userID: userID,
	}
}

func (i *DeleteUserInput) UserId() string {
	return i.userID
}
