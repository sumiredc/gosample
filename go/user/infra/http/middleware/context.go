package middleware

import (
	domainctx "sample/user/domain/context"
	"sample/user/infra/database/mysql"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func useContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		addRepositories(c)
		addValidator(c)

		return next(c)
	}

}

func addRepositories(c echo.Context) error {
	w, r, err := mysql.Connection()

	if err != nil {
		return err
	}

	c.Set(domainctx.UserRepository, mysql.NewUserRepository(w, r))

	return nil
}

func addValidator(c echo.Context) {
	c.Set(domainctx.Validator, validator.New())
}
