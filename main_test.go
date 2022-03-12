package main

import (
	"guestLedgerBookApi/comments"
	"guestLedgerBookApi/database"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestServerAddsComments(t *testing.T) {
	app := fiber.New()
	dbConfig := database.DbConfig{
		DbType: "sqlite3",
		Db:     "fake.db",
	}

	repo := database.NewRepo(dbConfig)

	service := comments.NewService(repo)

	handler := comments.NewHandler(service)

	handler.RegisterRoutes(app)

	commentToPost := `{"email":"test@gmail.com","content":"testaaa"}`
	req := httptest.NewRequest("POST", "/comments", strings.NewReader(commentToPost))
	req.Header.Add("Content-Type", `application/json"`)
	res, _ := app.Test(req)
	assert.Equal(t, 200, res.StatusCode, "Is status code expected")

	reqGetComments := httptest.NewRequest("GET", "/comments", nil)
	reqGetComments.Header.Add("Content-Type", `application/json"`)
	resComments, _ := app.Test(reqGetComments)
	commentsByte, _ := ioutil.ReadAll(resComments.Body)
	commentsString := string(commentsByte)
	expectedRes := `[{"ID":1,"email":"test@gmail.com","content":"testaaa"}]`
	assert.JSONEq(t, expectedRes, commentsString, "Is get comments matches to added comments")

	err := os.Remove("fake.db")
	if err != nil {
		log.Fatal(err)
	}
}

func TestServerReturnsCorrectCommentById(t *testing.T) {
	app := fiber.New()
	dbConfig := database.DbConfig{
		DbType: "sqlite3",
		Db:     "fake.db",
	}

	repo := database.NewRepo(dbConfig)

	service := comments.NewService(repo)

	handler := comments.NewHandler(service)

	handler.RegisterRoutes(app)

	commentToPost := `{"email":"test@gmail.com","content":"testaaa"}`
	req := httptest.NewRequest("POST", "/comments", strings.NewReader(commentToPost))
	req.Header.Add("Content-Type", `application/json"`)
	res, _ := app.Test(req)
	assert.Equal(t, 200, res.StatusCode, "Is status code expected")

	commentToPost2 := `{"email":"dddd@gmail.com","content":"xxxxxx"}`
	req2 := httptest.NewRequest("POST", "/comments", strings.NewReader(commentToPost2))
	req2.Header.Add("Content-Type", `application/json"`)
	res2, _ := app.Test(req2)
	assert.Equal(t, 200, res2.StatusCode, "Is status code expected")

	reqGetComments := httptest.NewRequest("GET", "/comments/1", nil)
	reqGetComments.Header.Add("Content-Type", `application/json"`)
	resComments, _ := app.Test(reqGetComments)
	commentsByte, _ := ioutil.ReadAll(resComments.Body)
	commentsString := string(commentsByte)
	expectedRes := `{"ID":1,"email":"test@gmail.com","content":"testaaa"}`
	assert.JSONEq(t, expectedRes, commentsString, "Is get comments matches to added comments")

	err := os.Remove("fake.db")
	if err != nil {
		log.Fatal(err)
	}
}

func TestServerDeletesAllComments(t *testing.T) {
	app := fiber.New()
	dbConfig := database.DbConfig{
		DbType: "sqlite3",
		Db:     "fake.db",
	}

	repo := database.NewRepo(dbConfig)

	service := comments.NewService(repo)

	handler := comments.NewHandler(service)

	handler.RegisterRoutes(app)

	commentToPost := `{"email":"test@gmail.com","content":"testaaa"}`
	req := httptest.NewRequest("POST", "/comments", strings.NewReader(commentToPost))
	req.Header.Add("Content-Type", `application/json"`)
	res, _ := app.Test(req)
	assert.Equal(t, 200, res.StatusCode, "Is status code expected")

	commentToPost2 := `{"email":"dddd@gmail.com","content":"xxxxxx"}`
	req2 := httptest.NewRequest("POST", "/comments", strings.NewReader(commentToPost2))
	req2.Header.Add("Content-Type", `application/json"`)
	res2, _ := app.Test(req2)
	assert.Equal(t, 200, res2.StatusCode, "Is status code expected")

	reqDelete := httptest.NewRequest("DELETE", "/comments", nil)
	resComments, _ := app.Test(reqDelete)
	assert.Equal(t, 200, resComments.StatusCode, "Is get comments matches to added comments")

	err := os.Remove("fake.db")
	if err != nil {
		log.Fatal(err)
	}
}
