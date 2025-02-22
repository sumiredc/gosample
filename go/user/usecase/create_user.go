package usecase

import (
	"sample/user/domain/model"
	"sample/user/domain/repository"
	"sample/user/domain/valueobject"
	"sample/user/usecase/dto"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(userRepository repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (u *CreateUserUseCase) Execute(i dto.CreateUserInput) (*dto.CreateUserOutput, error) {
	userID := valueobject.NewUserID()
	m := model.NewUser(userID, i.Name(), i.Email())

	user, err := u.userRepository.Create(m)

	if err != nil {
		return nil, err
	}

	o := dto.NewCreateUserOutput(*user)

	return o, nil
}
