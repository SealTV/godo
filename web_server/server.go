package web_server

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
)

func RunServer(db *sql.DB) {
	e := echo.New()
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