package dto

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

func (i *UpdateUserInput) UserID() string {
	return i.userID
}

func (i *UpdateUserInput) Name() string {
	return i.name
}

func (i *UpdateUserInput) Email() string {
	return i.email
}

type UpdateUserOutput struct {
}

func NewUpdateUserOutput() *UpdateUserOutput {
	return &UpdateUserOutput{}
}
