package repository

import "sample/app/user/domain/entity"

type UserRepository interface {
	List() ([]entity.User, error)
	Get() (*entity.User, error)
	Save() error
	Update(u *entity.User) error
	Delete(u *entity.User) error
}
