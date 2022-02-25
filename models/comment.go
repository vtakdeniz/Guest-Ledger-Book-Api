package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Email   string `json:"email"`
	Content string `json:"content"`
}

type MockComment struct {
	Id      int    `json:Id`
	Email   string `json:"email"`
	Content string `json:"content"`
}
