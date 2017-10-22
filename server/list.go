package server

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type H map[string]interface{}

func GetList(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		listId, err := strconv.Atoi(c.QueryParam("listId"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid list id")
		}

		list, err := db.GetUserById(listId)
		if err != nil {
			return c.String(http.StatusNotFound, fmt.Sprintf("List by id: %d are not found", listId))
		}
		return c.JSON(http.StatusOK, H{"list": list})
	}
}
func AddList(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		list := new(model.List)
		if err := c.Bind(list); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.AddList(*list)
		if err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		return c.JSON(http.StatusCreated, H{"list": result})
	}
}

func UpdateList(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		list := new(model.List)
		if err := c.Bind(list); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.UpdateList(*list)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteList(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		deleted, err := db.DeleteListById(id)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}

		log.Debug("Deleted todos:", deleted)
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
