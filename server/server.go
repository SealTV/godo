package server

import (
	"database/sql"

	"bitbucket.org/SealTV/go-site/server/api/v1"
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
	jwtGroup := e.Group("/jwt")

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))

	// this logs the server interaction
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//jwtGroup.Use(middleware.JWT([]byte("mySecret")))
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		//TokenLookup: "query:token",
	}))
	jwtGroup.GET("/main", mainJwt)

	cookieGroup.Use(checkCookie)
	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)

	e.GET("/login", login(db))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.File("/index", "static/index.html")
	e.File("/todo", "static/todo.html")
	e.Static("/static", "assets")

	e.GET("/tasks", v1.GetTasks(db))
	e.PUT("/tasks", v1.PutTask(db))
	e.DELETE("/tasks/:id", v1.DeleteTask(db))

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
