package main

import (
	"fmt"
	"guestLedgerBookApi/comments"
	"guestLedgerBookApi/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func StartServer(port int, dbConfig database.DbConfig) (fiber.App, error) {

	app := fiber.New()
	app.Use(cors.New())

	repo := database.NewRepo(dbConfig)

	service := comments.NewService(repo)

	handler := comments.NewHandler(service)

	handler.RegisterRoutes(app)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return fiber.App{}, err
	}
	return *app, nil
}

func StartServerWithSeed(port int, dbConfig database.DbConfig, seed []comments.Comment) (fiber.App, error) {

	app := fiber.New()
	app.Use(cors.New())

	repo := database.NewRepo(dbConfig)

	service := comments.NewService(repo)

	handler := comments.NewHandler(service)

	handler.RegisterRoutes(app)

	for _, comment := range seed {
		repo.AddComment(comment)
	}

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return fiber.App{}, err
	}
	return *app, nil
}

func main() {

	dbConfig := database.DbConfig{
		DbType: "sqlite3",
		Db:     "comments.db",
	}

	port := 8080
	_, err := StartServer(port, dbConfig)

	if err != nil {
		log.Fatal("Error starting server")
	}
}
