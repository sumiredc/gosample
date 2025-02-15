package http

import (
	"context"
	"sample/user/infra/database/mysql"

	"gorm.io/gorm"
)

func buildRepositoriesToContext(c context.Context, r, w *gorm.DB) context.Context {
	c = context.WithValue(c, "UserRepository", mysql.NewUserRepository(w, r))
	return c
}
