package usecase

import (
	"sample/user/domain/repository"
	"sample/user/usecase/dto"
)

type UpdateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewUpdateUserUseCase(userRepository repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: userRepository,
	}
}

func (u *UpdateUserUseCase) Execute(i dto.UpdateUserInput) (*dto.UpdateUserOutput, error) {
	return nil, nil
}
