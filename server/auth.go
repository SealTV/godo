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
		Name string `json:"name"`
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
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusCreated, u)
}

func (s *Server) login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	user, err := s.db.GetUserByLoginAndPassword(username, password)

	if err != nil {
		return echo.ErrUnauthorized
		//return c.String(http.StatusUnauthorized, "Your username or password were wrong")
	}

	// check username and password against DB after hashing the password
	cookie := &http.Cookie{}

	// this is the same
	//cookie := new(http.Cookie)

	cookie.Name = "sessionID"
	cookie.Value = "some_string"
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(cookie)

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
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"], "Token valid: ", token.Valid)

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}

func createJwtToken(user model.User) (string, error) {
	claims := jwtClaims{
		user.Login,
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
