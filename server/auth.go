package server

import (
	"encoding/json"
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

	auth struct {
		Token string     `json:"token"`
		User  model.User `json:"user"`
	}
	customValidator struct {
		validator *validator.Validate
	}
)

func (s *Server) register(c echo.Context) error {
	u := model.User{}

	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&u); err != nil {
		return sendResponse(c, http.StatusBadRequest, nil, err)
	}
	defer c.Request().Body.Close()

	// if err := c.Bind(&u); err != nil {
	// 	log.Println(err)
	// 	return sendResponse(c, http.StatusBadRequest, nil, err)
	// }

	u, err := s.db.AddUser(u)
	if err != nil {
		return sendResponse(c, http.StatusInternalServerError, nil, err)
	}
	return sendResponse(c, http.StatusCreated, u, nil)
}

func (s *Server) login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	user, err := s.db.GetUserByLoginAndPassword(username, password)
	if err != nil {
		log.Println(err)
		return sendResponse(c, http.StatusNotFound, nil, err)
	}

	// create jwt token
	token, err := createJwtToken(user)
	if err != nil {
		return sendResponse(c, http.StatusInternalServerError, nil, err)
	}

	return sendResponse(c, http.StatusOK, auth{token, user}, nil)
}

func (s *Server) mainJwt(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtClaims)
	name := claims.Name
	str := fmt.Sprintf("Hello %s, you is admin=%v", name, claims.Admin)
	return sendResponse(c, http.StatusOK, str, nil)
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
