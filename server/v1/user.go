package v1

import (
	"bitbucket.org/SealTV/go-site/data"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		id, _ := strconv.Atoi(claims["jti"].(string))
		user, err := data.GetUserById(db, int(id))

		if err != nil {
			return c.String(http.StatusBadRequest, "User not found")
		}

		return c.JSON(http.StatusOK, user)
	}
}

func GetUserModel(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		id, _ := strconv.Atoi(claims["jti"].(string))
		user, err := data.GetUserModel(db, id)

		if err != nil {
			return c.String(http.StatusBadRequest, "User not found")
		}

		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(data.User)
		if err := c.Bind(user); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := data.UpdateUser(db, *user)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "DeleteUser")
	}
}
