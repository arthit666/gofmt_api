package user

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Err struct {
	Massage string `json:"massage"`
}

func BasicAuth(user, pass string, c echo.Context) (bool, error) {

	if user == "apidesign" && pass == "45678" {
		return true, nil
	}

	return false, nil
}
