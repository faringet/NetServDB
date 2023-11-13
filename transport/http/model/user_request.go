package model

import (
	"NetServDB/domain"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type UserRequestAdd struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func (u UserRequestAdd) MapToDomain() domain.Users {
	return domain.Users{
		Model: gorm.Model{},
		Name:  u.Name,
		Age:   uint(u.Age),
	}
}

func (u *UserRequestAdd) ValidationUser() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required, validation.Length(2, 15)),
		validation.Field(&u.Age, validation.Required, validation.Min(1)),
		validation.Field(&u.Age, validation.Required, validation.Max(120)),
	)
}
