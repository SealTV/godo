package server

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func RunServer(db *sql.DB) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.File("/index", "public/view/index.html")
	e.File("/todo", "public/view/todo.html")
	e.Static("/static", "assets")

	e.GET("/tasks", getTasks(db))
	e.PUT("/tasks", putTask(db))
	e.DELETE("/tasks/:id", deleteTask(db))

	e.Logger.Fatal(e.Start(":1323"))
}
