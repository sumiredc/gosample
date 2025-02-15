package http

import (
	"sample/user/infra/database/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func UseMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(connectMySQL)
}

func connectMySQL(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		writer, reader, err := mysql.Connection()
		if err != nil {
			return err
		}

		ctx := buildRepositoriesToContext(c.Request().Context(), writer, reader)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
