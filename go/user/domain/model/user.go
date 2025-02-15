package model

type User struct {
	ID    string `gorm:"primaryKey;size:26"`
	Name  string
	Email string
}

func (User) TableName() string {
	return "user"
}
