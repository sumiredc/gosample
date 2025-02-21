package http

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/user", getUserList)
	api.POST("/user", createUser)
	api.GET("/user/:id", getUser)
	api.PUT("/user/:id", updateUser)
	api.DELETE("/user/:id", deleteUser)

	e.HTTPErrorHandler = jsonErrorHandler
}

func jsonErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	c.JSON(code, http.StatusText(code))
}
