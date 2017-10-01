package v1

import (
	"bitbucket.org/SealTV/go-site/data"
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func GetTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "tasks")
	}
}

func AddTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(data.Todo)
		if err := c.Bind(todo); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result := data.AddTodo(db, *todo)
		return c.JSON(http.StatusCreated, H{"todo": result})
	}
}

func UpdateTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(data.Todo)
		if err := c.Bind(todo); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := data.UpdateTodo(db, *todo)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		todosDelete, err := data.DeleteTodoById(db, id)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}

		log.Debug("Deleted todos:", todosDelete)
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
