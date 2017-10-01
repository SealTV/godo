package v1

import (
	"bitbucket.org/SealTV/go-site/data"
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetList(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		listId, err := strconv.Atoi(c.QueryParam("listId"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid list id")
		}

		list, err := data.GetUserById(db, listId)
		if err != nil {
			return c.String(http.StatusNotFound, fmt.Sprintf("List by id: %s are not found", listId))
		}
		return c.JSON(http.StatusOK, H{"list": list})
	}
}
func AddList(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		list := new(data.List)
		if err := c.Bind(list); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result := data.AddList(db, *list)
		return c.JSON(http.StatusCreated, H{"list": result})
	}
}

func UpdateList(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		list := new(data.List)
		if err := c.Bind(list); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := data.UpdateList(db, *list)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteList(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		deleted, err := data.DeleteListById(db, id)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}

		log.Debug("Deleted todos:", deleted)
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}