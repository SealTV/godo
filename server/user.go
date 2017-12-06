package server

import (
	"net/http"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
)

func (s *Server) getUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	user, err := s.db.GetUserById(u.Id)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) getUserModel(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	user, err := s.db.GetUserModel(u.Id)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) updateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.UpdateUser(*u)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, H{"updated": result})
}

func (s *Server) deleteUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.DeleteUser(*u)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, H{"delete": result})
}
