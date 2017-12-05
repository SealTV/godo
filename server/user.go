package server

import (
	"net/http"
	"strconv"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (s *Server) getUser(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	id, _ := strconv.Atoi(claims["jti"].(string))
	user, err := s.db.GetUserById(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) getUserModel(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	id, _ := strconv.Atoi(claims["jti"].(string))
	user, err := s.db.GetUserModel(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) updateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.UpdateUser(*user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, H{"updated": result})
}

func (s *Server) deleteUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "Invalid value")
	}

	result, err := s.db.DeleteUser(*user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, H{"delete": result})
}
