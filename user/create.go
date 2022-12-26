package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUsersHandler(c echo.Context) error {
	u := User{}

	err := c.Bind(&u) //unmashall

	if err != nil {
		c.JSON(http.StatusBadRequest, Err{Massage: err.Error()})

	}

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2)  RETURNING id", u.Name, u.Age)
	err = row.Scan(&u.Id) //scan = asign to variable

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Massage: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}
