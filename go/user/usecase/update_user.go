package usecase

type UpdateUserInput struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UpdateUserUseCase struct {
	Input *UpdateUserInput
}

func NewUpdateUserUseCase(input *UpdateUserInput) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		Input: input,
	}
}

func (u *UpdateUserUseCase) Execute() {

}
