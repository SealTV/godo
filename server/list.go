package server

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type H map[string]interface{}

func (s *Server) getList(c echo.Context) error {
	listId, err := strconv.Atoi(c.QueryParam("listId"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid list id")
	}

	list, err := s.db.GetUserById(listId)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("List by id: %d are not found", listId))
	}
	return c.JSON(http.StatusOK, H{"list": list})
}

func (s *Server) addList(c echo.Context) error {
	list := new(model.List)
	if err := c.Bind(list); err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	result, err := s.db.AddList(*list)
	if err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	return c.JSON(http.StatusCreated, H{"list": result})
}

func (s *Server) updateList(c echo.Context) error {
	list := new(model.List)
	if err := c.Bind(list); err != nil {
		c.String(http.StatusFailedDependency, "Invalid value")
	}

	result, err := s.db.UpdateList(*list)
	if err != nil {
		return c.String(http.StatusFailedDependency, err.Error())
	}
	return c.JSON(http.StatusOK, H{"updated": result})
}

func (s *Server) deleteList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	deleted, err := s.db.DeleteListById(id)
	if err != nil {
		return c.String(http.StatusFailedDependency, err.Error())
	}

	log.Debug("Deleted todos:", deleted)
	return c.JSON(http.StatusOK, H{
		"deleted": id,
	})
}
