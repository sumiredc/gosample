package mysql

import (
	"sample/user/domain/entity"
	"sample/user/domain/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Writer *gorm.DB
	Reader *gorm.DB
}

func NewUserRepository(w, r *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Writer: w,
		Reader: r,
	}
}

func (u *UserRepositoryImpl) List() []entity.User {
	models := make([]model.User, 0)
	users := make([]entity.User, 0)

	u.Reader.Model(&model.User{}).Find(&models)

	for _, model := range models {
		users = append(users, entity.NewUser(model.ID, model.Name, model.Email))
	}

	return users
}

func (*UserRepositoryImpl) Get() *entity.User {
	return nil
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
