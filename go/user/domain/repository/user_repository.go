package repository

import (
	"sample/user/domain/entity"
	"sample/user/domain/model"
	"sample/user/domain/valueobject"
)

type UserRepository interface {
	List() ([]*entity.User, error)
	Get(id valueobject.UserID) (*entity.User, error)
	Create(u *model.User) (*entity.User, error)
	Update(u *model.User) error
	Delete(id valueobject.UserID) error
}
