package server

import (
	"net/http"

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

	jwtGroup.GET("/main", s.mainJwt)

	// user
	jwtGroup.GET("/user", s.getUser)
	jwtGroup.PUT("/user", s.updateUser)
	jwtGroup.DELETE("/user/:id", s.deleteUser)

	// list
	jwtGroup.GET("/list", s.getList)
	jwtGroup.POST("/list", s.addList)
	jwtGroup.DELETE("/list/:id", s.deleteList)

	// todos
	jwtGroup.GET("/tasks", s.getTodos)
	jwtGroup.POST("/tasks", s.addTodo)
	jwtGroup.PUT("/tasks", s.updateTodo)
	jwtGroup.DELETE("/tasks/:id", s.deleteTodo)

	adminGroup.GET("/main", mainAdmin)

	e.POST("/register", s.register)
	e.GET("/login", s.login)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.File("/index", "static/index.html")
	e.File("/todo", "static/todo.html")

	e.Logger.Fatal(e.Start(":1323"))
}
