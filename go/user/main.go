package main

import (
	"sample/user/infra/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Middleware
	http.UseMiddlewares(e)

	// Routes
	http.SetupRoutes(e)

	// Error Handler

	e.Start(":8080")
}
