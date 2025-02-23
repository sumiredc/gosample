package entity

import "sample/user/domain/valueobject"

type User struct {
	ID    valueobject.UserID `json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

func NewUser(id valueobject.UserID, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
