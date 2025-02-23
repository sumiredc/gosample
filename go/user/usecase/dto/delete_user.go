package dto

type DeleteUserInput struct {
	userID string
}

func NewDeleteUserInput(userID string) *DeleteUserInput {
	return &DeleteUserInput{
		userID: userID,
	}
}

func (i *DeleteUserInput) UserID() string {
	return i.userID
}

type DeleteUserOutput struct {
}

func NewDeleteUserOutput() *DeleteUserOutput {
	return &DeleteUserOutput{}
}
