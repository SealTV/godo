package server

import (
	"net/http"

	"github.com/SealTV/godo/model"
	"github.com/labstack/echo"
)

// H - map
type H map[string]interface{}

func (s *Server) getLists(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return sendResponse(c, http.StatusFailedDependency, nil, err)
	}

	list, err := s.db.GetAllListsForUser(user)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, list, nil)
}

func (s *Server) addList(c echo.Context) error {
	var list model.List
	if err := c.Bind(&list); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	result, err := s.db.AddList(list)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusCreated, result, nil)
}

func (s *Server) updateList(c echo.Context) error {
	var list model.List
	if err := c.Bind(&list); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	result, err := s.db.UpdateList(list)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, result, nil)
}

func (s *Server) deleteList(c echo.Context) error {
	var list model.List
	if err := c.Bind(&list); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	deleted, err := s.db.DeleteListById(list.ID)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, deleted, nil)
}
