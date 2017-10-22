package server

import (
	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUser(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		id, _ := strconv.Atoi(claims["jti"].(string))
		user, err := db.GetUserById(id)

		if err != nil {
			return c.String(http.StatusBadRequest, "User not found")
		}

		return c.JSON(http.StatusOK, user)
	}
}

func GetUserModel(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		id, _ := strconv.Atoi(claims["jti"].(string))
		user, err := db.GetUserModel(id)

		if err != nil {
			return c.String(http.StatusBadRequest, "User not found")
		}

		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.UpdateUser(*user)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"updated": result})
	}
}

func DeleteUser(db data.DBConnector) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			c.String(http.StatusFailedDependency, "Invalid value")
		}

		result, err := db.DeleteUser(*user)
		if err != nil {
			return c.String(http.StatusFailedDependency, err.Error())
		}
		return c.JSON(http.StatusOK, H{"DeleteUser": result})
	}
}
