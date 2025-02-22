package usecase

import (
	"errors"
	"sample/user/domain/errkit"
	"sample/user/domain/repository"
	"sample/user/domain/valueobject"
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
	userID, err := valueobject.ParseUserID(i.UserId())
	if err != nil {
		return nil, errors.Join(errkit.ErrBadRequest, err)
	}

	if err := u.userRepository.Delete(*userID); err != nil {
		return nil, err
	}

	o := dto.NewDeleteUserOutput()

	return o, nil
}
