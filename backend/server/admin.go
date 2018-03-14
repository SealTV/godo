package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func mainAdmin(c echo.Context) error {
	return sendResponse(c, http.StatusOK, "hoary you are on the secret admin main page!", nil)
}
