package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type IncrRequest struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (i *IncrRequest) ValidationRedis() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.Key, validation.Required, validation.Length(2, 10)),
		validation.Field(&i.Value, validation.Required, validation.Min(0)),
		validation.Field(&i.Value, validation.Required, validation.Max(1000)),
	)
}

type Ihmacsha512Request struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

func (h *Ihmacsha512Request) ValidationHmac() error {
	return validation.ValidateStruct(h,
		validation.Field(&h.Text, validation.Required),
		validation.Field(&h.Key, validation.Required),
	)
}
