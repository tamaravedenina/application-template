package service

import (
	"context"

	"application-template/internal/app/user/datastruct"
)

// GetBooksByIDs ...
func (s *UserService) GetUsersByIDs(ctx context.Context, ids []int64) ([]*datastruct.User, error) {
	return s.repository.GetUsersByIDs(ctx, ids)
}
