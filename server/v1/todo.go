package v1

import (
	"bitbucket.org/SealTV/go-site/data"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"bitbucket.org/SealTV/go-site/model"
)

func GetTodos(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "tasks")
	}
}

func AddTodo(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(model.Todo)
		if err := c.Bind(todo); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.AddTodo(*todo)
		if err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}
		return c.JSON(http.StatusCreated, H{"todo": result})
	}
}

func UpdateTodo(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(model.Todo)
		if err := c.Bind(todo); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.UpdateTodo(*todo)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteTodo(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		todosDelete, err := db.DeleteTodoById(id)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}

		log.Debug("Deleted todos:", todosDelete)
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
