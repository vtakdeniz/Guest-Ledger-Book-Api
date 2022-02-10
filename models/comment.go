package model

import (
	"guestLedgerBookApi/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Email   string `json:"email"`
	Content string `json:"content"`
}

func GetComments(ctx *fiber.Ctx) error {
	db := database.DBGorm
	var comments []Comment
	db.Find(&comments)
	return ctx.JSON(comments)
}

func GetComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DBGorm
	var comment Comment
	db.Find(&comment, id)
	return ctx.JSON(comment)
}

func AddComment(ctx *fiber.Ctx) error {
	db := database.DBGorm
	comment := new(Comment)
	if err := ctx.BodyParser(comment); err != nil {
		ctx.Status(503)
		return err
	}
	db.Create(&comment)
	return ctx.JSON(comment)
}

func DeleteAllComments(ctx *fiber.Ctx) error {
	db := database.DBGorm
	db.Where("1 = 1").Delete(&Comment{})
	return ctx.Status(fiber.StatusOK).JSON(Comment{})
}
