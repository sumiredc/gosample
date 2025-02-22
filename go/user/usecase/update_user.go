package usecase

import (
	"errors"
	"sample/user/domain/errkit"
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

func (u *UpdateUserUseCase) Execute(i dto.UpdateUserInput) (*dto.UpdateUserOutput, error) {
	userId, err := valueobject.ParseUserID(i.UserId())

	if err != nil {
		return nil, errors.Join(errkit.ErrBadRequest, err)
	}

	m := model.NewUser(*userId, i.Name(), i.Email())

	if err := u.userRepository.Update(m); err != nil {
		return nil, err
	}

	o := dto.NewUpdateUserOutput()

	return o, nil
}
