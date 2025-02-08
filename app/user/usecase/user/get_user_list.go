package user

import (
	"log"
	"sample/app/user/domain/entity"
	"sample/app/user/domain/repository"
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
	users, err := u.userRepository.List()
	if err != nil {
		log.Printf("Failed UserRepository.List(): %v", err)
		return users, err
	}

	return users, nil
}
