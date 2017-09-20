package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.File("/index", "public/view/index.html")
	e.File("/todo", "public/view/todo.html")
	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":1323"))
}
