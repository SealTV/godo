package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"bitbucket.org/SealTV/go-site/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v8"
)

type (
	jwtClaims struct {
		Name  string `json:"name"`
		Admin bool   `json:"admin"`
		jwt.StandardClaims
	}
	user struct {
		Name     string `json:"name" form:"name" query:"name" validate:"required"`
		Email    string `json:"email" form:"email" query:"email" validate:"required, email"`
		Password string `json:"password" form:"password" query:"password"`
	}

	customValidator struct {
		validator *validator.Validate
	}
)

func (s *Server) register(c echo.Context) error {
	u := model.User{
		Login:    c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	u, err := s.db.AddUser(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, u)
}

func (s *Server) login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	user, err := s.db.GetUserByLoginAndPassword(username, password)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	// create jwt token
	token, err := createJwtToken(user)
	if err != nil {
		log.Println("Error Creating JWT token", err)
		return c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "You were logged in!",
		"token":   token,
	})
}

func (s *Server) mainJwt(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtClaims)
	name := claims.Name
	str := fmt.Sprintf("Hello %s, you is admin=%v", name, claims.Admin)
	return c.String(http.StatusOK, str)
}

func createJwtToken(user model.User) (string, error) {
	claims := jwtClaims{
		user.Login,
		false,
		jwt.StandardClaims{
			Id:        fmt.Sprint(user.Id),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
