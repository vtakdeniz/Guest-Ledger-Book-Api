package main

import (
	"guestLedgerBookApi/repo"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerAddsComments(t *testing.T) {
	repo := new(repo.MockRepo)
	app := initApp(repo)

	commentToPost := `{"email":"test@gmail.com","content":"testaaa"}`
	req := httptest.NewRequest("POST", "/api/comment", strings.NewReader(commentToPost))
	req.Header.Add("Content-Type", `application/json"`)
	res, _ := app.Test(req)
	assert.Equal(t, 200, res.StatusCode, "Is status code expected")

	reqGetComments := httptest.NewRequest("GET", "/api/comments", nil)
	reqGetComments.Header.Add("Content-Type", `application/json"`)
	resComments, _ := app.Test(reqGetComments)
	commentsByte, _ := ioutil.ReadAll(resComments.Body)
	commentsString := string(commentsByte)
	expectedRes := `[{"Id":1,"email":"test@gmail.com","content":"testaaa"}]`
	assert.JSONEq(t, expectedRes, commentsString, "Is get comments matches to added comments")
}

func TestServerReturnsCorrectCommentById(t *testing.T) {
	repo := new(repo.MockRepo)
	app := initApp(repo)

	commentToPost := `{"email":"test@gmail.com","content":"testaaa"}`
	req := httptest.NewRequest("POST", "/api/comment", strings.NewReader(commentToPost))
	req.Header.Add("Content-Type", `application/json"`)
	res, _ := app.Test(req)
	assert.Equal(t, 200, res.StatusCode, "Is status code expected")

	commentToPost2 := `{"email":"dddd@gmail.com","content":"xxxxxx"}`
	req2 := httptest.NewRequest("POST", "/api/comment", strings.NewReader(commentToPost2))
	req2.Header.Add("Content-Type", `application/json"`)
	res2, _ := app.Test(req2)
	assert.Equal(t, 200, res2.StatusCode, "Is status code expected")

	reqGetComments := httptest.NewRequest("GET", "/api/comment/1", nil)
	reqGetComments.Header.Add("Content-Type", `application/json"`)
	resComments, _ := app.Test(reqGetComments)
	commentsByte, _ := ioutil.ReadAll(resComments.Body)
	commentsString := string(commentsByte)
	expectedRes := `{"Id":1,"email":"test@gmail.com","content":"testaaa"}`
	assert.JSONEq(t, expectedRes, commentsString, "Is get comments matches to added comments")
}

func TestServerDeletesAllComments(t *testing.T) {
	repo := new(repo.MockRepo)
	app := initApp(repo)

	commentToPost := `{"email":"test@gmail.com","content":"testaaa"}`
	req := httptest.NewRequest("POST", "/api/comment", strings.NewReader(commentToPost))
	req.Header.Add("Content-Type", `application/json"`)
	res, _ := app.Test(req)
	assert.Equal(t, 200, res.StatusCode, "Is status code expected")

	commentToPost2 := `{"email":"dddd@gmail.com","content":"xxxxxx"}`
	req2 := httptest.NewRequest("POST", "/api/comment", strings.NewReader(commentToPost2))
	req2.Header.Add("Content-Type", `application/json"`)
	res2, _ := app.Test(req2)
	assert.Equal(t, 200, res2.StatusCode, "Is status code expected")

	reqDelete := httptest.NewRequest("DELETE", "/api/comments", nil)
	resComments, _ := app.Test(reqDelete)
	assert.Equal(t, 200, resComments.StatusCode, "Is get comments matches to added comments")
}
