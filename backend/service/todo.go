package service

import (
	"net/http"

	"github.com/SealTV/godo/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func (s *Service) addNewTodo(c *gin.Context) {
	var json model.Todo
	var err error

	claims := jwt.ExtractClaims(c)
	userID := claims["userId"].(int)

	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	json.UserID = userID
	var todo model.Todo
	if todo, err = s.db.AddTodo(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"result": todo})
}

func (s *Service) getAllTodos(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := claims["userId"].(int)

	var todos model.TodoCollection
	var err error

	if todos, err = s.db.GetAllTodosForUser(model.User{ID: userID}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": todos})
}

func (s *Service) updateAllTodos(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (s *Service) deleteAllTodos(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (s *Service) getTodo(c *gin.Context) {
	id := c.GetInt("id")
	todo, err := s.db.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": todo})
}

func (s *Service) addTodo(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
}

func (s *Service) updateTodo(c *gin.Context) {
	id := c.GetInt("id")
	var json model.Todo
	var err error
	if err = c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	json.ID = id

	var count int64
	if count, err = s.db.UpdateTodo(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"result": count})
}

func (s *Service) deleteTodo(c *gin.Context) {
	id := c.GetInt("id")
	var count int64
	var err error

	if count, err = s.db.DeleteTodoById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"result": count})
}
