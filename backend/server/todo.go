package server

import (
	"net/http"

	"bitbucket.org/SealTV/go-site/backend/model"
	"github.com/labstack/echo"
)

func (s *Server) getTodos(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return sendResponse(c, http.StatusFailedDependency, nil, err)
	}

	todos, err := s.db.GetAllTodosForUser(user)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, todos, nil)
}

func (s *Server) addTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	result, err := s.db.AddTodo(todo)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusCreated, result, nil)
}

func (s *Server) updateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	result, err := s.db.UpdateTodo(todo)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, result, nil)
}

func (s *Server) deleteTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	todosDelete, err := s.db.DeleteTodoById(todo.Id)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, todosDelete, nil)
}
