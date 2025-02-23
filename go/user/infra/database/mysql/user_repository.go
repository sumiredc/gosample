package mysql

import (
	"errors"
	"sample/user/domain/entity"
	"sample/user/domain/errkit"
	"sample/user/domain/model"
	"sample/user/domain/valueobject"

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

func (repo *UserRepositoryImpl) List() ([]*entity.User, error) {
	models := make([]*model.User, 0)
	users := make([]*entity.User, 0)

	err := repo.Reader.Model(&model.User{}).Find(&models).Error

	if err != nil {
		return nil, errors.Join(errkit.ErrDatabaseConnection, err)
	}

	for _, m := range models {
		userID, err := valueobject.ParseUserID(m.ID)

		if err != nil {
			return nil, errors.Join(errkit.ErrInternalServer, err)
		}

		users = append(users, entity.NewUser(*userID, m.Name, m.Email))
	}

	return users, nil
}

func (repo *UserRepositoryImpl) Get(id valueobject.UserID) (*entity.User, error) {
	var m *model.User

	err := repo.Reader.First(&m, "id = ?", id.Value().String()).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Join(errkit.ErrNotFound, err)
	}

	if err != nil {
		return nil, errors.Join(errkit.ErrDatabaseConnection, err)
	}

	user := entity.NewUser(id, m.Name, m.Email)

	return user, nil
}

func (repo *UserRepositoryImpl) Create(m *model.User) (*entity.User, error) {
	err := repo.Writer.Create(m).Error

	if err != nil {
		return nil, errors.Join(errkit.ErrDatabaseConnection, err)
	}

	userID, err := valueobject.ParseUserID(m.ID)

	if err != nil {
		return nil, errors.Join(errkit.ErrInternalServer, err)
	}

	user := entity.NewUser(*userID, m.Name, m.Email)

	return user, nil
}

func (repo *UserRepositoryImpl) Update(m *model.User) error {
	err := repo.Writer.Save(m).Error

	if err != nil {
		return errors.Join(errkit.ErrDatabaseConnection, err)
	}

	return nil
}

func (repo *UserRepositoryImpl) Delete(id valueobject.UserID) error {
	res := repo.Writer.Delete(&model.User{}, "id = ?", id.Value().String())

	if res.Error != nil {
		return errors.Join(errkit.ErrDatabaseConnection, res.Error)
	}

	if res.RowsAffected == 0 {
		return errkit.ErrNotFound
	}

	return nil
}
