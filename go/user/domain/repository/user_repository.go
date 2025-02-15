package repository

import "sample/user/domain/entity"

type UserRepository interface {
	List() []entity.User
	Get() *entity.User
	Save() error
	Update(u *entity.User) error
	Delete(u *entity.User) error
}
