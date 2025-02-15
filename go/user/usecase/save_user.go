package usecase

type SaveUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SaveUserUseCase struct {
	Input *SaveUserInput
}

func NewSaveUserUseCase(input *SaveUserInput) *SaveUserUseCase {
	return &SaveUserUseCase{
		Input: input,
	}
}

func (u *SaveUserUseCase) Execute() {

}
