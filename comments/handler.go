package comments

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	handlerService HandlerService
}

func NewHandler(service HandlerService) *Handler {
	return &Handler{
		handlerService: service,
	}
}

type HandlerService interface {
	GetComments(ctx *fiber.Ctx) ([]Comment, error)
	GetComment(ctx *fiber.Ctx, id int) (Comment, error)
	AddComment(ctx *fiber.Ctx, comment Comment) error
	DeleteAllComments(ctx *fiber.Ctx) error
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/comments", h.GetComments)
	app.Get("/comments/:id", h.GetComment)
	app.Post("/comments", h.AddComment)
	app.Delete("/comments", h.DeleteAllComments)
}

func (h *Handler) GetComments(ctx *fiber.Ctx) error {
	res, err := h.handlerService.GetComments(ctx)
	if err != nil {
		return errors.New("internal server error")
	}
	return ctx.Status(fiber.StatusOK).JSON(ToCommentResponseList(res))
}

func (h *Handler) GetComment(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return errors.New("error parsing url query")
	}
	res, err := h.handlerService.GetComment(ctx, id)
	if err != nil {
		return errors.New("internal server error")
	}
	return ctx.Status(fiber.StatusOK).JSON(ToCommentResponse(res))
}

func (h *Handler) AddComment(ctx *fiber.Ctx) error {
	var comment Comment
	if err := ctx.BodyParser(&comment); err != nil {
		return errors.New("error parsing request body")
	}
	err := h.handlerService.AddComment(ctx, comment)
	if err != nil {
		return errors.New("internal Server Error")
	}
	ctx.Status(fiber.StatusOK)
	return nil
}

func (h *Handler) DeleteAllComments(ctx *fiber.Ctx) error {
	err := h.handlerService.DeleteAllComments(ctx)
	if err != nil {
		return errors.New("internal server error")
	}
	ctx.Status(fiber.StatusOK)
	return nil
}

func ToCommentResponse(comment Comment) *CommentDto {
	return &CommentDto{
		ID:      int(comment.ID),
		Email:   comment.Email,
		Content: comment.Content,
	}
}

func ToCommentResponseList(comments []Comment) []CommentDto {
	var dto []CommentDto

	for _, val := range comments {
		dto = append(dto, *ToCommentResponse(val))
	}

	return dto
}
