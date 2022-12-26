package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"arthit/api/user"
)

func BasicAuth(user, pass string, c echo.Context) (bool, error) {

	if user == "apidesign" && pass == "45678" {
		return true, nil
	}

	return false, nil
}

func main() {

	user.InitDB()

	e := echo.New()

	// defer db.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(BasicAuth))

	e.GET("/users", user.GetUsersHandler)
	e.POST("/users", user.CreateUsersHandler)

	e.GET("/users/:id", user.GetOneUserHandler)

	log.Println("sever start")
	log.Fatal(e.Start(":2565"))

}
