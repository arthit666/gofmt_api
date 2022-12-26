//go:build integration

package main

import (
	"arthit/api/user"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGetAllUser(t *testing.T) {

// 	seedUser(t)

// 	var us []user.User
// 	res := request(http.MethodGet, uri("users"), nil)

// 	err := res.Decode(&us)

// 	assert.Nil(t, err)
// 	assert.EqualValues(t, http.StatusOK, res.StatusCode)
// 	assert.Greater(t, len(us), 0)
// }

// func TestCreateUser(t *testing.T) {
// 	body := bytes.NewBufferString(`{
// 		"name": "anuchito",
// 		"age": 19
// 	}`)

// 	var u user.User

// 	res := request(http.MethodPost, uri("users"), body)
// 	err := res.Decode(&u)

// 	println(u.Id)

// 	assert.Nil(t, err)
// 	assert.Equal(t, http.StatusCreated, res.StatusCode)
// 	assert.NotEqual(t, 0, u.Id)
// 	assert.Equal(t, "anuchito", u.Name)
// 	assert.Equal(t, 19, u.Age)
// }

func TestGetUserById(t *testing.T) {
	c := seedUser(t)

	var userId user.User
	res := request(http.MethodGet, uri("users", strconv.Itoa(c.Id)), nil)

	err := res.Decode(&userId)

	assert.Nil(t, err) //

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.Id, userId.Id)
	assert.NotEmpty(t, userId.Name)
	assert.NotEmpty(t, userId.Age)

}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func seedUser(t *testing.T) user.User {
	var c user.User
	body := bytes.NewBufferString(`{
			"name": "anuchito",
			"age" : 19
		}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c) //decode assign c
	if err != nil {
		t.Fatal("can't create uomer:", err)
	}
	return c
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err} //wrap res and err is return form sever
}
