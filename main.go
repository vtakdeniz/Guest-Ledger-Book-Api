package main

import (
	"fmt"
	"guestLedgerBookApi/database"
	model "guestLedgerBookApi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setRoutes(app *fiber.App) {
	app.Get("/api/comments", model.GetComments)
	app.Get("/api/comment/:id", model.GetComment)
	app.Post("/api/comment", model.AddComment)
	app.Delete("/api/comments", model.DeleteAllComments)
}

func initDB() {
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

func main() {
	port := ":9000"
	app := fiber.New()
	initDB()
	defer database.DBGorm.Close()
	app.Use(cors.New())
	setRoutes(app)
	app.Listen(port)
}
