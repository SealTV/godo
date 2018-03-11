package server

import (
	"fmt"

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

// Config - server config params
type Config struct {
	SecretKey string `json:"secret"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
}

// New - create new instance of servers
func New(db data.DBConnector, c Config) *Server {
	e := echo.New()
	s := Server{db, e, c}
	adminGroup := e.Group("/api/admin")

	// this logs the server interaction
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/api/jwt")
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(c.SecretKey),
	}))

	jwtGroup.GET("/api/main", s.mainJwt)

	// user
	jwtGroup.GET("/api/user", s.getUser)
	jwtGroup.PUT("/api/user", s.updateUser)
	jwtGroup.DELETE("/api/user/:id", s.deleteUser)

	// list
	jwtGroup.GET("/api/list", s.getLists)
	jwtGroup.POST("/api/list", s.addList)
	jwtGroup.DELETE("/api/list/:id", s.deleteList)

	// todos
	jwtGroup.GET("/api/tasks", s.getTodos)
	jwtGroup.POST("/api/tasks", s.addTodo)
	jwtGroup.PUT("/api/tasks", s.updateTodo)
	jwtGroup.DELETE("/api/tasks/:id", s.deleteTodo)

	adminGroup.GET("/api/main", mainAdmin)

	e.POST("/api/register", s.register)
	e.GET("/api/login", s.login)

	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root:   "static",
	// 	Browse: true,
	// }))
	// e.File("/", "static/index.html")
	// e.File("/index", "static/index.html")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", c.Host, c.Port)))

	return &s
}

// Run - start server host
func (s *Server) Run() {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port)))
}
