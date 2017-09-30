package server

import (
	"bitbucket.org/SealTV/go-site/data"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.QueryParam("username")
		password := c.QueryParam("password")

		user, err := data.GetUserByLoginAndPassword(db, username, password)

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
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token":   token,
		})
	}
}

func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}

func createJwtToken(user data.User) (string, error) {
	//userJson, _ := json.Marshal(&user)
	claims := JwtClaims{
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
