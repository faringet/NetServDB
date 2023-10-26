package models

import "gorm.io/gorm"

type IncrRequest struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type Ihmacsha512Request struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type Users struct {
	gorm.Model
	Name string
	Age  uint
}
