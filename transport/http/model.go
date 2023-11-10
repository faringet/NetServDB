package http

import (
	"NetServDB/domain"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type RequestAddUser struct {
}

type IncrRequest struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type Ihmacsha512Request struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type UserRequestAdd struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func (u UserRequestAdd) MapToDomain() domain.Users {
	return domain.Users{
		Model: gorm.Model{},
		Name:  u.Name,
		Age:   u.Age,
	}
}

func (u *UserRequestAdd) Validation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(5, 30)),
	)
}
