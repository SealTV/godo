package server

import (
	"log"
	"net/http"
	"strings"

	"bitbucket.org/SealTV/go-site/data"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	db data.DBConnector
}

func RunServer(db data.DBConnector) {
	e := echo.New()
	s := Server{db}
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

	jwtGroup.GET("/main", MainJwt)

	// user
	jwtGroup.GET("/user", GetUser(db))
	jwtGroup.PUT("/user", UpdateUser(db))
	jwtGroup.DELETE("/user/:id", DeleteUser(db))

	// list
	jwtGroup.GET("/list", GetList(db))
	jwtGroup.POST("/list", AddList(db))
	jwtGroup.DELETE("/list/:id", DeleteList(db))

	// todos
	jwtGroup.GET("/tasks", s.GetTodos)
	jwtGroup.POST("/tasks", AddTodo(db))
	jwtGroup.PUT("/tasks", UpdateTodo(db))
	jwtGroup.DELETE("/tasks/:id", DeleteTodo(db))

	cookieGroup.Use(checkCookie)
	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)

	e.POST("/register", Register(db))
	e.GET("/login", Login(db))

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
	return c.String(http.StatusOK, "hoary you are on the secret admin main page!")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie page!")
}
