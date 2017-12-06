package server

import (
	"net/http"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
)

func (s *Server) getTodos(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	todos, err := s.db.GetAllTodosForUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, todos)
}

func (s *Server) addTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.AddTodo(todo)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}
	return c.JSON(http.StatusCreated, result)
}

func (s *Server) updateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.UpdateTodo(todo)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) deleteTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	todosDelete, err := s.db.DeleteTodoById(todo.Id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todosDelete)
}
