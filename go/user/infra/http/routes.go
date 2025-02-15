package http

import "github.com/labstack/echo"

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/user", getUserList)
	api.POST("/user", saveUser)
	api.GET("/user/:id", getUser)
	api.PUT("/user/:id", updateUser)
	api.DELETE("/user/:id", deleteUser)
}
