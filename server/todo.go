package server

import (
	"net/http"
	"strconv"

	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func (s *Server) GetTodos(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	db := s.db
	id, _ := strconv.Atoi(claims["jti"].(string))
	user, err := db.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User are not found")
	}
	todos, err := s.db.GetAllTodosForUser(user)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}
	return c.JSON(http.StatusOK, todos)
}

func AddTodo(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(model.Todo)
		if err := c.Bind(todo); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.AddTodo(*todo)
		if err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}
		return c.JSON(http.StatusCreated, H{"todo": result})
	}
}

func UpdateTodo(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(model.Todo)
		if err := c.Bind(todo); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.UpdateTodo(*todo)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteTodo(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		todosDelete, err := db.DeleteTodoById(id)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}

		log.Debug("Deleted todos:", todosDelete)
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
