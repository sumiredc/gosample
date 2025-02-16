package http

import (
	"context"
	"sample/user/infra/database/mysql"

	"gorm.io/gorm"
)

type contextKey int

// NOTE: Context key names
const userRepository contextKey = iota

func buildRepositoriesToContext(c context.Context, r, w *gorm.DB) context.Context {
	c = context.WithValue(c, userRepository, mysql.NewUserRepository(w, r))
	return c
}
