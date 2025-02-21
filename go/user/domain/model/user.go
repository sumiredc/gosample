package model

import "sample/user/domain/valueobject"

type User struct {
	ID    string `gorm:"primaryKey;size:26"`
	Name  string
	Email string
}

func NewUser(id valueobject.UserID, name, email string) *User {
	return &User{
		ID:    id.Value().String(),
		Name:  name,
		Email: email,
	}
}

func (*User) TableName() string {
	return "user"
}
