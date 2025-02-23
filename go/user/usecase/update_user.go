package usecase

import (
	"sample/user/domain/model"
	"sample/user/domain/repository"
	"sample/user/domain/valueobject"
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

func (u *UpdateUserUseCase) Execute(i *dto.UpdateUserInput) (*dto.UpdateUserOutput, error) {
	userID, err := valueobject.ParseUserID(i.UserID())

	if err != nil {
		return nil, err
	}

	m := model.NewUser(*userID, i.Name(), i.Email())

	if err := u.userRepository.Update(m); err != nil {
		return nil, err
	}

	o := dto.NewUpdateUserOutput()

	return o, nil
}
