package middleware

import (
	"context"
	domainctx "sample/user/domain/context"
	"sample/user/infra/database/mysql"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func useContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var err error

		ctx, err = addRepositories(ctx)
		if err != nil {
			return err
		}

		ctx, err = addValidator(ctx)
		if err != nil {
			return err
		}

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}

}

func addRepositories(c context.Context) (context.Context, error) {
	w, r, err := mysql.Connection()

	if err != nil {
		return nil, err
	}

	return context.WithValue(c, domainctx.UserRepository, mysql.NewUserRepository(w, r)), nil
}

func addValidator(c context.Context) (context.Context, error) {
	v := validator.New()
	return context.WithValue(c, domainctx.Validator, v), nil
}
