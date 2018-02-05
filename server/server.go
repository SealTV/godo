package server

import (
	"fmt"
	"net/http"

	"bitbucket.org/SealTV/go-site/data"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Server - server object
type Server struct {
	db data.DBConnector
	e  *echo.Echo
	c  Config
}

type Config struct {
	SecretKey string `json:"secret"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
}

func New(db data.DBConnector, c Config) *Server {
	e := echo.New()
	s := Server{db, e, c}
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
		SigningKey:    []byte(c.SecretKey),
	}))

	jwtGroup.GET("/main", s.mainJwt)

	// user
	jwtGroup.GET("/user", s.getUser)
	jwtGroup.PUT("/user", s.updateUser)
	jwtGroup.DELETE("/user/:id", s.deleteUser)

	// list
	jwtGroup.GET("/list", s.getLists)
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
	e.File("/list", "static/list.html")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", c.Host, c.Port)))

	return &s
}

func (s *Server) Run() {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port)))
}
