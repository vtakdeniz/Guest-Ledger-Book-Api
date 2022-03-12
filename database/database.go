package database

import (
	"fmt"
	"guestLedgerBookApi/comments"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBGorm *gorm.DB
)

type Repository struct {
	db *gorm.DB
}

type DbConfig struct {
	Db     interface{}
	DbType string
}

func NewRepo(config DbConfig) *Repository {
	var err error
	DBGorm, err = gorm.Open(config.DbType, config.Db)
	fmt.Println("Setting up db connection")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully connected")
	DBGorm.AutoMigrate(comments.Comment{})
	fmt.Print("Database migrated")
	return &Repository{
		db: DBGorm,
	}
}

func (r *Repository) GetComments() ([]comments.Comment, error) {
	var comments []comments.Comment
	r.db.Find(&comments)
	return comments, nil
}

func (r *Repository) GetComment(id int) (comments.Comment, error) {
	var comment comments.Comment
	err := r.db.Find(&comment, id).Error
	if err != nil {
		return comments.Comment{}, err
	}
	return comment, nil
}

func (r *Repository) AddComment(comment comments.Comment) error {
	err := r.db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteAllComments() error {
	err := r.db.Where("1 = 1").Delete(&comments.Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}
