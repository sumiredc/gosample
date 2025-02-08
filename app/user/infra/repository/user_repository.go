package repository

import "sample/app/user/domain/entity"

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (*UserRepositoryImpl) List() ([]entity.User, error) {
	var list []entity.User

	return list, nil
}

func (*UserRepositoryImpl) Get() (*entity.User, error) {
	return nil, nil
}

func (*UserRepositoryImpl) Save() error {
	return nil
}

func (*UserRepositoryImpl) Update(u *entity.User) error {
	return nil
}

func (*UserRepositoryImpl) Delete(u *entity.User) error {
	return nil
}
