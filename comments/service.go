package comments

import (
	"github.com/gofiber/fiber/v2"
)

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{
		repository: repo,
	}
}

type Repository interface {
	GetComments() ([]Comment, error)
	GetComment(id int) (Comment, error)
	AddComment(comment Comment) error
	DeleteAllComments() error
}

func (s *service) GetComments(ctx *fiber.Ctx) ([]Comment, error) {
	comments, err := s.repository.GetComments()

	if err != nil {
		return []Comment{}, err
	}
	return comments, nil
}

func (s *service) GetComment(ctx *fiber.Ctx, id int) (Comment, error) {
	comment, err := s.repository.GetComment(id)
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func (s *service) AddComment(ctx *fiber.Ctx, comment Comment) error {
	err := s.repository.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteAllComments(ctx *fiber.Ctx) error {
	err := s.repository.DeleteAllComments()
	if err != nil {
		return err
	}
	return nil
}
