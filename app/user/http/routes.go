package http

import "github.com/labstack/echo"

func Setup(e *echo.Echo) {
	e.GET("/user", getUserList)
	e.POST("/user", saveUser)
	e.GET("/user/:id", getUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)
}
