package repo

import (
	"sort"

	"github.com/gofiber/fiber/v2"

	model "guestLedgerBookApi/models"
)

var db map[int]model.Comment
var globalId int

type MockRepo struct {
}

func (r *MockRepo) InitDB() {
	db = make(map[int]model.Comment)
	globalId = 1
}

func (mr *MockRepo) GetComments(ctx *fiber.Ctx) error {
	keys := make([]int, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var comments []model.MockComment
	for _, k := range keys {
		comm := db[k]
		mockCom := model.MockComment{
			Id:      k,
			Email:   comm.Email,
			Content: comm.Content,
		}
		comments = append(comments, mockCom)
	}

	return ctx.JSON(comments)
}

func (mr *MockRepo) GetComment(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	if comment, ok := db[id]; ok {
		mockCom := model.MockComment{
			Id:      id,
			Email:   comment.Email,
			Content: comment.Content,
		}
		return ctx.JSON(mockCom)
	}
	return ctx.SendStatus(fiber.StatusNotFound)
}

func (mr *MockRepo) AddComment(ctx *fiber.Ctx) error {
	comment := new(model.Comment)
	if err := ctx.BodyParser(comment); err != nil {
		ctx.Status(503)
		return err
	}
	db[globalId] = *comment
	mockComm := model.MockComment{
		Id:      globalId,
		Email:   comment.Email,
		Content: comment.Content,
	}
	globalId++
	return ctx.JSON(mockComm)
}

func (mr *MockRepo) DeleteAllComments(ctx *fiber.Ctx) error {
	for key, _ := range db {
		delete(db, key)
	}
	return nil
}
