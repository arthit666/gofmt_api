package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {

	stmt, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("can't prepare query all users statment", err)
	}
	println(stmt)

	rows, err := stmt.Query() // all many data  many row
	if err != nil {
		log.Fatal("can't query all users", err)
	}

	users := []User{}

	for rows.Next() {

		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)

}

func GetOneUserHandler(c echo.Context) error {
	id := c.Param("id")

	stmt, err := db.Prepare("SELECT id, name, age FROM users where id=$1") //>>return statement

	if err != nil {
		log.Fatal("can't prepare query all users statment", err)
	}

	row := stmt.QueryRow(id)

	u := User{}
	err = row.Scan(&u.Id, &u.Name, &u.Age)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "no id")
	}

	return c.JSON(http.StatusOK, u)
}
