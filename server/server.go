package server

import (
	"database/sql"

	"bitbucket.org/SealTV/go-site/server/v1"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strings"
)

func RunServer(db *sql.DB) {
	e := echo.New()

	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))

	// this logs the server interaction
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/jwt")
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))

	jwtGroup.GET("/main", v1.MainJwt)

	// user
	jwtGroup.GET("/user", v1.GetUser(db))
	jwtGroup.PUT("/user", v1.UpdateUser(db))
	jwtGroup.DELETE("/user/:id", v1.DeleteUser(db))

	// list
	jwtGroup.GET("/list", v1.GetList(db))
	jwtGroup.POST("/list", v1.AddList(db))
	jwtGroup.DELETE("/list/:id", v1.DeleteList(db))

	// todos
	jwtGroup.GET("/tasks", v1.GetTodos(db))
	jwtGroup.POST("/tasks", v1.AddTodo(db))
	jwtGroup.PUT("/tasks", v1.UpdateTodo(db))
	jwtGroup.DELETE("/tasks/:id", v1.DeleteTodo(db))

	cookieGroup.Use(checkCookie)
	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)

	e.POST("/register", v1.Register(db))
	e.GET("/login", v1.Login(db))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.File("/index", "static/index.html")
	e.File("/todo", "static/todo.html")

	e.Logger.Fatal(e.Start(":1323"))
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you don't have any cookie")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you don't have the right cookie, cookie")
	}
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret amdin main page!")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie page!")
}
