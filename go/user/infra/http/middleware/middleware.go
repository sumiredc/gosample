package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Use(e *echo.Echo) {
	// echo Logger
	e.Use(middleware.Logger())
	// CORS
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
			},
			AllowOrigins: []string{"*"},
		}))
	// Add Context
	e.Use(useContext)
}
