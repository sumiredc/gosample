package usecase

import (
	"sample/user/domain/repository"
	"sample/user/usecase/dto"
)

type GetUserListUseCase struct {
	userRepository repository.UserRepository
}

func NewGetUserListUseCase(userRepository repository.UserRepository) *GetUserListUseCase {
	return &GetUserListUseCase{
		userRepository: userRepository,
	}
}

func (u *GetUserListUseCase) Execute(i dto.GetUserListInput) (*dto.GetUserListOutput, error) {
	users, err := u.userRepository.List()

	if err != nil {
		return nil, err
	}

	o := dto.NewGetUserListOutput(users)
	return o, nil
}
