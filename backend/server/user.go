package server

import (
	"net/http"

	"github.com/SealTV/godo/model"
	"github.com/labstack/echo"
)

func (s *Server) getUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return sendResponse(c, http.StatusFailedDependency, nil, err)
	}

	user, err := s.db.GetUserById(u.ID)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	return sendResponse(c, http.StatusOK, user, nil)
}

func (s *Server) getUserModel(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return sendResponse(c, http.StatusFailedDependency, nil, err)
	}

	user, err := s.db.GetUserModel(u.ID)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	return sendResponse(c, http.StatusOK, user, nil)
}

func (s *Server) updateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	result, err := s.db.UpdateUser(*u)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, result, err)
}

func (s *Server) deleteUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}

	result, err := s.db.DeleteUser(*u)
	if err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	return sendResponse(c, http.StatusOK, result, nil)
}
