package service

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func (s *Service) authenticator(userID string, password string, c *gin.Context) (string, bool) {
	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return userID, true
	}

	return userID, false
}

func (s *Service) authorizator(userID string, c *gin.Context) bool {
	if userID == "admin" {
		return true
	}

	return false
}

func (s *Service) unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func (s *Service) helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userId": claims["id"],
		"text":   "Hello World.",
	})
}

func (s *Service) logout(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userId": claims["id"],
		"text":   "Hello World.",
	})
}
