package comments

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Email   string `json:"email"`
	Content string `json:"content"`
}

type CommentDto struct {
	ID      int    `json:"ID"`
	Email   string `json:"email"`
	Content string `json:"content"`
}
