package server

import (
	"fmt"

	"bitbucket.org/SealTV/go-site/backend/data"
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

type response struct {
	Result interface{} `json:"result"`
	Error  error       `json:"error"`
}

// New - create new instance of servers
func New(db data.DBConnector, c Config) *Server {
	e := echo.New()
	s := Server{db, e, c}
	adminGroup := e.Group("/api/admin")

	// this logs the server interaction
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: "[${time_unix}] ${id} " +
			"${status}  ${method} " +
			"${host}${path} ${latency_human} " +
			"${bytes_in}:${bytes_out} " +
			"${form} " +
			"\n",
	}))
	e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	log.Printf("REQUEST: %s\n", string(reqBody))
	// 	log.Printf("RESPONSE: %s\n", string(resBody))
	// }))

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:8080", "http://localhost:3000", "https://localhost:8080"},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))

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

	e.POST("/auth/register", s.register)
	e.POST("/auth/login", s.login)
	e.GET("/auth/user", s.user)
	e.POST("/auth/logout", s.logout)

	e.GET("/api/ping", s.ping)

	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root:   "static",
	// 	Browse: true,
	// }))
	// e.File("/", "static/index.html")
	// e.File("/index", "static/index.html")

	return &s
}

// Run - start server host
func (s *Server) Run() {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf("%s:%d", s.c.Host, s.c.Port)))
}

func sendResponse(c echo.Context, status int, result interface{}, err error) error {
	return c.JSON(status, response{result, err})
}
