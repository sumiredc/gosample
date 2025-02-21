package usecase

import (
	"errors"
	"sample/user/domain/errkit"
	"sample/user/domain/repository"
	"sample/user/domain/valueobject"
	"sample/user/usecase/dto"
)

type GetUserUseCase struct {
	userRepository repository.UserRepository
}

func NewGetUserUseCase(userRepository repository.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		userRepository: userRepository,
	}
}

func (u *GetUserUseCase) Execute(i dto.GetUserInput) (*dto.GetUserOutput, error) {
	userId, err := valueobject.ParseUserID(i.UserId())

	if err != nil {
		return nil, errors.Join(errkit.ErrBadRequest, err)
	}

	user, err := u.userRepository.Get(*userId)

	if err != nil {
		return nil, err
	}

	o := dto.NewGetUserOutput(*user)

	return o, nil
}
