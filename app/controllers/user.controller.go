package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "get list all user")
}
