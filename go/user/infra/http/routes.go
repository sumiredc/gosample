package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	apiV1 := api.Group("/v1")
	apiV1.GET("/user", getUserList)
	apiV1.POST("/user", createUser)
	apiV1.GET("/user/:id", getUser)
	apiV1.PUT("/user/:id", updateUser)
	apiV1.DELETE("/user/:id", deleteUser)

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
