package main

import (
	"guestLedgerBookApi/comments"
	"guestLedgerBookApi/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	dbConfig := database.DbConfig{
		DbType: "sqlite3",
		Db:     "comments.db",
	}

	repo := database.NewRepo(dbConfig)

	service := comments.NewService(repo)

	handler := comments.NewHandler(service)

	handler.RegisterRoutes(app)

	app.Listen(":8080")
}
