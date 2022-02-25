package repo

import (
	"fmt"
	"guestLedgerBookApi/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"

	model "guestLedgerBookApi/models"
)

type Repo struct {
}

func (r *Repo) InitDB() {
	var err error
	database.DBGorm, err = gorm.Open("sqlite3", "books.db")
	fmt.Println("Setting up db connection")
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database successfully connected")
	database.DBGorm.AutoMigrate(&model.Comment{})
	fmt.Print("Database migrated")
}

func (r *Repo) GetComments(ctx *fiber.Ctx) error {
	db := database.DBGorm
	var comments []model.Comment
	db.Find(&comments)
	return ctx.JSON(comments)
}

func (r *Repo) GetComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DBGorm
	var comment model.Comment
	db.Find(&comment, id)
	return ctx.JSON(comment)
}

func (r *Repo) AddComment(ctx *fiber.Ctx) error {
	db := database.DBGorm
	comment := new(model.Comment)
	if err := ctx.BodyParser(comment); err != nil {
		ctx.Status(503)
		return err
	}
	db.Create(&comment)
	return ctx.JSON(comment)
}

func (r *Repo) DeleteAllComments(ctx *fiber.Ctx) error {
	db := database.DBGorm
	db.Where("1 = 1").Delete(&model.Comment{})
	return ctx.Status(fiber.StatusOK).JSON(model.Comment{})
}
