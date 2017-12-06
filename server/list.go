package server

import (
	"net/http"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
)

// H - map
type H map[string]interface{}

func (s *Server) getLists(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	list, err := s.db.GetAllListsForUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, list)
}

func (s *Server) addList(c echo.Context) error {
	var list model.List
	if err := c.Bind(&list); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.AddList(list)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, result)
}

func (s *Server) updateList(c echo.Context) error {
	var list model.List
	if err := c.Bind(&list); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.UpdateList(list)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) deleteList(c echo.Context) error {
	var list model.List
	if err := c.Bind(&list); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	deleted, err := s.db.DeleteListById(list.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, deleted)
}
