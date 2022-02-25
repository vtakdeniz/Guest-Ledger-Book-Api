package repo

import "github.com/gofiber/fiber/v2"

type IRepo interface {
	GetComments(ctx *fiber.Ctx) error
	GetComment(ctx *fiber.Ctx) error
	AddComment(ctx *fiber.Ctx) error
	DeleteAllComments(ctx *fiber.Ctx) error
	InitDB()
}
