package v1

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "tasks")
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{"created": 123})
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
