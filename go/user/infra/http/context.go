package http

import (
	"context"
	"sample/user/infra/database/mysql"

	"gorm.io/gorm"
)

type contextKey string

// NOTE: Context key names
const (
	userRepositoryKey contextKey = "UserRepository"
)

func buildRepositoriesToContext(c context.Context, r, w *gorm.DB) context.Context {
	c = context.WithValue(c, userRepositoryKey, mysql.NewUserRepository(w, r))
	return c
}
