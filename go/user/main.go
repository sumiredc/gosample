package main

import (
	"sample/user/infra/http"
	"sample/user/infra/http/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware
	middleware.Use(e)

	// Routes
	http.SetupRoutes(e)

	// Error Handler

	e.Start(":8080")
}
