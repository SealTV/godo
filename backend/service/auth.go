package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SealTV/godo/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type (
	registerData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (s *Service) authenticator(c *gin.Context) (interface{}, error) {
	var username string
	var password string
	_, err := s.db.GetUserByLoginAndPassword(username, password)
	if err != nil {
		return username, nil
	}

	return username, nil
}

func (s *Service) payloadFunc(data interface{}) jwt.MapClaims {
	var username string
	log.Println("username:", username)
	u, err := s.db.GetUserByLogin(username)
	log.Println("User", u)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	m := make(map[string]interface{})
	m["userId"] = u.ID
	m["login"] = u.Login
	m["email"] = u.Email
	return m
}

func (s *Service) authorizator(data interface{}, c *gin.Context) bool {
	return true
}

func (s *Service) unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func (s *Service) register(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	fmt.Println(string(b))
	rd := registerData{}
	err = json.Unmarshal(b, &rd)
	fmt.Println(rd)
	if err != nil {
		log.Println("Can't unmarshal user data")
		c.Status(http.StatusBadRequest)
		return
	}

	u := model.User{
		Login:    rd.Username,
		Email:    rd.Email,
		Password: rd.Password,
	}
	u, err = s.db.AddUser(u)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, u)
}

func (s *Service) verify(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	c.JSON(200, gin.H{
		"userId": claims["userId"],
		"login":  claims["login"],
		"email":  claims["email"],
	})
}

func (s *Service) delete(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id := claims["userId"].(int)

	_, err := s.db.DeleteUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Service) logout(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	if claims["userId"] != nil && claims["login"] != nil {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusBadRequest)
}
