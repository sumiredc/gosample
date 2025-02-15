package usecase

import (
	"sample/user/domain/entity"
	"sample/user/domain/repository"
)

type GetUserListUseCase struct {
	userRepository repository.UserRepository
}

func NewGetUserListUseCase(userRepository repository.UserRepository) *GetUserListUseCase {
	return &GetUserListUseCase{
		userRepository: userRepository,
	}
}

func (u *GetUserListUseCase) Execute() ([]entity.User, error) {
	users := u.userRepository.List()

	return users, nil
}
