package repository

import (
	"application-template/internal/pkg/db"
	"context"
	"github.com/go-pg/pg"

	"application-template/internal/app/user/datastruct"
)

// GetUsersByIDs ...
func (r *UserRepository) GetUsersByIDs(ctx context.Context, ids []int64) ([]*datastruct.User, error) {
	users := make([]*datastruct.User, 0)
	err := db.GetDB(ctx).Model(&users).Where("id in (?)", pg.In(ids)).Select()
	return users, err
}
