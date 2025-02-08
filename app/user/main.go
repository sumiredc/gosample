package main

import (
	"sample/app/user/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	http.Setup(e)

	e.Start(":8080")
}
