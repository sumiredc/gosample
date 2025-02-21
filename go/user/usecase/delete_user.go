package usecase

import (
	"sample/user/domain/repository"
	"sample/user/usecase/dto"
)

type DeleteUserUseCase struct {
	userRepository repository.UserRepository
}

func NewDeleteUserUseCase(userRepository repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: userRepository,
	}
}

func (u *DeleteUserUseCase) Execute(i dto.DeleteUserInput) (*dto.DeleteUserOutput, error) {
	return nil, nil
}
