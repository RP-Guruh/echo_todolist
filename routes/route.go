package routes

import (
	"todolist/app/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", controllers.GetListUser)

	return e
}
