package domain

import (
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
	Age  int64  `json:"age"`
}

func (u UserRequestAdd) MapToDomain() Users {
	return Users{
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

func (i *IncrRequest) ValidationRedis() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Key, validation.Required, validation.Length(2, 10)),
		validation.Field(&i.Value, validation.Required, validation.Min(0)),
		validation.Field(&i.Value, validation.Required, validation.Max(1000)),
	)
}

func (h *Ihmacsha512Request) ValidationHmac() error {
	return validation.ValidateStruct(h,
		validation.Field(&h.Text, validation.Required),
		validation.Field(&h.Key, validation.Required),
	)
}
