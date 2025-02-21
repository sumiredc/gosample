package http

import (
	"context"
	"sample/user/infra/database/mysql"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type contextKey int

// NOTE: Context key names
const userRepository contextKey = iota
const requestValidator contextKey = iota

func buildRepositoriesToContext(c context.Context, r, w *gorm.DB) context.Context {
	return context.WithValue(c, userRepository, mysql.NewUserRepository(w, r))
}

func buildValidatorToContext(c context.Context, v *validator.Validate) context.Context {
	return context.WithValue(c, requestValidator, v)
}
