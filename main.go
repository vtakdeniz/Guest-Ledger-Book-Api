package main

import (
	"fmt"
	"guestLedgerBookApi/repo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initApp(r repo.IRepo) *fiber.App {
	app := fiber.New()
	r.InitDB()

	app.Get("/api/comments", r.GetComments)
	app.Get("/api/comment/:id", r.GetComment)
	app.Post("/api/comment", r.AddComment)
	app.Delete("/api/comments", r.DeleteAllComments)

	app.Use(cors.New())
	return app
}

func main() {
	repo := new(repo.MockRepo)
	app := initApp(repo)
	port := 8080
	app.Listen(fmt.Sprintf(":%d", port))
}
