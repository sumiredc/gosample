package usecase

import (
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

func (u *GetUserUseCase) Execute(i *dto.GetUserInput) (*dto.GetUserOutput, error) {
	userID, err := valueobject.ParseUserID(i.UserID())

	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.Get(*userID)

	if err != nil {
		return nil, err
	}

	o := dto.NewGetUserOutput(user)

	return o, nil
}
